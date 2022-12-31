package coordinate

import (
	"fmt"
)

type Coordinate struct {
	Bit uint64
}

func NewCoordinate(x int, y int) (*Coordinate, error) {
	if x < 0 || x >= 8 || y < 0 || y >= 8 {
		return &Coordinate{}, &OutOfRangeError{x, y}
	}
	return &Coordinate{0x8000000000000000 >> (x + y*8)}, nil
}

func (c *Coordinate) Transfer(k int) *Coordinate {
	switch k {
	case 0:
		return &Coordinate{(c.Bit << 8) & 0xffffffffffffff00}
	case 1:
		return &Coordinate{(c.Bit << 7) & 0x7f7f7f7f7f7f7f00}
	case 2:
		return &Coordinate{(c.Bit >> 1) & 0x7f7f7f7f7f7f7f7f}
	case 3:
		return &Coordinate{(c.Bit >> 9) & 0x007f7f7f7f7f7f7f}
	case 4:
		return &Coordinate{(c.Bit >> 8) & 0x00ffffffffffffff}
	case 5:
		return &Coordinate{(c.Bit >> 7) & 0x00fefefefefefefe}
	case 6:
		return &Coordinate{(c.Bit << 1) & 0xfefefefefefefefe}
	case 7:
		return &Coordinate{(c.Bit << 9) & 0xfefefefefefefe00}
	default:
		return &Coordinate{0}
	}
}

type OutOfRangeError struct {
	x int
	y int
}

func (e *OutOfRangeError) Error() string {
	return fmt.Sprintf("OutOfRangeError: coordinate (x: %d, y: %d) is out of the board.", e.x, e.y)
}
