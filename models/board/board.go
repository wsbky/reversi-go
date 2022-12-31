package board

import (
	cl "reversi/models/color"
	cd "reversi/models/coordinate"
)

type Board struct {
	Black uint64
	White uint64
	Turn  cl.Color

	showLegalputs bool
}

func (b *Board) LegalMode(color cl.Color) *Board {
	r := NewBoardFrom(b)
	r.showLegalputs = true
	r.Turn = color
	return r
}

func (b *Board) BitBoardOf(color cl.Color) *uint64 {
	switch color {
	case cl.Black:
		return &b.Black
	case cl.White:
		return &b.White
	}
	return &b.Black
}

func (b *Board) LegalBoard(color cl.Color) uint64 {
	var myBoard = *b.BitBoardOf(color)
	var opBoard = *b.BitBoardOf(color.Other())
	var blankBoard = ^(myBoard | opBoard)

	var horWatchOpBoard = opBoard & 0x7e7e7e7e7e7e7e7e
	var verWatchOpBoard = opBoard & 0x00ffffffffffff00
	var allWatchOpBoard = opBoard & 0x007e7e7e7e7e7e00

	var legalBoard uint64 = 0
	var temp uint64

	// left
	temp = horWatchOpBoard & (myBoard >> 1)
	for i := 0; i < 5; i++ {
		temp |= horWatchOpBoard & (temp >> 1)
	}
	legalBoard |= blankBoard & (temp >> 1)

	// right
	temp = horWatchOpBoard & (myBoard << 1)
	for i := 0; i < 5; i++ {
		temp |= horWatchOpBoard & (temp << 1)
	}
	legalBoard |= blankBoard & (temp << 1)

	// upper
	temp = verWatchOpBoard & (myBoard >> 8)
	for i := 0; i < 5; i++ {
		temp |= verWatchOpBoard & (temp >> 8)
	}
	legalBoard |= blankBoard & (temp >> 8)

	// lower
	temp = verWatchOpBoard & (myBoard << 8)
	for i := 0; i < 5; i++ {
		temp |= verWatchOpBoard & (temp << 8)
	}
	legalBoard |= blankBoard & (temp << 8)

	// upper left
	temp = allWatchOpBoard & (myBoard >> 9)
	for i := 0; i < 5; i++ {
		temp |= allWatchOpBoard & (temp >> 9)
	}
	legalBoard |= blankBoard & (temp >> 9)

	// upper right
	temp = allWatchOpBoard & (myBoard >> 7)
	for i := 0; i < 5; i++ {
		temp |= allWatchOpBoard & (temp >> 7)
	}
	legalBoard |= blankBoard & (temp >> 7)

	// lower left
	temp = allWatchOpBoard & (myBoard << 7)
	for i := 0; i < 5; i++ {
		temp |= allWatchOpBoard & (temp << 7)
	}
	legalBoard |= blankBoard & (temp << 7)

	// lower right
	temp = allWatchOpBoard & (myBoard << 9)
	for i := 0; i < 5; i++ {
		temp |= allWatchOpBoard & (temp << 9)
	}
	legalBoard |= blankBoard & (temp << 9)

	return legalBoard
}

func (b *Board) CanPut(place cd.Coordinate, color cl.Color) bool {
	return (place.Bit & b.LegalBoard(color)) == place.Bit
}

func (b *Board) IsPass(color cl.Color) bool {
	return b.LegalBoard(color) == 0 && b.LegalBoard(color.Other()) != 0
}

func (b *Board) IsOver() bool {
	return b.LegalBoard(cl.Black) == 0 && b.LegalBoard(cl.White) == 0
}

func (b *Board) TurnOver() {
	b.Turn = b.Turn.Other()
}

func (b *Board) popcountFor(color cl.Color) int {
	var x = *b.BitBoardOf(color)

	x = x - ((x >> 1) & 0x5555555555555555)

	x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)

	x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	x = x + (x >> 32)

	return int(x & 0x0000007f)
}
