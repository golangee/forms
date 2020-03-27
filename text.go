package wtk

import (
	"github.com/worldiety/wtk/dom"
)

type Text struct {
	Value String
	p     dom.Element
	f     Func
}

func (t *Text) init() {
	if t.f.Valid() {
		return
	}
	t.p = dom.CreateElement("p")
	t.p.SetInnerText(t.Value.Get())
	t.f = t.Value.AddListener(func(old interface{}, new interface{}) {
		t.p.SetInnerHTML(t.Value.Get())
	})
}

func (t *Text) attach(parent dom.Element) {
	t.init()
	parent.AppendChild(t.p)
}

func (t *Text) detach(parent dom.Element) {
	parent.RemoveChild(t.p)
}
