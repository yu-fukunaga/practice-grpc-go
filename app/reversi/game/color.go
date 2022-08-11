package game

type Color int

const (
	Empty Color = iota
	Black       // ◯
	White       // ◉
	Wall
	UNKNOWN
)

func (c Color) ToColorStr() string {
	switch c {
	case Black:
		return "◯"
	case White:
		return "◉"
	case Empty:
		return " "
	default:
		return ""
	}
}

func OpponentColor(c Color) Color {
	switch c {
	case Black:
		return White
	case White:
		return Black
	default:
		return UNKNOWN
	}
}
