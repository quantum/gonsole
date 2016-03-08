package gonsole

import "github.com/nsf/termbox-go"

// Control is the base model for a UI control
type BasicControl struct {
	window     *Window
	parent     Control
	id         string
	focussable bool
	dirty      bool

	Position Position
	Visible  bool
	Enabled  bool
	ZIndex   int
	TabIndex int

	Style      Style
	FocusStyle Style
}

func (ctrl *BasicControl) Init(id string) {
	ctrl.SetID(id)
	ctrl.Pollute()
}

func (ctrl *BasicControl) GetAbsolutePosition() Box {
	if parent := ctrl.Parent(); parent != nil {
		parentBox := parent.ContentBox()
		return ctrl.Position.Box(parentBox.Width, parentBox.Height).Absolute(parentBox)
	}
	w, h := termbox.Size()
	return ctrl.Position.Box(w, h)
}

func (ctrl *BasicControl) BorderBox() Box {
	style := ctrl.GetStyle()
	return ctrl.GetAbsolutePosition().Minus(style.Margin)
}

func (ctrl *BasicControl) ContentBox() Box {
	// substract padding and margin
	style := ctrl.GetStyle()
	contentBox := ctrl.GetAbsolutePosition().Minus(style.Margin).Minus(style.Padding)
	// substract border if applicable
	if style.Border != LineNone {
		contentBox = contentBox.Minus(Sides{1, 1, 1, 1})
	}
	return contentBox
}

func (ctrl *BasicControl) GetStyle() Style {
	if ctrl.Focussed() {
		return ctrl.FocusStyle
	}
	return ctrl.Style
}

func (ctrl *BasicControl) DrawBorder() {
	style := ctrl.GetStyle()

	if style.Border == LineNone {
		return
	}

	DrawBorder(ctrl.BorderBox(), style.Border, style.BorderFg, style.BorderBg)
}

func (ctrl *BasicControl) ParseEvent(ev *termbox.Event) bool {
	// to be implemented for individual controls
	return false
}

func (ctrl *BasicControl) ID() string {
	return ctrl.id
}

func (ctrl *BasicControl) SetID(id string) {
	ctrl.id = id
}

func (ctrl *BasicControl) Dirty() bool {
	return ctrl.dirty
}

func (ctrl *BasicControl) Pollute() {
	ctrl.dirty = true
}

func (ctrl *BasicControl) SetWindow(win *Window) {
	ctrl.window = win
}

func (ctrl *BasicControl) AddEventListener(eventType string, handler func(ev *Event) bool) {
	ctrl.Window().App.EventDispatcher.AddEventListener(ctrl, eventType, handler)
}

func (ctrl *BasicControl) SubmitEvent(ev *Event) {
	ctrl.Window().App.EventDispatcher.SubmitEvent(ev)
}

func (ctrl *BasicControl) Repaint() {
	if !ctrl.Dirty() {
		return
	}
	ClearRect(ctrl.BorderBox(), ColorDefault, ColorDefault)

	style := ctrl.GetStyle()

	if style.Bg != 0 {
		FillRect(ctrl.GetAbsolutePosition(), style.Fg, style.Bg)
	}
	ctrl.DrawBorder()
	// implement details in controls
}

func (ctrl *BasicControl) Focussed() bool {
	if ctrl.Window() != nil {
		return ctrl.Window().FocussedControl().ID() == ctrl.ID()
	}
	return false
}

func (ctrl *BasicControl) Focus() {
	ctrl.Window().SetFocussedControl(ctrl)
}

func (ctrl *BasicControl) Focussable() bool {
	return ctrl.focussable
}

func (ctrl *BasicControl) SetFocussable(focussable bool) {
	ctrl.focussable = focussable
}

func (ctrl *BasicControl) Parent() Control {
	return ctrl.parent
}

func (ctrl *BasicControl) SetParent(parent Control) {
	ctrl.parent = parent
}

func (ctrl *BasicControl) Window() *Window {
	if win := ctrl.window; win != nil {
		return win
	}
	if parent := ctrl.Parent(); parent != nil {
		return parent.Window()
	}
	return nil
}
