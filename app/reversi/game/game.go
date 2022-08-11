package game

import "fmt"

type Game struct {
	Board      *Board
	isStarted  bool
	isFinished bool
	me         Color
}

func NewGame(me Color) *Game {
	return &Game{
		Board: NewBoard(),
		me:    me,
	}
}

func (g *Game) MoveStone(x int32, y int32, c Color) (isFinished bool, error error) {
	if g.isFinished {
		return true, nil
	}
	err := g.Board.PutStone(x, y, c)
	if err != nil {
		return false, err
	}
	g.Display(g.me)

	return false, nil
}

func (g *Game) IsGameOver() bool {
	if g.Board.CountAvaliableCell(Black) > 0 {
		return false
	}
	if g.Board.CountAvaliableCell(White) > 0 {
		return false
	}
	return true
}

func (g *Game) JudgeWinner() Color {
	black, white, _ := g.Board.CountScore()
	if black == white {
		return UNKNOWN
	} else if black > white {
		return Black
	}
	return White
}

// Display board
func (g *Game) Display(me Color) {
	fmt.Println("")
	if me != UNKNOWN {
		fmt.Printf("You: %v\n", me.ToColorStr())
	}

	fmt.Print(" |")
	rs := []rune("ABCDEFGH")
	for i, r := range rs {
		fmt.Printf("%v", string(r))
		if i < len(rs)-1 {
			fmt.Print(" |")
		}
	}
	fmt.Print("\n")
	fmt.Println("----------------------")

	for j := 1; j < 9; j++ {
		fmt.Printf("%d", j)
		fmt.Print(" |")
		for i := 1; i < 9; i++ {
			fmt.Print(g.Board.Cells[i][j])
			fmt.Print(" |")
		}
		fmt.Print("\n")
	}
	fmt.Println("----------------------")

	black, white, empty := g.Board.CountScore()
	fmt.Printf("Score: BLACK=%d, White=%d REST=%d\n", black, white, empty)

	fmt.Print("\n")
}
