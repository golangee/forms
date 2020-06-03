// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package forms

import (
	h "github.com/golangee/forms/dom"
	js2 "github.com/golangee/forms/theme/material/js"
	"strconv"
	"syscall/js"
)

// Picker is also known as Combobox, Dropdown or Spinner.
type Picker struct {
	*absComponent
	menu           h.Element
	label          h.Element
	helper         h.Element
	fnd            js2.Foundation
	selectListener func(v *Picker)
	myOptions      []string
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
	t.myOptions = options
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

func (t *Picker) SetSelectedString(str string) *Picker {
	for idx, v := range t.myOptions {
		if v == str {
			t.SetSelected(idx)
			break
		}
	}
	return t
}

func (t *Picker) SelectedString() string {
	idx := t.Selected()
	for i, v := range t.myOptions {
		if i == idx {
			return v
		}
	}

	return ""
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
