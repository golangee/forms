package gowtk

type View interface {
	Children() []View
}

type Component struct {
	//SetPadding(l, t, r, b int) Component
	//OnClick(func())Component
}

func (c *Component) OnClick(func()) *Component {
	return c
}

func (c *Component) SetPadding(l, t, r, b int) *Component {
	return c
}

type Context interface{}

type ViewSpec func(ctx View)

type TextView struct {
	Content StringProperty
	Component
}

func (t *TextView) Children() []View {
	return nil
}

func (t *TextView) SetContent(p string) *TextView {
	return t
}

func (t *TextView) SetFontSize(p int) *TextView {
	return t
}

func NewTextView() *TextView {
	return &TextView{}
}

type Button struct {
	Content StringProperty
	Component
}

func (t *Button) Children() []View {
	return nil
}

func (t *Button) SetCaption(str string) *Button {
	return t
}

func (c *Button) OnClick(func()) *Button {
	return c
}

// nearly "overloading", as go can do that
func (c *TextView) OnClick(f func()) *TextView {
	c.Component.OnClick(f)
	return c
}

// nearly "overloading", as go can do that
func (c *TextView) SetPadding(l, t, r, b int) *TextView {
	c.Component.SetPadding(l, t, r, b)
	return c
}

type VBox struct {
	children []View
}

func (b *VBox) Children() []View {
	return b.children
}

func (b *VBox) Layout(view ...View) *VBox {
	b.children = append(b.children, view...)
	return nil
}

func NewVBox(views ...View) *VBox {
	return &VBox{}
}
