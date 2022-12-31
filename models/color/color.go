package color

type Color string

const (
	Black = Color("Black")
	White = Color("White")
)

func (c Color) Other() Color {
	switch c {
	case Black:
		return White
	case White:
		return Black
	default:
		return Black
	}
}
