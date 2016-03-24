package gonsole

import (
	"fmt"
	"strconv"

	"github.com/huandu/xstrings"
)

type SelectionDialog struct {
	BaseWindow

	list        *List
	buttonIndex int
}

func (d *SelectionDialog) SelectedItem() int {
	return d.list.SelectedItem()
}

func (d *SelectionDialog) SelectedButton() int {
	return d.buttonIndex
}

func NewSelectionDialog(app *App, id, title, message string, buttons []string, items []string) *SelectionDialog {
	d := &SelectionDialog{}
	d.Init(app, id)
	d.App().addWindow(d)
	d.SetTitle(title)
	d.SetPadding(Sides{1, 1, 1, 1})

	label := NewLabel(d, d, fmt.Sprintf("%s__message", id))
	label.SetPosition(Position{"0", "0", "100%", "20%"})
	label.SetText(message)

	d.list = NewList(d, d, "list")
	d.list.SetPosition(Position{"10%", "20%", "80%", "60%"})
	d.list.SetOptions(items)
	d.list.Focus()
	d.list.OnSumbit(func(selected int) {
		d.App().removeWindow(d)
		if d.onClose != nil {
			d.onClose()
		}
	})

	buttonCount := len(buttons)

	for i, button := range buttons {
		textLen := xstrings.Len(button)
		btn := NewButton(d, d, fmt.Sprintf("%s__button%d", id, i))
		btn.SetPosition(Position{fmt.Sprintf("%d%%-%d", (i*buttonCount+1)*100/(buttonCount*2), textLen/2), "90%", strconv.Itoa(textLen), "1"})
		btn.SetText(button)

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
