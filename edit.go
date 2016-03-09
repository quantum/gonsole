package gonsole

import (
	xs "github.com/huandu/xstrings"

	"github.com/nsf/termbox-go"
)

type Edit struct {
	BasicControl

	Value string
	// the max width of the value. By default set to the width of the control
	MaxWidth int

	cursorPos     int
	startingIndex int
}

func NewEdit(id string) *Edit {
	edit := &Edit{}
	edit.Init(id)
	edit.SetFocussable(true)
	edit.SetRequiresCursor(true)
	return edit
}

func (e *Edit) Repaint() {
	e.BasicControl.Repaint()

	style := e.GetStyle()
	box := e.ContentBox()

	var shownValue string
	var cursorOffset int
	length := xs.Len(e.Value)
	width := e.ContentBox().Width

	if e.startingIndex == 0 {
		if length < width {
			shownValue = e.Value
		} else {
			shownValue = xs.Slice(e.Value, 0, width)
		}
		cursorOffset = e.cursorPos
	} else {
		if length-e.startingIndex < width {
			shownValue = xs.Slice(e.Value, e.startingIndex, -1)
		} else {
			shownValue = xs.Slice(e.Value, e.startingIndex, e.startingIndex+width)
		}
		cursorOffset = e.cursorPos - e.startingIndex
	}

	DrawTextSimple(shownValue, true, box, style.Fg, style.Bg)

	if e.Focussed() {
		DrawCursor(box.Left+cursorOffset, box.Top)
	}
}

func (e *Edit) handleChar(ch rune) {
	length := xs.Len(e.Value)

	if e.MaxWidth > 0 && length >= e.MaxWidth {
		return
	}

	if e.cursorPos == 0 {
		e.Value = string(ch) + e.Value
	} else if e.cursorPos < length {
		e.Value = xs.Slice(e.Value, 0, e.cursorPos) + string(ch) + xs.Slice(e.Value, e.cursorPos, -1)
	} else {
		e.Value += string(ch)
	}

	e.cursorPos++

	width := e.ContentBox().Width

	if e.cursorPos >= width {
		e.startingIndex++
	}
}

func (e *Edit) handleBackspace() {
	if e.cursorPos == 0 || e.Value == "" {
		return
	}

	length := xs.Len(e.Value)

	if e.cursorPos >= length {
		e.cursorPos--
		e.Value = xs.Slice(e.Value, 0, length-1)
	} else if e.cursorPos == 1 {
		e.cursorPos = 0
		e.Value = xs.Slice(e.Value, 1, -1)
	} else {
		e.cursorPos--
		e.Value = xs.Slice(e.Value, 0, e.cursorPos) + xs.Slice(e.Value, e.cursorPos+1, -1)
	}

	width := e.ContentBox().Width

	if width >= length-1 {
		e.startingIndex = 0
	}
}

func (e *Edit) handleDelete() {
	length := xs.Len(e.Value)

	if e.cursorPos == length || e.Value == "" {
		return
	}

	if e.cursorPos == length-1 {
		e.Value = xs.Slice(e.Value, 0, length-1)
	} else {
		e.Value = xs.Slice(e.Value, 0, e.cursorPos) + xs.Slice(e.Value, e.cursorPos+1, -1)
	}

	width := e.ContentBox().Width

	if width >= length-1 {
		e.startingIndex = 0
	}
}

func (e *Edit) handleLeft() {
	if e.cursorPos == 0 || e.Value == "" {
		return
	}

	if e.cursorPos == e.startingIndex {
		e.startingIndex--
	}

	e.cursorPos--
}

func (e *Edit) handleRight() {
	length := xs.Len(e.Value)

	if e.cursorPos == length || e.Value == "" {
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
	length := xs.Len(e.Value)

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
			return true
		case termbox.KeyBackspace, termbox.KeyBackspace2:
			e.handleBackspace()
			return true
		case termbox.KeyDelete:
			e.handleDelete()
			return true
		case termbox.KeyArrowLeft:
			e.handleLeft()
			return true
		case termbox.KeyArrowRight:
			e.handleRight()
			return true
		case termbox.KeyHome:
			e.handleHome()
			return true
		case termbox.KeyEnd:
			e.handleEnd()
			return true
		case termbox.KeyEnter:
			m := make(map[string]interface{})
			m["value"] = e.Value
			e.SubmitEvent(&Event{"submit", e, m})
			return true
		default:
			if ev.Ch != 0 {
				e.handleChar(ev.Ch)
				return true
			}
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
