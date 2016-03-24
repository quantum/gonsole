package main

import (
	"github.com/nsf/termbox-go"
	g "github.com/quantum/gonsole"
)

func openConfirmExit(app *g.App) {
	title := "Exit"
	message := "Are you sure you want to exit"

	d := g.NewMessageDialog(app, "exit", title, message, []string{"Yes", "No"})
	d.SetPosition(g.Position{"30%", "30%", "40%", "30%"})
	d.OnClose(func() {
		if d.SelectedButton() == 0 {
			app.Stop()
		}
	})
}

func openConfirmInput(app *g.App, value string) {
	title := "Are you sure?"
	message := "Are you sure you want " + value

	d := g.NewMessageDialog(app, "confirminput", title, message, []string{"Yes", "No"})
	d.SetPosition(g.Position{"30%", "30%", "40%", "30%"})
	d.OnClose(func() {
		if d.SelectedButton() == 0 {
			openSelectionDialog(app)
		} else {
			openInputDialog(app)
		}
	})
}

func openConfirmColor(app *g.App, value string) {
	title := "Are you sure?"
	message := "Are you sure you like the color " + value

	d := g.NewMessageDialog(app, "confirmcolor", title, message, []string{"Yes", "No"})
	d.SetPosition(g.Position{"30%", "30%", "40%", "30%"})
	d.OnClose(func() {
		if d.SelectedButton() == 0 {
			app.Stop()
		} else {
			openSelectionDialog(app)
		}
	})
}

func openSelectionDialog(app *g.App) {
	title := "Choose a color"
	message := "What is your favorite color?"

	options := []string{"Red", "Blue", "Green"}
	d := g.NewSelectionDialog(app, "selection", title, message, []string{"OK", "Cancel"}, options)
	d.SetPosition(g.Position{"15%", "25%", "70%", "40%"})
	d.OnClose(func() {
		if d.SelectedButton() == 0 {
			openConfirmColor(app, options[d.SelectedItem()])
		} else {
			app.Stop()
		}
	})
}

func openInputDialog(app *g.App) {
	title := "Enter Something"
	message := "What do you like?"

	d := g.NewInputDialog(app, "input", title, message, []string{"OK", "Cancel"})
	d.SetPosition(g.Position{"15%", "25%", "70%", "40%"})
	d.OnClose(func() {
		if d.SelectedButton() == 0 {
			openConfirmInput(app, d.InputValue())
		} else {
			app.Stop()
		}
	})
}

func main() {
	app := g.NewApp()
	app.CloseKey = termbox.KeyCtrlQ

	app.AddEventListener(termbox.KeyF10, func() {
		openConfirmExit(app)
	})

	infoBar := g.NewWindow(app, "info")
	infoBar.SetBorderType(g.LineNone)
	infoBar.SetShadowType(g.LineNone)
	infoBar.Theme().SetBorder("focused.border", g.LineNone)
	infoBar.SetPosition(g.Position{"0", "100%-1", "100%", "1"})

	infoBarLabel := g.NewLabel(infoBar, infoBar, "infobar")
	infoBarLabel.SetText("[Tab] to move, [Enter] to select, [F10] to exit")
	infoBarLabel.SetPosition(g.Position{"0", "0", "100%", "100%"})

	openInputDialog(app)

	app.Run()
}
