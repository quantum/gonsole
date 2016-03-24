package gonsole

import (
	"fmt"
	"strconv"

	"github.com/huandu/xstrings"
)

type MessageDialog struct {
	BaseWindow

	buttonIndex int
}

func (d *MessageDialog) SelectedButton() int {
	return d.buttonIndex
}

func NewMessageDialog(app *App, id, title, message string, buttons []string) *MessageDialog {
	d := &MessageDialog{}
	d.Init(app, id)
	d.App().addWindow(d)
	d.SetTitle(title)
	d.SetPadding(Sides{1, 1, 1, 1})

	label := NewLabel(d, d, fmt.Sprintf("%s__message", id))
	label.SetPosition(Position{"0", "0", "100%", "80%"})
	label.SetText(message)

	buttonCount := len(buttons)

	for i, button := range buttons {
		textLen := xstrings.Len(button)
		btn := NewButton(d, d, fmt.Sprintf("%s__button%d", id, i))
		btn.SetPosition(Position{fmt.Sprintf("%d%%-%d", (i*buttonCount+1)*100/(buttonCount*2), textLen/2), "90%", strconv.Itoa(textLen), "1"})
		btn.SetText(button)

		if i == 0 {
			btn.Focus()
		}

		btnIndex := i

		btn.OnClick(func() {
			d.buttonIndex = btnIndex
			d.App().removeWindow(d)
			if d.onClose != nil {
				d.onClose()
			}
		})
	}
	return d
}
