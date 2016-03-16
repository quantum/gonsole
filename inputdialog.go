package gonsole

import (
	"fmt"
	"strconv"

	"github.com/huandu/xstrings"
)

type InputDialog struct {
	BaseWindow
}

func NewInputDialog(app *App, id, title, message string, buttons []string) *InputDialog {
	d := &InputDialog{}
	d.Init(app, id)
	d.App().addWindow(d)
	d.SetTitle(title)

	label := NewLabel(d, d, fmt.Sprintf("%s__message", id))
	label.SetPosition(Position{"0", "0", "100%", "80%"})
	label.SetText(message)

	edit := NewEdit(d, d, "edit")
	edit.SetPosition(Position{"10%", "10%", "80%", "1"})
	edit.Focus()
	edit.AddEventListener("submit", func(ev *Event) bool {
		d.App().eventDispatcher.SubmitEvent(&Event{"closed", d, ev.Data})
		d.Close()
		return true
	})

	buttonCount := len(buttons)

	for i, button := range buttons {
		textLen := xstrings.Len(button)
		btn := NewButton(d, d, fmt.Sprintf("%s__button%d", id, i))
		btn.SetPosition(Position{fmt.Sprintf("%d%%-%d", (i*buttonCount+1)*100/(buttonCount*2), textLen/2), "80%", strconv.Itoa(textLen), "1"})
		btn.SetText(button)

		btn.AddEventListener("clicked", func(ev *Event) bool {
			m := make(map[string]interface{})
			m["index"] = i
			d.App().eventDispatcher.SubmitEvent(&Event{"closed", d, m})
			d.Close()
			return true
		})
	}
	return d
}
