package gonsole

import "github.com/nsf/termbox-go"

// App holds the global gonsole state.
type App struct {
	CloseKey        termbox.Key
	EventDispatcher *EventDispatcher
	windows         []AppWindow
	activeWindow    AppWindow
	theme           *Theme
}

// NewApp creates a new app
func NewApp() *App {
	app := &App{
		EventDispatcher: NewEventDispatcher(),
		theme:           defaultTheme,
	}
	return app
}

func (app *App) Repaint() {
	dirty := false
	for _, window := range app.windows {
		if window.Dirty() {
			window.Repaint()
			dirty = true
		}
	}

	if dirty {
		termbox.Flush()
	}
}

func (app *App) Stop() {
	termbox.Interrupt()
}

func (app *App) Theme() *Theme {
	return app.theme
}

func (app *App) SetTheme(theme *Theme) {
	app.theme = theme
}

func (app *App) ActivateWindow(win AppWindow) {
	app.activeWindow = win
}

func (app *App) addWindow(win AppWindow) {
	app.windows = append(app.windows, win)

	// first window is automatically activated
	if len(app.windows) == 1 {
		app.ActivateWindow(win)
	}
}

func (app *App) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	termbox.Clear(app.Theme().ColorTermbox("app.fg"), app.Theme().ColorTermbox("app.bg"))
	app.Repaint()

mainloop:
	for {
		// poll events
		ev := termbox.PollEvent()
		switch ev.Type {
		case termbox.EventKey:
			if app.CloseKey == ev.Key {
				break mainloop
			}
		case termbox.EventInterrupt:
			break mainloop
		case termbox.EventError:
			panic(ev.Err)
		}

		// dispatch event to active window
		app.activeWindow.ParseEvent(&ev)

		app.Repaint()
	}
}
