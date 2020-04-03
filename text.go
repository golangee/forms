package wtk

type Text struct {
	Value string
	*absComponent
}

func NewText(str string) *Text {
	t := &Text{}
	t.absComponent = newComponent(t, "span")
	t.Set(str)
	return t
}

func (t *Text) Set(str string) *Text {
	t.Value = str
	t.absComponent.elem.SetInnerText(str)
	return t
}

func (t *Text) Style(style ...Style) *Text {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Text) Self(ref **Text) *Text {
	*ref = t
	return t
}
