package gonsole

type Panel struct {
	BaseControl
	BaseContainer
}

func NewPanel(win *Window, parent Container, id string) *Panel {
	panel := &Panel{}
	panel.BaseControl.Init(win, parent, id, "panel")
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

		t := p.Theme()
		fg, bg := t.ColorTermbox("fg"), t.ColorTermbox("bg")
		DrawTextSimple(" "+p.Title()+" ", false, p.BorderBox().Minus(Sides{Left: 2}), fg, bg)
	}

	// content area (ContainerControl already takes care of drawing the children)
}
