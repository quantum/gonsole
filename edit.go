package gonsole

import (
	xs "github.com/huandu/xstrings"

	"github.com/nsf/termbox-go"
)

type Edit struct {
	BaseControl

	value         string
	maxWidth      int
	cursorPos     int
	startingIndex int
}

func NewEdit(win AppWindow, parent Container, id string) *Edit {
	edit := &Edit{}
	edit.Init(win, parent, id, "edit")
	edit.SetFocusable(true)
	edit.SetCursorable(true)
	parent.AddControl(edit)
	return edit
}

func (e *Edit) Value() string {
	return e.value
}

func (e *Edit) SetValue(value string) {
	e.value = value
}

func (e *Edit) MaxWidth() int {
	return e.maxWidth
}

func (e *Edit) SetMaxWidth(width int) {
	e.maxWidth = width
}

func (e *Edit) Repaint() {
	if !e.Dirty() {
		return
	}
	e.BaseControl.Repaint()

	box := e.ContentBox()

	var shownValue string
	var cursorOffset int
	length := xs.Len(e.value)
	width := e.ContentBox().Width

	if e.startingIndex == 0 {
		if length < width {
			shownValue = e.value
		} else {
			shownValue = xs.Slice(e.value, 0, width)
		}
		cursorOffset = e.cursorPos
	} else {
		if length-e.startingIndex < width {
			shownValue = xs.Slice(e.value, e.startingIndex, -1)
		} else {
			shownValue = xs.Slice(e.value, e.startingIndex, e.startingIndex+width)
		}
		cursorOffset = e.cursorPos - e.startingIndex
	}

	t := e.Theme()
	fg, bg := t.ColorTermbox("fg"), t.ColorTermbox("bg")
	DrawTextSimple(shownValue, true, box, fg, bg)

	if e.Focused() {
		DrawTextSimple(" ", false, Box{box.Left+cursorOffset, box.Top, 1, 1}, t.ColorTermbox("cursor"), bg)
		DrawCursor(box.Left+cursorOffset, box.Top)
	}
}

func (e *Edit) handleChar(ch rune) {
	length := xs.Len(e.value)

	if e.maxWidth > 0 && length >= e.maxWidth {
		return
	}

	if e.cursorPos == 0 {
		e.value = string(ch) + e.value
	} else if e.cursorPos < length {
		e.value = xs.Slice(e.value, 0, e.cursorPos) + string(ch) + xs.Slice(e.value, e.cursorPos, -1)
	} else {
		e.value += string(ch)
	}

	e.cursorPos++

	width := e.ContentBox().Width

	if e.cursorPos >= width {
		e.startingIndex++
	}
}

func (e *Edit) handleBackspace() {
	if e.cursorPos == 0 || e.value == "" {
		return
	}

	length := xs.Len(e.value)

	if e.cursorPos >= length {
		e.cursorPos--
		e.value = xs.Slice(e.value, 0, length-1)
	} else if e.cursorPos == 1 {
		e.cursorPos = 0
		e.value = xs.Slice(e.value, 1, -1)
	} else {
		e.cursorPos--
		e.value = xs.Slice(e.value, 0, e.cursorPos) + xs.Slice(e.value, e.cursorPos+1, -1)
	}

	width := e.ContentBox().Width

	if width >= length-1 {
		e.startingIndex = 0
	}
}

func (e *Edit) handleDelete() {
	length := xs.Len(e.value)

	if e.cursorPos == length || e.value == "" {
		return
	}

	if e.cursorPos == length-1 {
		e.value = xs.Slice(e.value, 0, length-1)
	} else {
		e.value = xs.Slice(e.value, 0, e.cursorPos) + xs.Slice(e.value, e.cursorPos+1, -1)
	}

	width := e.ContentBox().Width

	if width >= length-1 {
		e.startingIndex = 0
	}
}

func (e *Edit) handleLeft() {
	if e.cursorPos == 0 || e.value == "" {
		return
	}

	if e.cursorPos == e.startingIndex {
		e.startingIndex--
	}

	e.cursorPos--
}

func (e *Edit) handleRight() {
	length := xs.Len(e.value)

	if e.cursorPos == length || e.value == "" {
		return
	}

	e.cursorPos++

	width := e.ContentBox().Width

	if e.cursorPos < length && e.cursorPos >= e.startingIndex+width-1 {
		e.startingIndex++
	}
}

func (e *Edit) handleHome() {
	e.cursorPos = 0
	e.startingIndex = 0
}

func (e *Edit) handleEnd() {
	length := xs.Len(e.value)

	e.cursorPos = length

	width := e.ContentBox().Width
	if length >= width {
		e.startingIndex = length - width + 1
	}
}

func (e *Edit) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeySpace:
			e.handleChar(' ')
			e.GetWindow().App().Redraw()
			return true
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			e.handleBackspace()
			e.GetWindow().App().Redraw()
			return true
		case termbox.KeyDelete:
			e.handleDelete()
			e.GetWindow().App().Redraw()
			return true
		case termbox.KeyArrowLeft:
			e.handleLeft()
			e.GetWindow().App().Redraw()
			return true
		case termbox.KeyArrowRight:
			e.handleRight()
			e.GetWindow().App().Redraw()
			return true
		case termbox.KeyHome:
			e.handleHome()
			e.GetWindow().App().Redraw()
			return true
		case termbox.KeyEnd:
			e.handleEnd()
			e.GetWindow().App().Redraw()
			return true
		case termbox.KeyEnter:
			m := make(map[string]interface{})
			m["value"] = e.value
			e.SubmitEvent(&Event{"submit", e, m})
			return true
		default:
			if ev.Ch != 0 {
				e.handleChar(ev.Ch)
				e.GetWindow().App().Redraw()
				return true
			}
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
