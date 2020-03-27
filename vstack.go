package wtk

import (
	"github.com/worldiety/wtk/dom"
)

type VStack struct {
	model Model
	div   dom.Element
	cb    Handle
	views []View
}

func (s *VStack) init() {
	if s.cb.Valid() {
		return
	}
	s.div = dom.CreateElement("div")
	s.cb = s.model.AddListener(func(model interface{}) {
		s.div.SetInnerText("")
		for _, view := range model.([]View) {
			view.attach(s.div)
		}
	})
}

func (s *VStack) attach(parent dom.Element) {
	s.init()
	parent.AppendChild(s.div)
}

func (s *VStack) detach(parent dom.Element) {
	parent.RemoveChild(s.div)
}

func (s *VStack) AddView(view View) {
	s.init()
	s.views = append(s.views, view)
	s.model.Value = s.views
	s.model.Notify()
}
