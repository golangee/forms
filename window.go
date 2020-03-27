package wtk

import "github.com/worldiety/wtk/dom"

var Root = newWindow()

type Window struct {
	window dom.Window
}

func newWindow() Window {
	return Window{window: dom.GetWindow()}
}

func (w Window) RemoveAll() {
	w.window.Document().Body().SetInnerHTML("")
}

func (w Window) AddView(v View) {
	v.attach(w.window.Document().Body())
}
