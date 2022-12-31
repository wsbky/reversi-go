package board

import (
	"reversi/models/color"
)

func NewBoard(black uint64, white uint64) *Board {
	return &Board{
		Black:         black,
		White:         white,
		showLegalputs: false,
		Turn:          color.Black,
	}
}

func NewBoardDefault() *Board {
	return NewBoard(0x0000000810000000, 0x0000001008000000)
}

func NewBoardFrom(other *Board) *Board {
	ret := *other
	return &ret
}
