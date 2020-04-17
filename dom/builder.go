package dom

type Modifier interface {
	Apply(e Element)
}

type funcMod func(e Element)

func (f funcMod) Apply(e Element) {
	f(e)
}

type ChildHolder struct {
	Element Element
	f       func(e Element)
}

func (f ChildHolder) Self(e *Element) ChildHolder {
	*e = f.Element
	return f
}

func (f ChildHolder) Build(e Element) Element {
	f.f(e)
	return f.Element
}

func (f ChildHolder) Apply(e Element) {
	f.f(e)
}

func Class(c ...string) Modifier {
	return funcMod(func(e Element) {
		for _, s := range c {
			e.AddClass(s)
		}
	})
}

func ViewBox(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("viewBox", v)
	})
}

func Type(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("type", v)
	})
}


func Cx(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("cx", v)
	})
}

func Cy(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("cy", v)
	})
}

func R(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("r", v)
	})
}

func Fill(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("fill", v)
	})
}

func StrokeWidth(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("stroke-width", v)
	})
}

func StrokeMiterlimit(v string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("stroke-miterlimit", v)
	})
}

func Elem(name string, mods ...Modifier) ChildHolder {
	var elem Element
	switch name {
	case "svg":
		elem = CreateElementNS("http://www.w3.org/2000/svg", "svg")
	case "circle":
		elem = CreateElementNS("http://www.w3.org/2000/svg", "circle")
	default:
		elem = CreateElement(name)
	}

	return ChildHolder{
		Element: elem,
		f: func(e Element) {
			for _, m := range mods {
				m.Apply(elem)
			}
			e.AppendChild(elem)
		},
	}
}

func Wrap(elem Element, mods ...Modifier) {
	for _, m := range mods {
		m.Apply(elem)
	}
}

func Div(mods ...Modifier) ChildHolder {
	return Elem("div", mods...)
}

func Input(mods ...Modifier) ChildHolder {
	return Elem("input", mods...)
}

func Button(mods ...Modifier) ChildHolder {
	return Elem("button", mods...)
}

func Svg(mods ...Modifier) ChildHolder {
	return Elem("svg", mods...)
}
func Circle(mods ...Modifier) ChildHolder {
	return Elem("circle", mods...)
}

func P(mods ...Modifier) ChildHolder {
	return Elem("p", mods...)
}
func Span(mods ...Modifier) ChildHolder {
	return Elem("div", mods...)
}

func I(mods ...Modifier) ChildHolder {
	return Elem("i", mods...)
}

func AriaLabelledby(label string) Modifier {
	return funcMod(func(e Element) {
		e.SetAttr("aria-labelledby", label)
	})
}

func TextContent(r string) Modifier {
	return funcMod(func(e Element) {
		e.SetTextContent(r)
	})
}

func Role(r string) Modifier {
	return funcMod(func(e Element) {
		e.SetRole(r)
	})
}

func Id(id string) Modifier {
	return funcMod(func(e Element) {
		e.SetId(id)
	})
}
