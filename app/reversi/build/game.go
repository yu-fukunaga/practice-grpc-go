package build

import (
	"fmt"
	"practice/app/reversi/game"
	"practice/gen/proto"
)

func ToRoom(r *proto.Room) *game.Room {
	return &game.Room{
		ID:    r.GetId(),
		Host:  ToPlayer(r.GetHost()),
		Guest: ToPlayer(r.GetGuest()),
	}
}
func ToPlayer(p *proto.Player) *game.Player {
	return &game.Player{
		ID:    p.GetId(),
		Color: ToColor(p.GetColor()),
	}
}
func ToColor(c proto.Color) game.Color {
	switch c {

	case proto.Color_BLACK:
		return game.Black
	case proto.Color_WHITE:
		return game.White
	case proto.Color_EMPTY:
		return game.Empty
	case proto.Color_WALL:
		return game.Wall
	}
	panic(fmt.Sprintf("unknwon color=%v", c))
}
