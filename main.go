package main

import (
	"fmt"
	bd "reversi/models/board"
	pl "reversi/players"
)

func main() {
	var board = bd.NewBoardDefault()

	var is_over = false
	for !is_over {
		fmt.Printf("\n%s\n", board.LegalMode(board.Turn))
		pl.Human(board)
		is_over = board.IsOver()
	}
	fmt.Println()
	switch board.Judge() {
	case bd.BlackWins:
		fmt.Println("Black wins.")
	case bd.WhiteWins:
		fmt.Println("White wins.")
	case bd.Draw:
		fmt.Println("Draw.")
	}
	fmt.Printf("\n%s", board.LegalMode(board.Turn))
}
