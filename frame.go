package wtk

type Frame struct {
	v View
	*absComponent
}

func NewFrame() *Frame {
	t := &Frame{}
	t.absComponent = newComponent(t, "div")
	/* //this is center inside, but any flex-grow will not work inside it
	t.node().Style().Set("margin", "0")
	t.node().Style().Set("position", "absolute")
	t.node().Style().Set("top", "50%")
	t.node().Style().Set("left", "50%")
	t.node().Style().Set("transform", "translate(-50%,-50%)")
	*/

	t.node().Style().Set("max-width", "600")
	t.node().Style().Set("margin-left", "auto")
	t.node().Style().Set("margin-right", "auto")
	return t
}

func (t *Frame) SetView(v View) *Frame {
	if t.v != nil {
		t.absComponent.removeView(t.v)
		t.v = nil
	}
	t.absComponent.addView(v)
	t.v = v
	return t
}

func (t *Frame) Style(style ...Style) *Frame {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Frame) Self(ref **Frame) *Frame {
	*ref = t
	return t
}
