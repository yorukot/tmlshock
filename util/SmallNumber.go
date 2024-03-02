package util

import "github.com/nsf/termbox-go"

func SmallColon(termWidth int, termHeight int, color termbox.Attribute, diff int) {

	drawString(termWidth/2+diff+2, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff+2, termHeight/2+1, "  ", color, color)

}

func SmallDot(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff+1, termHeight/2+2, "  ", color, color)
}

func SmallOne(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff+4, termHeight/2-2, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+2, "  ", color, color)
}

func SmallTwo(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2-0, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2+2, "      ", color, color)
}

func SmallThree(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2-0, "      ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2+2, "      ", color, color)
}

func SmallFour(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2-2, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2-1, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2, "      ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+2, "  ", color, color)
}

func SmallFive(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2-0, "      ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2+2, "      ", color, color)
}

func SmallSix(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2-0, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2+1, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2+2, "      ", color, color)

}

func SmallSeven(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+2, "  ", color, color)
}

func SmallEight(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2-1, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2+1, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2+2, "      ", color, color)
}

func SmallNine(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2-1, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2, "      ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff+4, termHeight/2+2, "  ", color, color)
}

func SmallZero(termWidth int, termHeight int, color termbox.Attribute, diff int) {
	drawString(termWidth/2+diff, termHeight/2-2, "      ", color, color)

	drawString(termWidth/2+diff, termHeight/2-1, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2-1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2+1, "  ", color, color)
	drawString(termWidth/2+diff+4, termHeight/2+1, "  ", color, color)

	drawString(termWidth/2+diff, termHeight/2+2, "      ", color, color)
}
