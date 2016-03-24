package gonsole

import "github.com/nsf/termbox-go"

// Window is the top-level struct in gonsole library.
type BaseWindow struct {
	BaseElement
	BaseContainer

	app            *App
	focusedControl Control
	onClose        func()
}

func (win *BaseWindow) Init(app *App, id string) {
	win.app = app
	win.BaseElement.Init(win, nil, id, "window")
}

func (win *BaseWindow) App() *App {
	return win.app
}

func (win *BaseWindow) Focused() bool {
	aw := win.app.activeWindow()
	if aw != nil {
		if win.ID() == aw.ID() {
			return true
		}
	}
	return false
}

func (win *BaseWindow) Focus() {
	win.app.moveWindowToTop(win)
}

func (win *BaseWindow) FocusedControl() Control {
	return win.focusedControl
}

func (win *BaseWindow) FocusControl(control Control) {
	for _, loopFC := range win.ChildrenDeep() {
		if loopFC.ID() == control.ID() {
			win.focusedControl = loopFC
			return
		}
	}
}

func (win *BaseWindow) Close() {
	win.App().removeWindow(win)
	win.App().Redraw()
}

func (win *BaseWindow) OnClose(handler func()) {
	win.onClose = handler
}

func (win *BaseWindow) moveFocus(num int) {
	focusControls := win.getFocusableControls()
	if len(focusControls) == 0 {
		return
	}

	currentFocusControl := win.FocusedControl()
	// get focus index
	index := -1
	for i, loopFC := range focusControls {
		if loopFC.ID() == currentFocusControl.ID() {
			index = i
		}
	}
	newIndex := (index + num + len(focusControls)) % len(focusControls)
	if index == -1 {
		newIndex = 0
	}
	newFocusControl := focusControls[newIndex]
	win.FocusControl(newFocusControl)

	// update focus, mark dirty
	currentFocusControl.SetDirty(true)
	newFocusControl.SetDirty(true)
}

func (win *BaseWindow) getFocusableControls() []Control {
	// TODO order by tabIndex and filter non-focussable controls
	focusControls := []Control{}
	for _, control := range win.ChildrenDeep() {
		if control.Focusable() {
			focusControls = append(focusControls, control)
		}
	}
	return focusControls
}

// return true if event was parsed and should not continue bubbling up
func (win *BaseWindow) ParseEvent(ev *termbox.Event) (handled, repaint bool) {
	// TODO window level event parsing, support tabbing for changing focus

	// dispatch event to currently focussed control
	if win.FocusedControl() != nil {
		handled, repaint = win.FocusedControl().ParseEvent(ev)
		if handled {
			return handled, repaint
		}
	}

	// focus navigation events
	// catch tab key if the focussed control did not need it
	// catch arrow keys if the focussed control did not need them
	if ev.Type == termbox.EventKey {
		switch ev.Key {
		case termbox.KeyTab:
			win.moveFocus(1)
			return true, true
		case termbox.KeyArrowDown, termbox.KeyArrowRight:
			win.moveFocus(1)
			return true, true
		case termbox.KeyArrowUp, termbox.KeyArrowLeft:
			win.moveFocus(-1)
			return true, true
		}
	}

	return false, false
}

func (win *BaseWindow) Repaint() {
	if !win.Dirty() {
		return
	}

	prefix := ""
	if win.Focused() {
		prefix = "focused."
	}

	win.drawBox(prefix)
	win.BaseContainer.RepaintChildren()

	// draw title
	if win.Title() != "" {
		if win.BorderType() == LineNone {
			win.SetPadding(win.Padding().Plus(Sides{Top: 1}))
		}
		fg, bg := win.Theme().ColorTermbox(prefix+"title.fg"), win.Theme().ColorTermbox(prefix+"title.bg")
		DrawTextSimple(" "+win.Title()+" ", false, win.BorderBox().Minus(Sides{Left: 2}), fg, bg)
	}
}
