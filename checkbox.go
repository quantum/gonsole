package gonsole

import "github.com/nsf/termbox-go"

type Checkbox struct {
	BaseControl

	text    string
	checked bool
}

func NewCheckbox(win AppWindow, parent Container, id string) *Checkbox {
	checkbox := &Checkbox{}
	checkbox.Init(win, parent, id)
	checkbox.SetFocusable(true)
	parent.AddControl(checkbox)
	return checkbox
}

func (c *Checkbox) Text() string {
	return c.text
}

func (c *Checkbox) SetText(text string) {
	c.text = text
}

func (c *Checkbox) Checked() bool {
	return c.checked
}

func (c *Checkbox) SeChecked(checked bool) {
	c.checked = checked
}

func (c *Checkbox) Repaint() {
	if !c.Dirty() {
		return
	}
	c.BaseControl.Repaint()

	var icon string
	if c.checked {
		icon = "☑"
	} else {
		icon = "☐"
	}

	contentBox := c.ContentBox()
	DrawTextSimple(icon, false, contentBox, c.fg, c.bg)
	DrawTextBox(c.text, contentBox.Minus(Sides{Left: 2}), c.fg, c.bg)
}

func (chk *Checkbox) ParseEvent(ev *termbox.Event) bool {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEnter:
			fallthrough
		case termbox.KeySpace:
			// change state
			chk.checked = !chk.checked
			// events
			if chk.checked {
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
