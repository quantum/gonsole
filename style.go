package gonsole

import "github.com/nsf/termbox-go"

type Attribute uint16

const (
	ColorDefault Attribute = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

const (
	AttrBold Attribute = 1 << (iota + 9)
	AttrUnderline
	AttrReverse
)

type LineType int

const (
	LineNone = iota
	LineTransparent
	LineSingle
	LineSingleCorners
	LineDouble
	LineDoubleCorners
	LineDashed
	LineDotted
)

type HorizontalAlignment int

const (
	HorizontalAlignmentLeft = iota
	HorizontalAlignmentCenter
	HorizontalAlignmentRight
)

type VerticalAlignment int

const (
	HorizontalAlignmentTop = iota
	HorizontalAlignmentMiddle
	HorizontalAlignmentBottom
)

type Style struct {
	Fg Attribute
	Bg Attribute

	Margin  Sides
	Padding Sides

	Border   LineType
	BorderBg Attribute
	BorderFg Attribute

	ScrollFg Attribute
	ScrollBg Attribute

	SelectedFg Attribute
	SelectedBg Attribute

	HAlign HorizontalAlignment
	VAlign VerticalAlignment
}

func ColorRGB(r, g, b int) Attribute {
	within := func(n int) int {
		if n < 0 {
			return 0
		}

		if n > 5 {
			return 5
		}

		return n
	}

	r, b, g = within(r), within(b), within(g)
	return Attribute(0x0f + 36*r + 6*g + b)
}

func termboxAttr(attr Attribute) termbox.Attribute {
	return termbox.Attribute(attr)
}
