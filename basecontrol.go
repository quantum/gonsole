package gonsole

// Control is the base model for a UI control
type BaseControl struct {
	BaseElement

	focusable  bool
	cursorable bool
}

func (c *BaseControl) Focusable() bool {
	return c.focusable
}

func (c *BaseControl) SetFocusable(focusable bool) {
	c.focusable = focusable
}

func (c *BaseControl) Focused() bool {
	if c.window.FocusedControl() != nil {
		if c.window.FocusedControl().ID() == c.ID() {
			return true
		}
	}
	return false
}

func (c *BaseControl) Focus() {
	c.window.FocusControl(c)
}

func (c *BaseControl) Cursorable() bool {
	return c.cursorable
}

func (c *BaseControl) SetCursorable(cursorable bool) {
	c.cursorable = cursorable
}
