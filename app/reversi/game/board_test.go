package game

import (
	"testing"
)

func TestCanPutStone(t *testing.T) {
	board := NewBoard()
	tests := map[string]struct {
		board Board
		x     int32
		y     int32
		c     Color
		want  bool
	}{
		"ng: not empty": {
			board: *board,
			x:     4,
			y:     4,
			c:     Black,
			want:  false,
		},
		"ng: no turnable stones": {
			board: *board,
			x:     3,
			y:     4,
			c:     White,
			want:  false,
		},
		"ok": {
			board: *board,
			x:     3,
			y:     4,
			c:     Black,
			want:  true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.board.CanPutStone(tt.x, tt.y, tt.c)
			if tt.want != got {
				t.Errorf("mismatch got=%v want=%v\n", got, tt.want)
			}
		})
	}
}

func TestPutStone(t *testing.T) {
	board := NewBoard()
	tests := map[string]struct {
		board     Board
		x         int32
		y         int32
		c         Color
		wantErr   bool
		wantBoard Board
	}{
		"ok: 1st Black": {
			board:   *board,
			x:       4,
			y:       3,
			c:       Black,
			wantErr: false,
			wantBoard: Board{
				[][]Color{
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=0
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=1
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=2
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=3
					{Wall, Empty, Empty, Black, Black, Black, Empty, Empty, Empty, Wall}, // x=4
					{Wall, Empty, Empty, Empty, Black, White, Empty, Empty, Empty, Wall}, // x=5
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=6
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=7
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=8
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=9
				},
			},
		},
		"ok: 2nd White": {
			board:   *board,
			x:       5,
			y:       3,
			c:       White,
			wantErr: false,
			wantBoard: Board{
				[][]Color{
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=0
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=1
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=2
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=3
					{Wall, Empty, Empty, Black, Black, Black, Empty, Empty, Empty, Wall}, // x=4
					{Wall, Empty, Empty, White, White, White, Empty, Empty, Empty, Wall}, // x=5
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=6
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=7
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=8
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=9
				},
			},
		},
		"ok: 3th Black turn 2 directions": {
			board:   *board,
			x:       6,
			y:       3,
			c:       Black,
			wantErr: false,
			wantBoard: Board{
				[][]Color{
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=0
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=1
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=2
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=3
					{Wall, Empty, Empty, Black, Black, Black, Empty, Empty, Empty, Wall}, // x=4
					{Wall, Empty, Empty, Black, Black, White, Empty, Empty, Empty, Wall}, // x=5
					{Wall, Empty, Empty, Black, Empty, Empty, Empty, Empty, Empty, Wall}, // x=6
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=7
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=8
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=9
				},
			},
		},
		"ok: 4th White tuen 2 stone in 1 direction ": {
			board:   *board,
			x:       5,
			y:       2,
			c:       White,
			wantErr: false,
			wantBoard: Board{
				[][]Color{
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=0
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=1
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=2
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=3
					{Wall, Empty, Empty, Black, Black, Black, Empty, Empty, Empty, Wall}, // x=4
					{Wall, Empty, White, White, White, White, Empty, Empty, Empty, Wall}, // x=5
					{Wall, Empty, Empty, Black, Empty, Empty, Empty, Empty, Empty, Wall}, // x=6
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=7
					{Wall, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Empty, Wall}, // x=8
					{Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall, Wall},         // x=9
				},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.board.PutStone(tt.x, tt.y, tt.c)
			if err != nil && !tt.wantErr {
				t.Errorf("return err : %s", err)
			}
			for dx := 0; dx < 10; dx++ {
				for dy := 0; dy < 10; dy++ {
					if tt.board.Cells[dx][dy] != tt.wantBoard.Cells[dx][dy] {
						t.Errorf("mismatch after borad got(%v,%v)=%v, want(%v,%v)=%v", dx, dy, tt.board.Cells[dx][dy], dx, dy, tt.wantBoard.Cells[dx][dy].ToColorStr())
						return
					}
				}
			}
		})
	}
}

func TestCountTurnableStonesByDirection(t *testing.T) {
	board := NewBoard()
	tests := map[string]struct {
		board Board
		x     int32
		y     int32
		c     Color
		dx    int32
		dy    int32
		want  int
	}{
		"ok: Upper Left": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    -1,
			dy:    -1,
			want:  0,
		},
		"ok: Upper Middle": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    0,
			dy:    -1,
			want:  0,
		},
		"ok: Upper Right": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    1,
			dy:    -1,
			want:  0,
		},
		"ok: Middle Left": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    -1,
			dy:    0,
			want:  0,
		},
		"ok: Middle Right": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    1,
			dy:    0,
			want:  0,
		},
		"ok: Lower Left": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    -1,
			dy:    1,
			want:  0,
		},
		"ok: Lower Midle": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    0,
			dy:    1,
			want:  1,
		},
		"ok: Lower Right": {
			board: *board,
			x:     4,
			y:     3,
			c:     Black,
			dx:    1,
			dy:    1,
			want:  0,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := tt.board.CountTurnableStonesByDirection(tt.x, tt.y, tt.c, tt.dx, tt.dy)
			if got != tt.want {
				t.Errorf("mismatch got=%v want=%v\n", got, tt.want)
			}
		})
	}

}
