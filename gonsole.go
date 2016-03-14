package gonsole

import (
	"fmt"

	"github.com/nsf/termbox-go"
)

// App holds the global gonsole state.
type App struct {
	CloseKey        termbox.Key
	eventDispatcher *EventDispatcher
	windows         []AppWindow
	theme           *Theme
}

// NewApp creates a new app
func NewApp() *App {
	app := &App{
		eventDispatcher: NewEventDispatcher(),
		theme:           defaultTheme,
	}
	return app
}

func (app *App) ID() string {
	return "__global__app"
}

func (app *App) Repaint() {
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

func (app *App) Stop() {
	termbox.Interrupt()
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

func (app *App) addWindow(win AppWindow) {
	app.windows = append(app.windows, win)
}

func (app *App) removeWindow(win AppWindow) {
	for i, w := range app.windows {
		if w == win {
			if i < len(app.windows)-1 {
				app.windows = append(app.windows[:i], app.windows[i+1:]...)
			} else {
				app.windows = app.windows[:i]
			}
			app.eventDispatcher.RemoveEventListener(win)
			return
		}
	}
}

func (app *App) parseGlobalEvent(ev *termbox.Event) bool {
	if ev.Type == termbox.EventKey {
		key := app.eventDispatcher.getKey(app, fmt.Sprintf("%d", ev.Key))
		if funcs, ok := app.eventDispatcher.registeredEvents[key]; ok {
			for _, function := range funcs {
				function(&Event{"key", app, nil})
			}
		}
	}

	return false
}

func (app *App) AddEventListener(key termbox.Key, handler func(ev *Event) bool) {
	app.eventDispatcher.AddEventListener(app, fmt.Sprintf("%d", key), handler)
}

func (app *App) Run() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

mainloop:
	for {
		termbox.Clear(app.Theme().ColorTermbox("app.fg"), app.Theme().ColorTermbox("app.bg"))
		app.Repaint()

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

		handled := false
		activeWindow := app.activeWindow()
		if activeWindow != nil {
			handled = activeWindow.ParseEvent(&ev)
		}

		if !handled {
			app.parseGlobalEvent(&ev)
		}
	}
}
