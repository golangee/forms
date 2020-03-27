package wtk

import (
	"github.com/worldiety/wtk/dom"
	"syscall/js"
)

type Button struct {
	model Model
	btn   dom.Element
	f     Handle
	Text  string
}

func (t *Button) init() {
	if t.f.Valid() {
		return
	}
	t.btn = dom.CreateElement("button")
	t.btn.SetInnerText(t.Text)
	t.f = t.model.AddListener(func(v interface{}) {
		t.btn.SetInnerHTML(t.Text)
	})

}

func (t *Button) attach(parent dom.Element) {
	t.init()
	parent.AppendChild(t.btn)
}

func (t *Button) detach(parent dom.Element) {
	parent.RemoveChild(t.btn)
}

func (t *Button) AddOnClickListener(f func()) {
	t.init()
	_ = t.btn.AddEventListener("click", func(value js.Value) {
		f()
	}, false)
}
