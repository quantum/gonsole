package gonsole

import "github.com/nsf/termbox-go"

// App holds the global gonsole state.
type App struct {
	CloseKey     termbox.Key
	windows      []AppWindow
	theme        *Theme
	channel      chan string
	globalEvents map[termbox.Key][]func()
}

// NewApp creates a new app
func NewApp() *App {
	app := &App{
		theme:        defaultTheme,
		channel:      make(chan string),
		globalEvents: make(map[termbox.Key][]func(), 0),
	}
	return app
}

func (app *App) ID() string {
	return "__global__app"
}

func (app *App) Repaint() {
	termbox.Clear(app.Theme().ColorTermbox("app.fg"), app.Theme().ColorTermbox("app.bg"))

	dirty := false
	for _, window := range app.windows {
		if window.Enabled() && window.Dirty() {
			window.Repaint()
			dirty = true
		}
	}

	aw := app.activeWindow()
	var fc Control
	if aw != nil {
		fc = aw.FocusedControl()
	}

	if aw == nil || fc == nil || !fc.Cursorable() {
		termbox.HideCursor()
	}

	if dirty {
		termbox.Flush()
	}
}

func (app *App) putEvent(eventType string) {
	go func() {
		app.channel <- eventType
	}()
}

func (app *App) Redraw() {
	app.putEvent("redraw")
}

func (app *App) Stop() {
	app.putEvent("quit")
}

func (app *App) Theme() *Theme {
	return app.theme
}

func (app *App) SetTheme(theme *Theme) {
	app.theme = theme
}

func (app *App) activeWindow() AppWindow {
	for i := len(app.windows) - 1; i >= 0; i-- {
		if app.windows[i].Enabled() {
			return app.windows[i]
		}
	}
	return nil
}

func (app *App) moveWindowToTop(win AppWindow) {
	if app.removeWindow(win) {
		app.addWindow(win)
	}
}

func (app *App) addWindow(win AppWindow) {
	app.windows = append(app.windows, win)
	app.Redraw()
}

func (app *App) removeWindow(win AppWindow) bool {
	for i, w := range app.windows {
		if w.ID() == win.ID() {
			if i < len(app.windows)-1 {
				app.windows = append(app.windows[:i], app.windows[i+1:]...)
			} else {
				app.windows = app.windows[:i]
			}
			app.Redraw()
			return true
		}
	}
	return false
}

func (app *App) parseGlobalEvent(ev *termbox.Event) bool {
	if ev.Type == termbox.EventKey {
		if funcs, ok := app.globalEvents[ev.Key]; ok {
			for _, function := range funcs {
				function()
			}
		}
	}

	return false
}

func (app *App) AddEventListener(key termbox.Key, handler func()) {
	funcArray, ok := app.globalEvents[key]
	if !ok {
		funcArray = make([]func(), 0)
	}
	funcArray = append(funcArray, handler)
	app.globalEvents[key] = funcArray
}

func (app *App) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	app.Repaint()

	for {
		select {
		case ev := <-eventQueue:
			switch ev.Type {
			case termbox.EventKey:
				if app.CloseKey == ev.Key {
					return
				}

				handled := false
				repaint := false
				activeWindow := app.activeWindow()
				if activeWindow != nil {
					handled, repaint = activeWindow.ParseEvent(&ev)
				}

				if !handled {
					app.parseGlobalEvent(&ev)
				} else if repaint {
					app.Repaint()
				}

			case termbox.EventResize:
				app.Repaint()
			case termbox.EventInterrupt:
				return
			case termbox.EventError:
				panic(ev.Err)
			}

		case ev := <-app.channel:
			if ev == "redraw" {
				app.Repaint()
			} else if ev == "quit" {
				return
			}
		}
	}
}
