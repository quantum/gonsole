package gonsole

import (
	"fmt"
	"strconv"

	"github.com/huandu/xstrings"
)

type MessageDialog struct {
	BaseWindow
}

func NewMessageDialog(app *App, id, title, message string, buttons []string) *MessageDialog {
	mb := &MessageDialog{}
	mb.Init(app, id)
	mb.SetTitle(title)

	label := NewLabel(mb, mb, fmt.Sprintf("%s__message", id))
	label.SetText(message)
	label.SetPosition(Position{"0", "0", "100%", "80%"})

	buttonCount := len(buttons)

	for i, button := range buttons {
		textLen := xstrings.Len(button)
		btn := NewButton(mb, mb, fmt.Sprintf("%s__button%d", id, i))
		btn.SetPosition(Position{fmt.Sprintf("%d%%-%d", (i*buttonCount+1)*100/(buttonCount*2), textLen/2), "90%", strconv.Itoa(textLen), "1"})
		btn.SetFocusColors(ColorGreen, ColorRed)
		btn.SetText(button)

		if i == 0 {
			btn.Focus()
		}

		btn.AddEventListener("clicked", func(ev *Event) bool {
			m := make(map[string]interface{})
			m["index"] = i
			mb.App().EventDispatcher.SubmitEvent(&Event{"closed", mb, m})
			return true
		})
	}

	return mb
}
