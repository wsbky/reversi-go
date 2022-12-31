package players

import (
	"fmt"
	"os"
	bd "reversi/models/board"
	cl "reversi/models/color"
	cd "reversi/models/coordinate"
	"strconv"
)

func colorIcon(color cl.Color) string {
	switch color {
	case cl.Black:
		return "●"
	case cl.White:
		return "○"
	default:
		return "?"
	}
}

func Human(board *bd.Board) {
	var x int
	var y int

	var color = board.Turn

	if board.IsPass(board.Turn) {
		fmt.Printf("%s> passed\n", color)
		board.TurnOver()
		return
	}

	for {
		fmt.Printf("%s (x y): ", colorIcon(color))
		var input [2]string
		fmt.Scan(&input[0])
		if input[0] == "p" || input[0] == "pass" {
			fmt.Printf("%s> passed\n", color)
			board.TurnOver()
			return
		}
		if _, e := fmt.Scan(&input[1]); e != nil {
			fmt.Fprintln(os.Stderr, "Error occured while reading input.")
			continue
		}

		if s, e := strconv.Atoi(input[0]); e == nil {
			x = s - 1
		} else {
			continue
		}
		if s, e := strconv.Atoi(input[1]); e == nil {
			y = s - 1
		} else {
			continue
		}

		if place, err := cd.NewCoordinate(x, y); err == nil {
			if board.CanPut(*place, color) {
				break
			}
		}
	}

	fmt.Printf("%s> %d %d\n", colorIcon(color), x+1, y+1)
	place, _ := cd.NewCoordinate(x, y)
	board.Put(*place, color)
	board.TurnOver()
}
