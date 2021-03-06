package gonsole

import "github.com/nsf/termbox-go"

// An UI element. It has a position, margin, padding, colors and other style properties.
type Element interface {
	GetWindow() AppWindow
	Parent() Container
	ID() string

	Dirty() bool
	SetDirty(dirty bool)

	Enabled() bool
	SetEnabled(enabled bool)

	Focused() bool
	Focus()

	Position() Position
	SetPosition(pos Position)

	Margin() Sides
	SetMargin(margins Sides)

	Padding() Sides
	SetPadding(margins Sides)

	Colors() (fg Attribute, bg Attribute)
	SetColors(fg Attribute, bg Attribute)

	FocusColors() (fg Attribute, bg Attribute)
	SetFocusColors(fg Attribute, bg Attribute)

	BorderType() LineType
	SetBorderType(border LineType)

	BorderColors() (fg Attribute, bg Attribute)
	SetBorderColors(fg Attribute, bg Attribute)

	AbsolutePosition() Box
	BorderBox() Box
	ContentBox() Box

	ParseEvent(ev *termbox.Event) (handled, repaint bool)

	Repaint()

	Theme() *Theme
}

// An element that is a container for controls
type Container interface {
	Element

	Title() string
	SetTitle(title string)

	AddControl(control Control)

	Children() []Control
	ChildrenDeep() []Control
}

// A control which is an element that can optional be focused
type Control interface {
	Element

	Focusable() bool
	SetFocusable(active bool)

	Cursorable() bool
	SetCursorable(cursorable bool)
}

// A window which is a top level container for controls.
// TODO: a better name for this interface. ideally "window"
type AppWindow interface {
	Container

	App() *App
	Close()

	FocusedControl() Control
	FocusControl(control Control)
}
