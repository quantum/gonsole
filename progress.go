package gonsole

import (
	"fmt"
	"strings"

	xs "github.com/huandu/xstrings"
)

type Progress struct {
	BasicControl

	// a value between 0 and 1
	Value float32
}

func NewProgress(id string) *Progress {
	p := &Progress{}
	p.Init(id)
	return p
}

func (p *Progress) Repaint() {
	p.BasicControl.Repaint()

	cb := p.ContentBox()

	text := strings.Repeat(" ", (cb.Width/2)-1)
	text += fmt.Sprintf("%d%%", int(p.Value*100))
	text += strings.Repeat(" ", (cb.Width/2)-3)

	style := p.GetStyle()
	percent := int(p.Value * float32(cb.Width))
	DrawTextSimple(xs.Slice(text, 0, percent), false, p.ContentBox(), style.Fg|AttrReverse, style.Bg)
	DrawTextSimple(xs.Slice(text, percent, -1), false, Box{cb.Left + percent, cb.Top, cb.Width - percent, cb.Height}, style.Fg, style.Bg)
}
