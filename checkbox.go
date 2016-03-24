package gonsole

import "github.com/nsf/termbox-go"

type Checkbox struct {
	BaseControl

	text     string
	checked  bool
	onChange func(checked bool)
}

func NewCheckbox(win AppWindow, parent Container, id string) *Checkbox {
	checkbox := &Checkbox{}
	checkbox.Init(win, parent, id, "checkbox")
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

func (c *Checkbox) OnChange(handler func(checked bool)) {
	c.onChange = handler
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

	t := c.Theme()
	fg, bg := t.ColorTermbox("fg"), t.ColorTermbox("bg")
	contentBox := c.ContentBox()
	DrawTextSimple(icon, false, contentBox, fg, bg)
	DrawTextBox(c.text, contentBox.Minus(Sides{Left: 2}), fg, bg)
}

func (chk *Checkbox) ParseEvent(ev *termbox.Event) (handled, repaint bool) {
	switch ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyEnter:
			fallthrough
		case termbox.KeySpace:
			chk.checked = !chk.checked
			if chk.onChange != nil {
				chk.onChange(chk.checked)
			}
			return true, true
		}
	case termbox.EventError:
		panic(ev.Err)
	}

	return false, false
}
