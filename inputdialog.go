package gonsole

import (
	"fmt"
	"strconv"

	"github.com/quantum/castle-installer/Godeps/_workspace/src/github.com/huandu/xstrings"
)

type InputDialog struct {
	BaseWindow

	edit        *Edit
	buttonIndex int
}

func (d *InputDialog) InputValue() string {
	return d.edit.Value()
}

func (d *InputDialog) SelectedButton() int {
	return d.buttonIndex
}

func NewInputDialog(app *App, id, title, message string, buttons []string) *InputDialog {
	d := &InputDialog{}
	d.Init(app, id)
	d.App().addWindow(d)
	d.SetTitle(title)
	d.SetPadding(Sides{1, 1, 1, 1})

	label := NewLabel(d, d, fmt.Sprintf("%s__message", id))
	label.SetPosition(Position{"0", "0", "100%", "50%"})
	label.SetText(message)

	d.edit = NewEdit(d, d, "edit")
	d.edit.SetPosition(Position{"0", "50%+1", "100%", "1"})
	d.edit.Focus()

	buttonCount := len(buttons)

	for i, button := range buttons {
		textLen := xstrings.Len(button)
		btn := NewButton(d, d, fmt.Sprintf("%s__button%d", id, i))
		btn.SetPosition(Position{fmt.Sprintf("%d%%-%d", (i*buttonCount+1)*100/(buttonCount*2), textLen/2), "90%", strconv.Itoa(textLen), "1"})
		btn.SetText(button)

		btnIndex := i

		btn.AddEventListener("clicked", func(ev *Event) bool {
			d.buttonIndex = btnIndex
			d.App().eventDispatcher.SubmitEvent(&Event{"closed", d, nil})
			d.Close()
			return true
		})
	}
	return d
}
