package gowtk




type Context interface{}

type ViewSpec func(ctx View)

type TextView struct {
	Content StringProperty
	Component
}

func (t *TextView) Attach() {

}

func (t *TextView) Children() []View {
	return nil
}

func (t *TextView) SetContent(p string) *TextView {
	t.Content.SetString(p)
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

func (t *Button) Attach() {

}

func (t *Button) Children() []View {
	return nil
}

func (t *Button) SetCaption(str string) *Button {
	t.Content.SetString(str)
	return t
}

func (c *Button) OnClick(func()) *Button {
	return c
}

// nearly "overloading", as go can do that
func (c *TextView) OnClick(f func()) *TextView {
	//c.Component.OnClick(f)
	return c
}

// nearly "overloading", as go can do that
func (c *TextView) SetPadding(l, t, r, b int) *TextView {
	//c.Component.SetPadding(l, t, r, b)
	return c
}
