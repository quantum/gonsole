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
	channel         chan Event
}

// NewApp creates a new app
func NewApp() *App {
	app := &App{
		eventDispatcher: NewEventDispatcher(),
		theme:           defaultTheme,
		channel:         make(chan Event),
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
		app.channel <- Event{Type: eventType, Source: app}
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

func (app *App) addWindow(win AppWindow) {
	app.windows = append(app.windows, win)
}

func (app *App) removeWindow(win AppWindow) bool {
	for i, w := range app.windows {
		if w.ID() == win.ID() {
			if i < len(app.windows)-1 {
				app.windows = append(app.windows[:i], app.windows[i+1:]...)
			} else {
				app.windows = app.windows[:i]
			}
			app.eventDispatcher.RemoveEventListener(win)
			return true
		}
	}
	return false
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
				activeWindow := app.activeWindow()
				if activeWindow != nil {
					handled = activeWindow.ParseEvent(&ev)
				}

				if !handled {
					app.parseGlobalEvent(&ev)
				}
			case termbox.EventResize:
				app.Repaint()
			case termbox.EventInterrupt:
				return
			case termbox.EventError:
				panic(ev.Err)
			}

		case ev := <-app.channel:
			if ev.Type == "redraw" {
				app.Repaint()
			} else if ev.Type == "quit" {
				return
			}
		}
	}
}
