package gonsole

type BaseContainer struct {
	title    string
	children []Control
}

func (c *BaseContainer) Title() string {
	return c.title
}

func (c *BaseContainer) SetTitle(title string) {
	c.title = title
}

func (c *BaseContainer) AddControl(ctrl Control) {
	c.children = append(c.children, ctrl)
}

func (c *BaseContainer) Children() []Control {
	return c.children
}

func (c *BaseContainer) ChildrenDeep() []Control {
	controls := make([]Control, 0)
	for _, control := range c.children {
		container, ok := control.(Container)
		if ok {
			children := container.ChildrenDeep()
			for _, child := range children {
				controls = append(controls, child)
			}
		} else {
			controls = append(controls, control)
		}
	}
	return controls
}

func (c *BaseContainer) DirtyChildren() bool {
	for _, child := range c.ChildrenDeep() {
		if child.Dirty() {
			return true
		}
	}

	return false
}

func (c *BaseContainer) RepaintChildren() {
	for _, child := range c.Children() {
		child.Repaint()
	}
}
