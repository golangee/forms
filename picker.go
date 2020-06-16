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
	"github.com/golangee/forms/property"
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
	selectAnchor   h.Element
	textProperty   property.String
}

// NewPicker creates a new Combobox/Dropdown/Spinner component.
func NewPicker(options ...string) *Picker {
	t := &Picker{}
	t.absComponent = newComponent(t, "div")
	t.node().SetClassName("mdc-select mdc-select--outlined")
	t.textProperty = property.NewString()

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
		).Self(&t.selectAnchor),
		h.Div(h.Class("mdc-select__menu", "mdc-menu", "mdc-menu-surface"), h.Role("listbox")).Self(&t.menu),
		h.Div(h.Class("mdc-text-field-helper-line"),
			h.P(h.Class("mdc-text-field-helper-text", "mdc-text-field-helper-text--persistent", "mdc-text-field-helper-text--validation-msg")).Self(&t.helper),
		),
	)

	t.fnd = js2.Attach(js2.Select, t.node())
	if len(options) > 0 {
		t.SetOptions(options...)
	}

	t.textProperty.Observe(func(old, new string) {
		t.SetText(new)
	})

	return t
}

// SetText tries to find the text in "options" and selects the index, if possible. Otherwise does nothing.
func (t *Picker) SetText(str string) *Picker {
	for idx, v := range t.myOptions {
		if v == str {
			t.SetSelected(idx)
			break
		}
	}

	return t
}

// Text returns the current selected option text. If nothing is selected returns the empty string.
func (t *Picker) Text() string {
	return t.selectedString()
}

// TextProperty returns a text property to or get the text.
func (t *Picker) TextProperty() property.String {
	return t.textProperty
}

// BindText is a shortcut for TextProperty().Bind() and returning self.
// If used together with #SetSelected() the order matters. When populated from/to a model,
// you likely want to first select a default (by index) and then bind to the model, which reads the value
// from the given pointer. However if that is invalid, it does nothing, otherwise selects the right index by name.
func (t *Picker) BindText(s *string) *Picker {
	t.TextProperty().Bind(s)
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
		t.textProperty.Set(t.selectedString())
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
	for _, s := range style {
		s.applyCSS(t.selectAnchor)
	}
	return t
}

func (t *Picker) Selected() int {
	return t.fnd.Unwrap().Get("selectedIndex").Int()
}

func (t *Picker) SetSelected(idx int) *Picker {
	if t.Selected() == idx {
		return t // debounce
	}

	t.fnd.Unwrap().Set("selectedIndex", idx)
	return t
}

func (t *Picker) selectedString() string {
	idx := t.Selected()
	if idx >= 0 && idx < len(t.myOptions) {
		return t.myOptions[idx]
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
