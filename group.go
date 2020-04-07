package wtk

type Group struct {
	*absComponent
}

func NewGroup() *Group {
	t := &Group{}
	t.absComponent = newComponent(t, "div")
	return t
}

func (t *Group) AddViews(views ...View) *Group {
	for _, v := range views {
		t.addView(v)
	}
	return t
}

func (t *Group) Style(style ...Style) *Group {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Group) Self(ref **Group) *Group {
	*ref = t
	return t
}
