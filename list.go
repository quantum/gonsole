package gonsole

import "github.com/nsf/termbox-go"

type List struct {
	BasicControl

	Options []string

	SelectedIndex int
	topIndex      int
}

func NewList(id string) *List {
	list := &List{}
	list.Init(id)
	list.SetFocussable(true)
	return list
}

func (l *List) Repaint() {
	l.BasicControl.Repaint()

	contentBox := l.ContentBox()

	count := len(l.Options)
	if count > contentBox.Height {
		count = contentBox.Height

		style := l.GetStyle()
		pos := ScrollPos(l.SelectedIndex, len(l.Options), contentBox.Height)
		DrawScrollBar(contentBox.Right(), contentBox.Top, contentBox.Height, pos, style.ScrollFg, style.ScrollBg)
		contentBox = contentBox.Minus(Sides{Right: 1})
	}

	for i := 0; i < count; i++ {
		style := l.GetStyle()
		fg := style.Fg
		bg := style.Bg

		if i+l.topIndex == l.SelectedIndex {
			fg = style.SelectedFg
			bg = style.SelectedBg
		}

		DrawTextSimple(l.Options[l.topIndex+i], true, Box{contentBox.Left, contentBox.Top + i, contentBox.Width, 1}, fg, bg)
	}
}

func (l *List) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowDown:
			if l.SelectedIndex < len(l.Options)-1 {
				l.SelectedIndex++
				contentBox := l.ContentBox()
				if l.SelectedIndex == l.topIndex+contentBox.Height {
					l.topIndex++
				}
			}
			return true

		case termbox.KeyArrowUp:
			if l.SelectedIndex > 0 {
				l.SelectedIndex--
				if l.SelectedIndex < l.topIndex {
					l.topIndex--
				}
			}
			return true

		case termbox.KeyHome:
			l.SelectedIndex = 0
			l.topIndex = 0
			return true

		case termbox.KeyEnd:
			l.SelectedIndex = len(l.Options) - 1
			contentBox := l.ContentBox()
			if contentBox.Height > 0 {
				l.topIndex = len(l.Options) - contentBox.Height
			}
			return true

		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			m := make(map[string]interface{})
			m["index"] = l.SelectedIndex
			l.SubmitEvent(&Event{"selected", l, m})
			return true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
