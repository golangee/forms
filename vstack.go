package wtk

type VStack struct {
	*absComponent
}

func NewVStack() *VStack {
	t := &VStack{}
	t.absComponent = newComponent(t, "div")
	return t
}

func (t *VStack) AddViews(views ...View) *VStack {
	for _, v := range views {
		t.addView(v)
	}
	return t
}

func (t *VStack) Style(style ...Style) *VStack {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *VStack) Self(ref **VStack) *VStack {
	*ref = t
	return t
}
