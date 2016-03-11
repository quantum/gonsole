package gonsole

type Label struct {
	BaseControl

	text string
}

func NewLabel(win AppWindow, parent Container, id string) *Label {
	label := &Label{}
	label.Init(win, parent, id)
	parent.AddControl(label)
	return label
}

func (l *Label) Text() string {
	return l.text
}

func (l *Label) SetText(text string) {
	l.text = text
}

func (l *Label) Repaint() {
	if !l.Dirty() {
		return
	}
	l.BaseControl.Repaint()

	DrawTextBox(l.text, l.ContentBox(), l.fg, l.bg)
}
