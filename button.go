package gonsole

import "github.com/nsf/termbox-go"

type Button struct {
	BaseControl

	text string
}

func NewButton(win AppWindow, parent Container, id string) *Button {
	button := &Button{}
	button.Init(win, parent, id)
	button.SetFocusable(true)
	parent.AddControl(button)
	return button
}

func (b *Button) Text() string {
	return b.text
}

func (b *Button) SetText(text string) {
	b.text = text
}

func (b *Button) Repaint() {
	if !b.Dirty() {
		return
	}
	b.BaseControl.Repaint()

	fg, bg := b.Colors()

	if b.Focused() {
		fg, bg = b.FocusColors()
	}
	DrawTextSimple(b.text, false, b.ContentBox(), fg, bg)
}

func (b *Button) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			b.SubmitEvent(&Event{"clicked", b, nil})
			return true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
