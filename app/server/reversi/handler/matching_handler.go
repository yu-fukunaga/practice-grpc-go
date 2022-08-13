package handler

import (
	"context"
	"fmt"
	"practice/app/reversi/build"
	"practice/app/reversi/game"
	"practice/gen/proto"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MatchingHandler struct {
	proto.UnimplementedReversiMathingServiceServer
	sync.RWMutex
	Rooms       map[int32]*game.Room
	maxPlayerID int32
}

func NewMatchingHandler() *MatchingHandler {
	return &MatchingHandler{
		Rooms: make(map[int32]*game.Room),
	}
}

func (h *MatchingHandler) JoinRoom(req *proto.JoinRoomRequest, stream proto.ReversiMathingService_JoinRoomServer) error {
	ctx, cancel := context.WithTimeout(stream.Context(), 2*time.Minute)
	defer cancel()

	h.Lock()

	// プレイヤーの新規作成
	h.maxPlayerID++
	me := &game.Player{
		ID: h.maxPlayerID,
	}

	// 空いている部屋を探す
	for _, room := range h.Rooms {
		if room.Guest == nil {
			me.Color = game.White
			room.Guest = me
			err := stream.Send(&proto.JoinRoomResponse{
				Status: proto.JoinRoomResponse_MATCHD,
				Room:   build.ToPBRoom(room),
				Me:     build.ToPBPlayer(room.Guest),
			})
			if err != nil {
				return err
			}
			h.Unlock()
			fmt.Printf("matched room_id=%v\n", room.ID)
			return nil
		}
	}

	// 開いてる部屋がなかったので部屋を作る
	me.Color = game.Black
	room := &game.Room{
		ID:   int32(len(h.Rooms)) + 1,
		Host: me,
	}

	h.Rooms[room.ID] = room
	h.Unlock()

	err := stream.Send(&proto.JoinRoomResponse{
		Status: proto.JoinRoomResponse_WAITING,
		Room:   build.ToPBRoom(room),
	})
	if err != nil {
		return err
	}

	ch := make(chan int)
	go func(ch chan<- int) {
		for {
			h.RLock()
			guest := room.Guest
			h.RUnlock()
			if guest != nil {
				err := stream.Send(&proto.JoinRoomResponse{
					Status: proto.JoinRoomResponse_MATCHD,
					Room:   build.ToPBRoom(room),
					Me:     build.ToPBPlayer(room.Host),
				})
				if err != nil {
					return
				}
				ch <- 0
				break
			}
			time.Sleep(1 * time.Second)

			select {
			case <-ctx.Done():
				return
			default:
			}
		}
	}(ch)

	select {
	case <-ch:
	case <-ctx.Done():
		return status.Errorf(codes.DeadlineExceeded, "マッチングできませんでした")
	}
	return nil
}
