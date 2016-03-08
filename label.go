package gonsole

type Label struct {
	BasicControl
	Text string
	//Alignment
}

func NewLabel(id string) *Label {
	label := &Label{}
	label.Init(id)
	return label
}

func (l *Label) Repaint() {
	l.BasicControl.Repaint()

	style := l.GetStyle()
	DrawTextBox(l.Text, l.ContentBox(), style.Fg, style.Bg)
}
