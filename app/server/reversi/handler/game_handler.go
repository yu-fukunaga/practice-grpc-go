package handler

import (
	"fmt"
	"practice/app/reversi/build"
	"practice/app/reversi/game"
	"practice/gen/proto"
	"sync"
)

type GameHandler struct {
	proto.UnimplementedGameServiceServer
	sync.RWMutex
	games  map[int32]*game.Game                     // ゲーム情報（盤面など）を格納する
	client map[int32][]proto.GameService_PlayServer // 状態変更時にクライアントにストリーミングを返すために格納する
}

func NewGameHandler() *GameHandler {
	return &GameHandler{
		games:  make(map[int32]*game.Game),
		client: make(map[int32][]proto.GameService_PlayServer),
	}
}

func (h *GameHandler) Play(stream proto.GameService_PlayServer) error {
	for {
		// クライアントからリクエストを受信したらreqにリクエストが代入されます
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		roomID := req.GetRoomId()
		player := build.ToPlayer(req.GetPlayer())

		// oneofで複数の型のリクエストがくるのでswitch文で処理します
		switch req.GetAction().(type) {
		case *proto.PlayRequest_Start:
			// ゲーム開始リクエスト
			err := h.start(stream, roomID, player)
			if err != nil {
				return err
			}
		case *proto.PlayRequest_Move:
			// 石を置いたときのリクエスト
			action := req.GetMove()
			x := action.GetMove().GetX()
			y := action.GetMove().GetY()
			err := h.move(roomID, x, y, player)
			if err != nil {
				return err
			}
		}
	}
}

func (h *GameHandler) start(stream proto.GameService_PlayServer, roomID int32, me *game.Player) error {

	h.Lock()
	defer h.Unlock()

	// ゲーム情報がなければ作成する
	g := h.games[roomID]
	if g == nil {
		g = game.NewGame(game.UNKNOWN)
		h.games[roomID] = g
		h.client[roomID] = make([]proto.GameService_PlayServer, 0, 2)
	}

	// 自分のクライアントを格納する
	h.client[roomID] = append(h.client[roomID], stream)

	if len(h.client[roomID]) == 2 {
		// 二人揃ったので開始する
		for _, s := range h.client[roomID] {
			// クライアントにゲーム開始を通知する
			err := s.Send(&proto.PlayResponse{
				Event: &proto.PlayResponse_Ready{
					Ready: &proto.PlayResponse_ReadyEvent{},
				},
			})
			if err != nil {
				return err
			}
		}
		fmt.Printf("game has started room_id=%v\n", roomID)
	} else {
		// まだ揃ってないので待機中であることをクライアントに通知する
		err := stream.Send(&proto.PlayResponse{
			Event: &proto.PlayResponse_Waiting{
				Waiting: &proto.PlayResponse_WaitingEvent{},
			},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (h *GameHandler) move(roomID int32, x int32, y int32, p *game.Player) error {
	h.Lock()
	defer h.Unlock()

	g := h.games[roomID]

	isFinished, err := g.MoveStone(x, y, p.Color)
	if err != nil {
		return err
	}

	for _, s := range h.client[roomID] {
		// 手が打たれたことをクライアントに通知する
		err := s.Send(&proto.PlayResponse{
			Event: &proto.PlayResponse_Move{
				Move: &proto.PlayResponse_MoveEvent{
					Player: build.ToPBPlayer(p),
					Move: &proto.Move{
						X: x,
						Y: y,
					},
					Board: build.ToPBBoard(g.Board),
				},
			},
		})
		if err != nil {
			return err
		}

		if isFinished {
			// ゲーム終了を通知する
			err := s.Send(
				&proto.PlayResponse{
					Event: &proto.PlayResponse_Finished{
						Finished: &proto.PlayResponse_FinishedEvent{
							Winner: build.ToPBColor(g.JudgeWinner()),
							Board:  build.ToPBBoard(g.Board),
						},
					},
				},
			)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
