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

func (f ChildHolder) Build() Element {
	f.f(CreateElement("div"))
	return f.Element
}

func (f ChildHolder) Apply(e Element) {
	f.f(e)
}

func Class(s string) Modifier {
	return funcMod(func(e Element) {
		e.AddClass(s)
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

func Div(mods ...Modifier) ChildHolder {
	return Elem("div", mods...)
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

func Id(id string) Modifier {
	return funcMod(func(e Element) {
		e.SetId(id)
	})
}
