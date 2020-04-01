package wtk

import "github.com/worldiety/wtk/dom"

var Root = newWindow()

type Window struct {
	window dom.Window
}

func (w Window) attach(parent View) {
}

func (w Window) detach() {
}

func (w Window) parent() View {
	return nil
}

func (w Window) node() dom.Element {
	return w.window.Document().Body()
}

func (w Window) Release() {
}

func newWindow() Window {
	return Window{window: dom.GetWindow()}
}

func (w Window) RemoveAll() {
	w.window.Document().Body().SetInnerHTML("")
}

func (w Window) AddView(v View) {
	v.attach(w)
	w.node().AppendChild(v.node())
}
