package gonsole

import "github.com/nsf/termbox-go"

type BaseElement struct {
	window             AppWindow
	parent             Container
	id                 string
	dirty              bool
	enabled            bool
	position           Position
	margin             Sides
	padding            Sides
	fg, bg             Attribute
	fgFocus, bgFocus   Attribute
	border             LineType
	fgBorder, bgBorder Attribute
}

func (e *BaseElement) Init(window AppWindow, parent Container, id string) {
	e.window = window
	e.parent = parent
	e.id = id
	e.enabled = true
	e.dirty = true
}

func (e *BaseElement) GetWindow() AppWindow {
	return e.window
}

func (e *BaseElement) Parent() Container {
	return e.parent
}

func (e *BaseElement) ID() string {
	return e.id
}

func (e *BaseElement) Dirty() bool {
	return e.dirty
}

func (e *BaseElement) SetDirty(dirty bool) {
	e.dirty = dirty
}

func (e *BaseElement) Enabled() bool {
	return e.enabled
}

func (e *BaseElement) SetEnabled(enabled bool) {
	e.enabled = enabled
}

func (e *BaseElement) Position() Position {
	return e.position
}

func (e *BaseElement) SetPosition(pos Position) {
	e.position = pos
}

func (e *BaseElement) Margin() Sides {
	return e.margin
}

func (e *BaseElement) SetMargin(margins Sides) {
	e.margin = margins
}

func (e *BaseElement) Padding() Sides {
	return e.padding
}

func (e *BaseElement) SetPadding(paddings Sides) {
	e.padding = paddings
}

func (e *BaseElement) Colors() (fg Attribute, bg Attribute) {
	return e.fg, e.bg
}

func (e *BaseElement) SetColors(fg Attribute, bg Attribute) {
	e.fg = fg
	e.bg = bg
}

func (e *BaseElement) FocusColors() (fg Attribute, bg Attribute) {
	return e.fgFocus, e.bgFocus
}

func (e *BaseElement) SetFocusColors(fg Attribute, bg Attribute) {
	e.fgFocus = fg
	e.bgFocus = bg
}

func (e *BaseElement) BorderType() LineType {
	return e.border
}

func (e *BaseElement) SetBorderType(border LineType) {
	e.border = border
}

func (e *BaseElement) BorderColors() (fg Attribute, bg Attribute) {
	return e.fgBorder, e.bgBorder
}

func (e *BaseElement) SetBorderColors(fg Attribute, bg Attribute) {
	e.fgBorder = fg
	e.bgBorder = bg
}

func (e *BaseElement) AbsolutePosition() Box {
	if parent := e.Parent(); parent != nil {
		parentBox := parent.ContentBox()
		return e.position.Box(parentBox.Width, parentBox.Height).Absolute(parentBox)
	}
	w, h := termbox.Size()
	return e.position.Box(w, h)
}

func (e *BaseElement) BorderBox() Box {
	return e.AbsolutePosition().Minus(e.margin)
}

func (e *BaseElement) ContentBox() Box {
	// substract padding and margin
	contentBox := e.AbsolutePosition().Minus(e.margin).Minus(e.padding)
	// substract border if applicable
	if e.border != LineNone {
		contentBox = contentBox.Minus(Sides{1, 1, 1, 1})
	}
	return contentBox
}

func (e *BaseElement) AddEventListener(eventType string, handler func(ev *Event) bool) {
	e.window.App().EventDispatcher.AddEventListener(e, eventType, handler)
}

func (e *BaseElement) SubmitEvent(ev *Event) {
	e.window.App().EventDispatcher.SubmitEvent(ev)
}

func (e *BaseElement) ParseEvent(ev *termbox.Event) bool {
	return false
}

func (e *BaseElement) Repaint() {
	if !e.Dirty() {
		return
	}
	ClearRect(e.BorderBox(), ColorDefault, ColorDefault)

	if e.bg != ColorDefault {
		FillRect(e.AbsolutePosition(), e.fg, e.bg)
	}

	if e.border != LineNone {
		DrawBorder(e.BorderBox(), e.border, e.fgBorder, e.bgBorder)
	}

	// implement details in controls
}
