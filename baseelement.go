package gonsole

import "github.com/nsf/termbox-go"

type BaseElement struct {
	window   AppWindow
	parent   Container
	id       string
	dirty    bool
	enabled  bool
	position Position
	margin   Sides
	padding  Sides
	theme    *Theme
}

func (e *BaseElement) Init(window AppWindow, parent Container, id, themePrefix string) {
	e.window = window
	e.parent = parent
	e.id = id
	e.enabled = true
	e.dirty = true
	e.theme = NewTheme(themePrefix, e.window.App().Theme())
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
	return e.Theme().Color("fg"), e.Theme().Color("bg")
}

func (e *BaseElement) SetColors(fg Attribute, bg Attribute) {
	e.Theme().SetColor("fg", fg)
	e.Theme().SetColor("bg", bg)
}

func (e *BaseElement) FocusColors() (fg Attribute, bg Attribute) {
	return e.Theme().Color("focus.fg"), e.Theme().Color("focus.fg")
}

func (e *BaseElement) SetFocusColors(fg Attribute, bg Attribute) {
	e.Theme().SetColor("focus.fg", fg)
	e.Theme().SetColor("focus.bg", bg)
}

func (e *BaseElement) BorderType() LineType {
	return e.Theme().Border("border")
}

func (e *BaseElement) SetBorderType(border LineType) {
	e.Theme().SetBorder("border", border)
}

func (e *BaseElement) BorderColors() (fg Attribute, bg Attribute) {
	return e.Theme().Color("border.fg"), e.Theme().Color("border.bg")
}

func (e *BaseElement) SetBorderColors(fg Attribute, bg Attribute) {
	e.Theme().SetColor("border.fg", fg)
	e.Theme().SetColor("border.bg", bg)
}

func (e *BaseElement) ShadowType() LineType {
	return e.Theme().Border("shadow")
}

func (e *BaseElement) SetShadowType(shadow LineType) {
	e.Theme().SetBorder("shadow", shadow)
}

func (e *BaseElement) ShadowColor() Attribute {
	return e.Theme().Color("shadow.fg")
}

func (e *BaseElement) SetShadowColor(color Attribute) {
	e.Theme().SetColor("shadow.fg", color)
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
	borderBox := e.AbsolutePosition().Minus(e.margin)

	if e.ShadowType() != LineNone {
		borderBox = borderBox.Minus(Sides{0, 1, 1, 0})
	}

	return borderBox
}

func (e *BaseElement) ContentBox() Box {
	// substract padding and margin
	contentBox := e.AbsolutePosition().Minus(e.margin).Minus(e.padding)
	// substract border if applicable
	if e.BorderType() != LineNone {
		contentBox = contentBox.Minus(Sides{1, 1, 1, 1})
	}
	if e.ShadowType() != LineNone {
		contentBox = contentBox.Minus(Sides{0, 1, 1, 0})
	}
	return contentBox
}

func (e *BaseElement) AddEventListener(eventType string, handler func(ev *Event) bool) {
	e.window.App().eventDispatcher.AddEventListener(e, eventType, handler)
}

func (e *BaseElement) SubmitEvent(ev *Event) {
	e.window.App().eventDispatcher.SubmitEvent(ev)
}

func (e *BaseElement) ParseEvent(ev *termbox.Event) bool {
	return false
}

func (e *BaseElement) drawBox(status string) {
	t := e.Theme()

	FillRect(e.BorderBox(), t.ColorTermbox(status+"fg"), t.ColorTermbox(status+"bg"))

	border := e.BorderType()
	if border != LineNone {
		DrawBorder(e.BorderBox(), border, t.ColorTermbox(status+"border.fg"), t.ColorTermbox(status+"border.bg"))
	}

	shadow := e.ShadowType()
	if shadow != LineNone {
		DrawShadow(e.AbsolutePosition(), t.ColorTermbox(status+"shadow.fg"))
	}
}

func (e *BaseElement) Theme() *Theme {
	return e.theme
}
