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

type Theme struct {
	baseTheme *Theme
	prefix    string
	dict      map[string]interface{}
}

func NewTheme(prefix string, baseTheme *Theme) *Theme {
	theme := &Theme{
		baseTheme: baseTheme,
		prefix:    prefix,
		dict:      make(map[string]interface{}),
	}
	return theme
}

func (t *Theme) getKey(key string) string {
	if t.prefix != "" {
		return t.prefix + "." + key
	}
	return key
}

func (t *Theme) getValue(key string) interface{} {

	prefixAndKey := t.getKey(key)

	value, found := t.dict[prefixAndKey]
	if found {
		return value
	}

	if t.baseTheme != nil {
		value, found := t.baseTheme.dict[prefixAndKey]
		if found {
			return value
		}
	}

	return nil
}

func (t *Theme) setValue(key string, value interface{}) {

	t.dict[t.getKey(key)] = value
}

func (t *Theme) Color(key string) Attribute {

	value := t.getValue(key)
	if value != nil {
		return value.(Attribute)
	}

	return ColorDefault
}

func (t *Theme) ColorTermbox(key string) termbox.Attribute {

	return termbox.Attribute(t.Color(key))
}

func (t *Theme) SetColor(key string, color Attribute) {

	t.setValue(key, color)
}

func (t *Theme) Border(key string) LineType {
	value := t.getValue(key)
	if value != nil {
		return LineType(value.(int))
	}

	return LineNone
}

func (t *Theme) SetBorder(key string, border LineType) {

	t.setValue(key, int(border))
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

var defaultTheme *Theme

func init() {
	defaultTheme = NewTheme("", nil)

	defaultTheme.setValue("app.fg", ColorBlue)
	defaultTheme.setValue("app.bg", ColorBlue)

	defaultTheme.setValue("window.fg", ColorBlack)
	defaultTheme.setValue("window.bg", ColorWhite)
	defaultTheme.setValue("window.border", LineSingle)
	defaultTheme.setValue("window.border.fg", ColorBlack)
	defaultTheme.setValue("window.border.bg", ColorWhite)
	defaultTheme.setValue("window.title.fg", ColorBlack)
	defaultTheme.setValue("window.title.bg", ColorWhite)

	defaultTheme.setValue("label.fg", ColorBlack)
	defaultTheme.setValue("label.bg", ColorWhite)

	defaultTheme.setValue("button.fg", ColorBlack)
	defaultTheme.setValue("button.bg", ColorWhite)
	defaultTheme.setValue("button.border", LineNone)
	defaultTheme.setValue("button.border.fg", ColorBlack)
	defaultTheme.setValue("button.border.bg", ColorWhite)
	defaultTheme.setValue("button.focused.fg", ColorMagenta)
	defaultTheme.setValue("button.focused.bg", ColorBlue)
}
