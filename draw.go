// Higher level helper functions for termbox
// TODO support theming
package gonsole

import (
	"github.com/mitchellh/go-wordwrap"
	"github.com/nsf/termbox-go"
)

func ClearRect(box Box, fg, bg termbox.Attribute) {
	for x := box.Left; x < box.Right(); x++ {
		for y := box.Top; y < box.Bottom(); y++ {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
	}
}

func FillRect(box Box, fg, bg termbox.Attribute) {
	for x := box.Left; x <= box.Right(); x++ {
		for y := box.Top; y <= box.Bottom(); y++ {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
	}
}

func DrawBorder(box Box, lineType LineType, fg, bg termbox.Attribute) {
	right := box.Right()
	bottom := box.Bottom()
	runes := getLineRunes(lineType)

	DrawLineHorizontal(box.Left, box.Top, box.Width, runes[0], fg, bg)
	DrawLineHorizontal(box.Left, box.Bottom(), box.Width, runes[0], fg, bg)
	DrawLineVertical(box.Left, box.Top, box.Height, runes[1], fg, bg)
	DrawLineVertical(box.Right(), box.Top, box.Height, runes[1], fg, bg)

	termbox.SetCell(box.Left, box.Top, runes[2], fg, bg)
	termbox.SetCell(right, box.Top, runes[3], fg, bg)
	termbox.SetCell(box.Left, bottom, runes[4], fg, bg)
	termbox.SetCell(right, bottom, runes[5], fg, bg)
}

func DrawShadow(box Box, shadow termbox.Attribute) {
	bottom := box.Bottom()
	right := box.Right()
	DrawLineHorizontal(box.Left+1, bottom, box.Width, ' ', shadow, shadow)
	DrawLineVertical(right, box.Top+1, box.Height-1, ' ', shadow, shadow)
}

func DrawLineHorizontal(left, top, width int, ch rune, fg, bg termbox.Attribute) {
	for x := left; x < left+width-1; x++ {
		termbox.SetCell(x, top, ch, fg, bg)
	}
}

func DrawLineVertical(left, top, height int, ch rune, fg, bg termbox.Attribute) {
	for y := top; y < top+height-1; y++ {
		termbox.SetCell(left, y, ch, fg, bg)
	}
}

func ScrollPos(index, count, height int) int {
	pos := int(float32(index) / float32(count) * float32(height-2))
	return pos
}

func DrawScrollBar(left, top, height, pos int, fg, bg termbox.Attribute) {
	runes := []rune("░■▲▼")

	termbox.SetCell(left, top, runes[2], fg, bg)
	termbox.SetCell(left, top+height-1, runes[3], fg, bg)
	if height > 2 {
		for y := top + 1; y < top+height-1; y++ {
			termbox.SetCell(left, y, runes[0], fg, bg)
		}
	}
	if pos != -1 {
		termbox.SetCell(left, top+pos+1, runes[1], fg, bg)
	}
}

func DrawCursor(x, y int) {
	termbox.SetCursor(x, y)
}

func HideCursor() {
	termbox.HideCursor()
}

// TODO support line breaking for multiline strings
// TODO support alignment
func DrawTextBox(text string, box Box, fg, bg termbox.Attribute) {
	wrapText := wordwrap.WrapString(text, uint(box.Width))

	x := box.Left
	y := box.Top

	for _, char := range wrapText {
		if char == '\n' {
			x = box.Left
			y++
		} else {
			termbox.SetCell(x, y, char, fg, bg)
			x++
		}
	}
}

func DrawTextSimple(text string, fill bool, box Box, fg, bg termbox.Attribute) {
	index := 0
	for _, char := range text {
		termbox.SetCell(box.Left+index, box.Top, char, fg, bg)
		index++
	}
	if fill {
		for x := index; x < box.Width; x++ {
			termbox.SetCell(box.Left+x, box.Top, ' ', bg, bg)
		}
	}
}

func getLineRunes(lineType LineType) []rune {
	// https://en.wikipedia.org/wiki/Box-drawing_character
	var runes []rune
	switch lineType {
	case LineNone:
		panic("LineNone is not valid")
	case LineTransparent:
		runes = []rune{' ', ' ', ' ', ' ', ' ', ' '}
	case LineSingle:
		runes = []rune{'─', '│', '┌', '┐', '└', '┘'}
	case LineSingleCorners:
		runes = []rune{' ', ' ', '┌', '┐', '└', '┘'}
	case LineDouble:
		runes = []rune{'═', '║', '╔', '╗', '╚', '╝'}
	case LineDoubleCorners:
		runes = []rune{' ', ' ', '╔', '╗', '╚', '╝'}
	case LineDashed:
		runes = []rune{'╌', '╎', '┌', '┐', '└', '┘'}
	case LineDotted:
		runes = []rune{'┄', '┆', '┌', '┐', '└', '┘'}
	}
	return runes
}
