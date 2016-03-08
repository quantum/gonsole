// Higher level helper functions for termbox
// TODO support theming
package gonsole

import (
	"github.com/mitchellh/go-wordwrap"
	"github.com/nsf/termbox-go"
)

func ClearRect(box Box, foreground, background Attribute) {
	fg := termboxAttr(foreground)
	bg := termboxAttr(background)
	for x := box.Left; x < box.Right(); x++ {
		for y := box.Top; y < box.Bottom(); y++ {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
	}
}

func FillRect(box Box, foreground, background Attribute) {
	fg := termboxAttr(foreground)
	bg := termboxAttr(background)
	for x := box.Left; x < box.Right(); x++ {
		for y := box.Top; y < box.Bottom(); y++ {
			termbox.SetCell(x, y, ' ', fg, bg)
		}
	}
}

func DrawBorder(box Box, lineType LineType, foreground, background Attribute) {
	right := box.Right()
	bottom := box.Bottom()
	runes := getLineRunes(lineType)
	// draw box.Top and bottom lines
	for y := box.Top; y < box.Top+box.Height; y = y + box.Height - 1 {
		DrawLineHorizontal(box.Left+1, y, box.Width, lineType, foreground, background)
	}
	// draw box.Left and right lines
	for x := box.Left; x < box.Left+box.Width; x = x + box.Width - 1 {
		DrawLineVertical(x, box.Top+1, box.Height, lineType, foreground, background)
	}
	// draw corners
	fg := termboxAttr(foreground)
	bg := termboxAttr(background)
	termbox.SetCell(box.Left, box.Top, runes[2], fg, bg)
	termbox.SetCell(right, box.Top, runes[3], fg, bg)
	termbox.SetCell(box.Left, bottom, runes[4], fg, bg)
	termbox.SetCell(right, bottom, runes[5], fg, bg)
}

func DrawLineHorizontal(left, top, width int, lineType LineType, foreground, background Attribute) {
	fg := termboxAttr(foreground)
	bg := termboxAttr(background)

	for x := left; x < left+width-1; x++ {
		termbox.SetCell(x, top, getLineRunes(lineType)[0], fg, bg)
	}
}

func DrawLineVertical(left, top, height int, lineType LineType, foreground, background Attribute) {
	fg := termboxAttr(foreground)
	bg := termboxAttr(background)

	for y := top; y < top+height-1; y++ {
		termbox.SetCell(left, y, getLineRunes(lineType)[1], fg, bg)
	}
}

func DrawCursor() {
}

// TODO support line breaking for multiline strings
// TODO support alignment
func DrawTextBox(text string, box Box, foreground, background Attribute) {

	fg := termboxAttr(foreground)
	bg := termboxAttr(background)

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

func DrawTextSimple(text string, box Box, foreground, background Attribute) {
	fg := termboxAttr(foreground)
	bg := termboxAttr(background)
	index := 0
	for _, char := range text {
		termbox.SetCell(box.Left+index, box.Top, char, fg, bg)
		index++
	}
}

func getLineRunes(lineType LineType) []rune {
	// https://en.wikipedia.org/wiki/Box-drawing_character
	var runes []rune
	switch lineType {
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
