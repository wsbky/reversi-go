package board

import (
	cl "reversi/models/color"
)

type Judge int

const (
	BlackWins Judge = iota
	WhiteWins
	Draw
)

func (b *Board) Judge() Judge {
	var blackCount = b.popcountFor(cl.Black)
	var whiteCount = b.popcountFor(cl.White)

	if blackCount > whiteCount {
		return BlackWins
	} else if blackCount < whiteCount {
		return WhiteWins
	} else {
		return Draw
	}
}
