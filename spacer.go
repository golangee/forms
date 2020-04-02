package wtk

type Spacer struct {
	*absComponent
}

func NewSpacer() *Spacer {
	t := &Spacer{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("flex-grow", "1")
	return t
}

func (t *Spacer) Style(style ...Style) *Spacer {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Spacer) Self(ref **Spacer) *Spacer {
	*ref = t
	return t
}
