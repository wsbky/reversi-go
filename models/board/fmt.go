package board

import (
	"fmt"
	"reversi/models/coordinate"
)

func (b *Board) String() string {
	r := ""

	r += "   1 2 3 4 5 6 7 8  \n"
	r += "  ----------------- \n"

	for i := 0; i < 8; i++ {
		r += fmt.Sprintf("%d| ", i+1)

		for j := 0; j < 8; j++ {
			place, _ := coordinate.NewCoordinate(j, i)
			if b.Black&place.Bit != 0 {
				r += "●"
			} else if b.White&place.Bit != 0 {
				r += "○"
			} else if b.showLegalputs && b.CanPut(*place, b.Turn) {
				r += "+"
			} else {
				r += "∙"
			}

			if j != 7 {
				r += " "
			}
		}
		r += " |\n"
	}
	r += "  ----------------- "

	return r + "\n"
}
