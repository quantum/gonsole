package main

import (
	"github.com/jojomi/gonsole"
	"github.com/nsf/termbox-go"
)

func main() {
	app := gonsole.NewApp()
	app.CloseKey = termbox.KeyEsc
	//app.CloseKey = 'q'
	win := gonsole.NewWindow("winMain")

	panel := gonsole.NewPanel("panel1")
	panel.Position = gonsole.Box{4, 8, 50, 22}
	panel.Title = "Test Controls"
	//panel.TitleAlignment =
	//panel.Background = termbox.ColorWhite
	panel.SetBorder(gonsole.LineDouble)
	win.AddControl(panel)
	//win.Background = termbox.ColorBlue

	ctrl := gonsole.NewLabel("lblStatus")
	ctrl.Position = gonsole.Box{2, 2, 30, 3}
	ctrl.Text = "Test"
	ctrl.SetBorder(gonsole.LineSingle)
	ctrl.Margin = gonsole.Sides{0, 1, 0, 1}
	win.AddControl(ctrl)

	ctrlChk := gonsole.NewCheckbox("chkActive")
	ctrlChk.Position = gonsole.Box{2, 2, 30, 3}
	ctrlChk.Checked = true
	ctrlChk.Text = "Test"
	ctrlChk.SetBorder(gonsole.LineDouble)
	panel.AddControl(ctrlChk)

	ctrlChk2 := gonsole.NewCheckbox("chkActive2")
	ctrlChk2.Position = gonsole.Box{2, 7, 30, 3}
	ctrlChk2.Checked = false
	ctrlChk2.Text = "Test with more text"
	panel.AddControl(ctrlChk2)

	ctrlBtn := gonsole.NewButton("MyButton")
	ctrlBtn.Position = gonsole.Box{2, 10, 40, 3}
	ctrlBtn.Text = "This is a button. Push me!"
	ctrlBtn.SetBorder(gonsole.LineSingle)
	panel.AddControl(ctrlBtn)

	ctrlBtn2 := gonsole.NewButton("MyButton2")
	ctrlBtn2.Position = gonsole.Box{2, 14, 40, 3}
	ctrlBtn2.Text = "This is my second magic button..."
	ctrlBtn2.SetBorder(gonsole.LineSingle)
	panel.AddControl(ctrlBtn2)

	ctrlBtn.Focus()

	app.AddWindow(win)

	// events
	ctrlBtn.AddEventListener("click", func(ev *gonsole.Event) bool {
		ctrlBtn.Text = "--- clicked ---"
		win.Pollute()
		//win.Repaint()
		win.FullRepaint()
		return true
	})

	app.Run()
}
