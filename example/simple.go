package main

import (
	"github.com/nsf/termbox-go"
	g "github.com/quantum/gonsole"
)

func main() {
	app := g.NewApp()
	app.CloseKey = termbox.KeyEsc
	//app.CloseKey = 'q'
	win := g.NewWindow("winMain")

	panel := g.NewPanel("panel1")
	panel.Position = g.Position{"25%", "25%", "50%", "50%"}
	panel.Style = g.Style{Border: g.LineDashed}
	panel.FocusStyle = g.Style{Border: g.LineDashed, BorderFg: g.ColorYellow}
	panel.Title = "Test Controls"
	//panel.TitleAlignment =
	//panel.Background = termbox.ColorWhite
	win.AddControl(panel)
	//win.Background = termbox.ColorBlue

	ctrl := g.NewLabel("lblStatus")
	ctrl.Position = g.Position{"2", "2", "30", "3"}
	ctrl.Style = g.Style{Border: g.LineSingle, Margin: g.Sides{0, 1, 0, 1}}
	ctrl.FocusStyle = g.Style{Border: g.LineSingle, Margin: g.Sides{0, 1, 0, 1}, BorderFg: g.ColorYellow}
	ctrl.Text = "Test"
	win.AddControl(ctrl)

	ctrlChk := g.NewCheckbox("chkActive")
	ctrlChk.Position = g.Position{"2", "2", "30", "3"}
	ctrlChk.Style = g.Style{Border: g.LineDouble}
	ctrlChk.FocusStyle = g.Style{Border: g.LineDouble, BorderFg: g.ColorYellow}
	ctrlChk.Checked = true
	ctrlChk.Text = "Test"
	panel.AddControl(ctrlChk)

	ctrlChk2 := g.NewCheckbox("chkActive2")
	ctrlChk2.Position = g.Position{"2", "7", "30", "3"}
	ctrlChk2.Checked = false
	ctrlChk2.Text = "Test with more text"
	panel.AddControl(ctrlChk2)

	ctrlBtn := g.NewButton("MyButton")
	ctrlBtn.Position = g.Position{"2", "10", "40", "3"}
	ctrlBtn.Style = g.Style{Border: g.LineSingle}
	ctrlBtn.FocusStyle = g.Style{Border: g.LineSingle, BorderFg: g.ColorYellow}
	ctrlBtn.Text = "This is a button. Push me!"
	panel.AddControl(ctrlBtn)

	ctrlBtn2 := g.NewButton("MyButton2")
	ctrlBtn2.Position = g.Position{"2", "14", "40", "3"}
	ctrlBtn2.Style = g.Style{Border: g.LineSingle}
	ctrlBtn2.FocusStyle = g.Style{Border: g.LineSingle, BorderFg: g.ColorYellow}
	ctrlBtn2.Text = "This is my second magic button..."
	panel.AddControl(ctrlBtn2)

	ctrlChk2.Focus()

	app.AddWindow(win)

	// events
	ctrlBtn.AddEventListener("clicked", func(ev *g.Event) bool {
		ctrlBtn.Text = "--- clicked ---"
		return true
	})

	ctrlBtn2.AddEventListener("clicked", func(ev *g.Event) bool {
		btn := ev.Source.(*g.Button)
		btn.Text = "Clicked button"
		return true
	})

	ctrlChk2.AddEventListener("checked", func(ev *g.Event) bool {
		chk := ev.Source.(*g.Checkbox)
		chk.Text = "works"
		return true
	})

	ctrlChk2.AddEventListener("unchecked", func(ev *g.Event) bool {
		chk := ev.Source.(*g.Checkbox)
		chk.Text = "does not work"
		return true
	})

	app.Run()
}
