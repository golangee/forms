package wtk

type HStack struct {
	*absComponent
}

func NewHStack() *HStack {
	t := &HStack{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("display", "flex")
	return t
}

func (t *HStack) AddViews(views ...View) *HStack {
	for _, v := range views {
		t.addView(v)
	}
	return t
}

func (t *HStack) Style(style ...Style) *HStack {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *HStack) Self(ref **HStack) *HStack {
	*ref = t
	return t
}
