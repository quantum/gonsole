package gonsole

import "github.com/nsf/termbox-go"

type Checkbox struct {
	BasicControl

	// custom
	Text    string
	Checked bool
}

func NewCheckbox(id string) *Checkbox {
	checkbox := &Checkbox{}
	checkbox.Init(id)
	checkbox.SetFocussable(true)
	return checkbox
}

func (c *Checkbox) Repaint() {
	if !c.Dirty() {
		return
	}
	c.BasicControl.Repaint()

	// Box
	var icon string
	if c.Checked {
		icon = "☑"
	} else {
		icon = "☐"
	}

	contentBox := c.ContentBox()

	style := c.GetStyle()

	DrawTextSimple(icon, false, contentBox, style.Fg, style.Bg)

	DrawTextBox(c.Text, contentBox.Minus(Sides{Left: 2}), style.Fg, style.Bg)
}

func (chk *Checkbox) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEnter:
			fallthrough
		case termbox.KeySpace:
			// change state
			chk.Checked = !chk.Checked
			// events
			if chk.Checked {
				chk.SubmitEvent(&Event{"checked", chk, nil})
			} else {
				chk.SubmitEvent(&Event{"unchecked", chk, nil})
			}
			chk.SubmitEvent(&Event{"changed", chk, nil})
			return true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false
}
