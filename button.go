package gonsole

import "github.com/nsf/termbox-go"

type Button struct {
	BaseControl

	text    string
	onClick func()
}

func NewButton(win AppWindow, parent Container, id string) *Button {
	button := &Button{}
	button.Init(win, parent, id, "button")
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

func (b *Button) OnClick(handler func()) {
	b.onClick = handler
}

func (b *Button) Repaint() {
	if !b.Dirty() {
		return
	}

	b.BaseControl.Repaint()

	t := b.Theme()
	cb := b.ContentBox()
	if b.Focused() {
		DrawTextSimple(b.text, false, cb, t.ColorTermbox("focused.fg"), t.ColorTermbox("focused.bg"))
	} else {
		DrawTextSimple(b.text, false, cb, t.ColorTermbox("fg"), t.ColorTermbox("bg"))
	}
}

func (b *Button) ParseEvent(ev *termbox.Event) (handled, repaint bool) {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			if b.onClick != nil {
				b.onClick()
			}
			return true, true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false, false
}
