package wtk

import (
	h "github.com/worldiety/wtk/dom"
	js2 "github.com/worldiety/wtk/theme/material/js"
	"strconv"
	"syscall/js"
)

type Picker struct {
	*absComponent
	menu           h.Element
	label          h.Element
	helper         h.Element
	fnd            js2.Foundation
	selectListener func(v *Picker)
}

func NewPicker() *Picker {
	t := &Picker{}
	t.absComponent = newComponent(t, "div")
	t.node().SetClassName("mdc-select mdc-select--outlined")

	labelId := nextId()
	h.Wrap(t.node(),
		h.Div(h.Class("mdc-select__anchor"),
			h.I(h.Class("mdc-select__dropdown-icon")),
			h.Div(h.Class("mdc-select__selected-text"), h.AriaLabelledby(labelId), h.Id(nextId())),
			h.Div(h.Class("mdc-notched-outline"),
				h.Div(h.Class("mdc-notched-outline__leading")),
				h.Div(h.Class("mdc-notched-outline__notch"),
					h.Span(h.Id(labelId), h.Class("mdc-floating-label")).Self(&t.label),
				),
				h.Div(h.Class("mdc-notched-outline__trailing")),
			),
		),
		h.Div(h.Class("mdc-select__menu", "mdc-menu", "mdc-menu-surface"), h.Role("listbox")).Self(&t.menu),
		h.Div(h.Class("mdc-text-field-helper-line"),
			h.P(h.Class("mdc-text-field-helper-text", "mdc-text-field-helper-text--persistent", "mdc-text-field-helper-text--validation-msg")).Self(&t.helper),
		),
	)

	t.fnd = js2.Attach(js2.Select, t.node())
	return t
}

func (t *Picker) SetLabel(str string) *Picker {
	t.label.SetTextContent(str)
	return t
}

func (t *Picker) SetHelper(str string) *Picker {
	t.helper.SetTextContent(str)
	return t
}

func (t *Picker) SetInvalid(b bool) *Picker {
	if b {
		t.node().AddClass("mdc-text-field--invalid")
	} else {
		t.node().RemoveClass("mdc-text-field--invalid")
	}
	return t
}

func (t *Picker) SetOptions(options ...string) *Picker {
	t.fnd.Release()
	t.menu.SetTextContent("")
	ul := h.CreateElement("ul").AddClass("mdc-list")
	for i, opt := range options {
		li := h.CreateElement("li").AddClass("mdc-list-item").SetAttr("data-value", strconv.Itoa(i)).SetRole("option")
		li.SetTextContent(opt)
		ul.AppendChild(li)
	}
	t.menu.AppendChild(ul)

	t.fnd = js2.Attach(js2.Select, t.node())
	t.addResource(t.node().AddEventListener("MDCSelect:change", func(this js.Value, args []js.Value) interface{} {
		if t.selectListener != nil {
			t.selectListener(t)
		}
		return nil
	}, false))
	return t
}

func (t *Picker) SetSelectListener(f func(v *Picker)) *Picker {
	t.selectListener = f
	return t
}

func (t *Picker) Style(style ...Style) *Picker {
	t.absComponent.style(style...)
	return t
}

func (t *Picker) Selected() int {
	return t.fnd.Unwrap().Get("selectedIndex").Int()
}

func (t *Picker) SetSelected(idx int) *Picker {
	t.fnd.Unwrap().Set("selectedIndex", idx)
	return t
}

func (t *Picker) SetEnabled(b bool) *Picker {
	t.fnd.Unwrap().Set("disabled", !b)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Picker) Self(ref **Picker) *Picker {
	*ref = t
	return t
}

func (t *Picker) Release() {
	t.fnd.Release()
	t.absComponent.Release()
}
