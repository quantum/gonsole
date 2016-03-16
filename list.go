package gonsole

import "github.com/nsf/termbox-go"

type List struct {
	BaseControl

	options       []string
	selectedIndex int
	topIndex      int
}

func NewList(win AppWindow, parent Container, id string) *List {
	list := &List{}
	list.Init(win, parent, id, "list")
	list.SetFocusable(true)
	parent.AddControl(list)
	return list
}

func (l *List) Options() []string {
	return l.options
}

func (l *List) SetOptions(options []string) {
	l.options = options
}

func (l *List) Repaint() {
	if !l.Dirty() {
		return
	}
	l.BaseControl.Repaint()

	contentBox := l.ContentBox()

	t := l.Theme()
	focused := ""
	if l.Focused() {
		focused = "focused."
	}

	count := len(l.options)
	if count > contentBox.Height {
		count = contentBox.Height

		pos := ScrollPos(l.selectedIndex, len(l.options), contentBox.Height)
		fg, bg := t.ColorTermbox(focused+"scroll.fg"), t.ColorTermbox(focused+"scroll.bg")
		DrawScrollBar(contentBox.Right(), contentBox.Top, contentBox.Height, pos, fg, bg)
		contentBox = contentBox.Minus(Sides{Right: 1})
	}

	for i := 0; i < count; i++ {
		fg, bg := t.ColorTermbox(focused+"fg"), t.ColorTermbox(focused+"bg")

		if i+l.topIndex == l.selectedIndex {
			fg, bg = t.ColorTermbox(focused+"selected.fg"), t.ColorTermbox(focused+"selected.bg")
		}

		DrawTextSimple(l.options[l.topIndex+i], true, Box{contentBox.Left, contentBox.Top + i, contentBox.Width, 1}, fg, bg)
	}
}

func (l *List) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowDown:
			if l.selectedIndex < len(l.options)-1 {
				l.selectedIndex++
				contentBox := l.ContentBox()
				if l.selectedIndex == l.topIndex+contentBox.Height {
					l.topIndex++
				}
			}
			l.GetWindow().App().Redraw()
			return true

		case termbox.KeyArrowUp:
			if l.selectedIndex > 0 {
				l.selectedIndex--
				if l.selectedIndex < l.topIndex {
					l.topIndex--
				}
			}
			l.GetWindow().App().Redraw()
			return true

		case termbox.KeyHome:
			l.selectedIndex = 0
			l.topIndex = 0
			return true

		case termbox.KeyEnd:
			l.selectedIndex = len(l.options) - 1
			contentBox := l.ContentBox()
			if contentBox.Height > 0 {
				l.topIndex = len(l.options) - contentBox.Height
			}
			l.GetWindow().App().Redraw()
			return true

		case termbox.KeySpace:
			fallthrough
		case termbox.KeyEnter:
			m := make(map[string]interface{})
			m["index"] = l.selectedIndex
			l.SubmitEvent(&Event{"selected", l, m})
			return true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
