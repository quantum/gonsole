package gonsole

type Panel struct {
	BaseControl
	BaseContainer
}

func NewPanel(win *Window, parent Container, id string) *Panel {
	panel := &Panel{}
	panel.BaseControl.Init(win, parent, id)
	parent.AddControl(panel)
	return panel
}

func (p *Panel) Repaint() {
	if !p.Dirty() {
		return
	}
	p.BaseControl.Repaint()
	p.BaseContainer.RepaintChildren()

	// draw title
	if p.Title() != "" {
		if p.BorderType() == LineNone {
			p.SetPadding(p.Padding().Plus(Sides{Top: 1}))
		}

		DrawTextSimple(" "+p.Title()+" ", false, p.BorderBox().Minus(Sides{Left: 2}), p.fg, p.bg)
	}

	// content area (ContainerControl already takes care of drawing the children)
}
