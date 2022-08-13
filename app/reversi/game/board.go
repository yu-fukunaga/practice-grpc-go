package game

import (
	"fmt"
)

type Board struct {
	Cells [][]Color
}

// 盤面を作成
func NewBoard() *Board {
	// セルを作成
	b := &Board{
		Cells: make([][]Color, 10), // 列を10個作成
	}
	for i := 0; i < 10; i++ {
		b.Cells[i] = make([]Color, 10) // 列ごとにセルを10個作成
	}

	// 外周を壁で埋める
	for i := 0; i < 10; i++ {
		b.Cells[0][i] = Wall
		b.Cells[9][i] = Wall
	}
	for i := 1; i < 9; i++ {
		b.Cells[i][0] = Wall
		b.Cells[i][9] = Wall
	}

	// 初期石を置く
	b.Cells[4][4] = White
	b.Cells[5][5] = White
	b.Cells[5][4] = Black
	b.Cells[4][5] = Black

	return b
}

func (b *Board) CanPutStone(x int32, y int32, c Color) bool {
	// 置いた場所がEmptyでなければFALSE
	if b.Cells[x][y] != Empty {
		return false
	}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if b.CountTurnableStonesByDirection(x, y, c, int32(dx), int32(dy)) > 0 {
				return true
			}
		}
	}
	return false
}

func (b *Board) PutStone(x int32, y int32, c Color) error {
	if !b.CanPutStone(x, y, c) {
		return fmt.Errorf("Can not put stone x=%v, y=%v color=%v\n", x, y, c)
	}
	b.Cells[x][y] = c

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			b.TurnStonesByDirection(x, y, c, int32(dx), int32(dy))
		}
	}
	return nil
}

func (b *Board) CountTurnableStonesByDirection(x int32, y int32, c Color, dx int32, dy int32) int {
	cnt := 0
	nx := x + dx
	ny := y + dy
	for {
		nc := b.Cells[nx][ny]
		// 敵の色であればカウント
		if nc == OpponentColor(c) {
			cnt++
		} else {
			break
		}
		nx += dx
		ny += dy
	}
	if cnt > 0 && b.Cells[nx][ny] == c {
		return cnt
	}
	return 0
}

func (b *Board) TurnStonesByDirection(x int32, y int32, c Color, dx int32, dy int32) {
	nx := x + dx
	ny := y + dy
	if b.CountTurnableStonesByDirection(x, y, c, dx, dy) > 0 {
		for {
			nc := b.Cells[nx][ny]
			// 敵の色であればひっくり返す
			if nc == OpponentColor(c) {
				b.Cells[nx][ny] = c
			} else {
				break
			}
			nx += dx
			ny += dy
		}
	}
}

func (b *Board) CountScore() (black int, white int, empty int) {
	for _, col := range b.Cells {
		for _, cell := range col {
			switch cell {
			case Black:
				black++
			case White:
				white++
			case Empty:
				empty++
			default:
			}
		}
	}
	return
}

func (b *Board) CountAvaliableCell(c Color) int {
	cnt := 0

	for i := 1; i < 9; i++ {
		for j := 1; j < 9; j++ {
			if b.CanPutStone(int32(i), int32(j), c) {
				cnt++
			}
		}
	}

	return cnt
}
