package build

import (
	"practice/app/reversi/game"
	"practice/gen/proto"
)

func ToPBRoom(r *game.Room) *proto.Room {
	return &proto.Room{
		Id:   r.ID,
		Host: ToPBPlayer(r.Host),
	}
}

func ToPBPlayer(p *game.Player) *proto.Player {
	if p == nil {
		return nil
	}
	return &proto.Player{
		Id:    p.ID,
		Color: ToPBColor(p.Color),
	}
}

func ToPBColor(c game.Color) proto.Color {
	switch c {
	case game.Black:
		return proto.Color_BLACK
	case game.White:
		return proto.Color_WHITE
	case game.Empty:
		return proto.Color_EMPTY
	case game.Wall:
		return proto.Color_BLACK
	}
	return proto.Color_WALL
}

func ToPBBoard(b *game.Board) *proto.Board {
	pbCols := make([]*proto.Board_Col, 0, 10)
	for _, col := range b.Cells {
		pbCells := make([]proto.Color, 0, 10)
		for _, c := range col {
			pbCells = append(pbCells, ToPBColor(c))
		}
		pbCols = append(pbCols, &proto.Board_Col{
			Cells: pbCells,
		})
	}
	return &proto.Board{
		Cols: pbCols,
	}
}
