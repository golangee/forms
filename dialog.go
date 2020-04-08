package wtk

import (
	"github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/theme/material/js"
	"log"
	js2 "syscall/js"
)

type Dialog struct {
	Value string
	*absComponent
	title  dom.Element
	body   dom.Element
	fnd    js.Foundation
	footer dom.Element
}

func NewDialog() *Dialog {
	t := &Dialog{}
	t.absComponent = newComponent(t, "div")
	t.node().AddClass("mdc-dialog")
	dlgContainer := dom.CreateElement("div").AddClass("mdc-dialog__container")
	t.node().AppendChild(dlgContainer)

	dlgSurface := dom.CreateElement("div").AddClass("mdc-dialog__surface")
	dlgContainer.AppendChild(dlgSurface)

	t.title = dom.CreateElement("h2").AddClass("mdc-dialog__title")
	dlgSurface.AppendChild(t.title)

	t.body = dom.CreateElement("div").AddClass("mdc-dialog__content")
	dlgSurface.AppendChild(t.body)

	t.footer = dom.CreateElement("footer").AddClass("mdc-dialog__actions")
	dlgSurface.AppendChild(t.footer)

	dlgSurface.AppendChild(dom.CreateElement("div").AddClass("mdc-dialog__scrim"))

	t.fnd = js.Attach(js.Dialog, t.node())
	t.addResource(t.fnd)
	return t
}

func (t *Dialog) Show(parent View) {
	wnd := getWindow(parent)
	if wnd == nil {
		log.Println("cannot show dialog, view is not attached")
		return
	}
	wnd.AddView(t)
	var closedFunc dom.Func
	closedFunc = t.node().AddEventListener("MDCDialog:closed", func(this js2.Value, args []js2.Value) interface{} {
		t.destroy(parent)
		closedFunc.Release()
		return nil
	}, true)
	t.addResource(closedFunc)
	t.fnd.Unwrap().Call("open")
}

func (t *Dialog) destroy(parent View) {
	wnd := getWindow(parent)
	if wnd == nil {
		log.Println("cannot show dialog, view is not attached")
		return
	}
	wnd.RemoveView(t)
	t.Release()
}

func (t *Dialog) SetBody(v View) *Dialog {
	//hm, breaking our type-system, no attach here
	t.body.SetTextContent("")
	t.body.AppendChild(v.node())
	t.addResource(v)
	return t
}

func (t *Dialog) AddAction(caption string, onClick func(dlg *Dialog)) *Dialog {
	btn := NewButton(caption).SetStyleKind(Dlg)
	btn.AddClickListener(func(v View) {
		onClick(t)
	})
	//hm, breaking our type-system, no attach here
	t.footer.AppendChild(btn.node())
	t.addResource(btn)
	return t
}

func (t *Dialog) Close() {
	t.fnd.Unwrap().Call("close")
}

func (t *Dialog) SetTitle(s string) *Dialog {
	t.title.SetTextContent(s)
	return t
}

func (t *Dialog) Style(style ...Style) *Dialog {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Dialog) Self(ref **Dialog) *Dialog {
	*ref = t
	return t
}
