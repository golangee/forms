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
func Elem(name string, mods ...Modifier) ChildHolder {
	elem := CreateElement(name)
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
