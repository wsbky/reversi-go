package board

import (
	cl "reversi/models/color"
	cd "reversi/models/coordinate"
)

func (b *Board) Put(place cd.Coordinate, color cl.Color) {
	var myBoard = *b.BitBoardOf(color)
	var opBoard = *b.BitBoardOf(color.Other())

	var rev uint64 = 0

	for i := 0; i < 8; i++ {
		var temp uint64 = 0
		var mask = place.Transfer(i)

		if mask.Bit != 0 && (mask.Bit&opBoard) != 0 {
			temp |= mask.Bit
			mask = mask.Transfer(i)
		}
		if mask.Bit&myBoard != 0 {
			rev |= temp
		}
	}

	myBoard ^= place.Bit | rev
	opBoard ^= rev

	*b.BitBoardOf(color) = myBoard
	*b.BitBoardOf(color.Other()) = opBoard
}
