package wtk

import "github.com/worldiety/wtk/event"

type Button struct {
	Text string
	*absComponent
}

func NewButton(text string) *Button {
	t := &Button{}
	t.absComponent = newComponent(t, "button")
	t.SetText(text)
	return t
}

func (t *Button) SetText(str string) *Button {
	t.Text = str
	t.absComponent.elem.SetInnerText(str)
	return t
}

func (t *Button) Style(style ...Style) *Button {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Button) Self(ref **Button) *Button {
	*ref = t
	return t
}

func (t *Button) AddClickListener(f func(v View)) *Button {
	t.addEventListener(event.Click, f)
	return t
}
