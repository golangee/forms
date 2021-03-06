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
	"github.com/golangee/forms/event"
	"github.com/golangee/forms/property"
	"github.com/golangee/forms/theme/material/icon"
	js2 "github.com/golangee/forms/theme/material/js"
	"strconv"
	"syscall/js"
)

type InputType string

const Password InputType = "password"
const Txt InputType = "text"
const EMail InputType = "email"
const URL InputType = "url"
const Range InputType = "range"
const Number InputType = "number"
const Date InputType = "date"
const DateTime InputType = "datetime"
const DateTimeLocal InputType = "datetime-local"
const Month InputType = "month"
const Search InputType = "search"
const Tel InputType = "tel"
const Time InputType = "time"
const Week InputType = "week"

type TextField struct {
	*absComponent
	label          h.Element
	labelNotch     h.Element
	helperLine     h.Element
	helperText     h.Element
	leadingIco     h.Element
	trailingIco    h.Element
	characterCount h.Element
	mdcTextField   h.Element
	input          h.Element
	fndTF          js2.Foundation
	dirty          bool
	textProperty   property.String
	intProperty    property.Int
}

func NewTextField() *TextField {
	t := &TextField{}
	t.absComponent = newComponent(t, "div")

	t.textProperty = property.NewString()
	t.textProperty.Observe(func(old, new string) {
		t.SetText(new)
	})

	t.intProperty = property.NewInt()
	t.intProperty.Observe(func(old, new int) {
		t.SetText(strconv.Itoa(new))
	})

	labelId := nextId()
	h.Wrap(t.node(), h.Class("text-field-container"),
		h.Div(h.Class("mdc-text-field", "mdc-text-field--outlined"), h.Class("mdc-text-field--no-label"),
			h.I(h.Class("material-icons", "mdc-text-field__icon")).Self(&t.leadingIco),
			h.I(h.Class("material-icons", "mdc-text-field__icon")).Self(&t.trailingIco),
			h.Input(h.Class("mdc-text-field__input"), h.Type("text"), h.AriaLabelledby(labelId)).Self(&t.input),
			h.Div(h.Class("mdc-notched-outline"),
				h.Div(h.Class("mdc-notched-outline__leading")),
				h.Div(h.Class("mdc-notched-outline__notch"),
					h.Span(h.Class("mdc-floating-label"), h.Id(labelId)).Self(&t.label),
				).Self(&t.labelNotch),
				h.Div(h.Class("mdc-notched-outline__trailing")),
			),
		).Self(&t.mdcTextField),
		h.Div(h.Class("mdc-text-field-helper-line"),
			h.Div(h.Class("mdc-text-field-helper-text", "mdc-text-field-helper-text--persistent", "mdc-text-field-helper-text--validation-msg")).Self(&t.helperText),

			h.Div(h.Class()).Self(&t.characterCount),
		).Self(&t.helperLine),
	)
	t.addResource(t.input.AddEventListener("change", func(this js.Value, args []js.Value) interface{} {
		v := t.Text()

		t.textProperty.Set(v)
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			t.intProperty.Set(int(i))
		}

		return nil
	}, false))
	t.helperLine.Style().Set("display", "none")
	t.labelNotch.Style().Set("display", "none")
	t.invalidate(false)
	return t
}

// invalidate postpones foundation construction until view is attached
func (t *TextField) invalidate(force bool) {
	if t.isAttached() || force {
		t.dirty = false
		t.initFoundation()
	} else {
		t.dirty = true
	}
}

func (t *TextField) initFoundation() {
	t.fndTF.Release()
	t.fndTF = js2.Attach(js2.TextField, t.mdcTextField)
}

func (t *TextField) SetLeadingIcon(ico icon.Icon) *TextField {
	t.mdcTextField.RemoveClass("mdc-text-field--with-leading-icon")
	if len(ico) > 0 {
		t.mdcTextField.AddClass("mdc-text-field--with-leading-icon")
		t.leadingIco.Style().Set("display", "block")
	} else {
		t.leadingIco.Style().Set("display", "none")
	}
	t.leadingIco.SetText(string(ico))
	return t
}

func (t *TextField) AddTrailingIconClickListener(f func(v View)) *TextField {
	t.trailingIco.SetTabIndex(1)
	t.trailingIco.SetRole("button")
	t.absComponent.addResource(t.trailingIco.AddEventListener(string(event.Click), func(this js.Value, args []js.Value) interface{} {
		f(t)
		return nil
	}, false))

	return t
}

func (t *TextField) SetTrailingIcon(ico icon.Icon) *TextField {
	t.mdcTextField.RemoveClass("mdc-text-field--with-trailing-icon")
	if len(ico) > 0 {
		t.mdcTextField.AddClass("mdc-text-field--with-trailing-icon")
		t.trailingIco.Style().Set("display", "block")
	} else {
		t.trailingIco.Style().Set("display", "none")
	}
	t.trailingIco.SetText(string(ico))
	return t
}

func (t *TextField) AddLeadingIconClickListener(f func(v View)) *TextField {
	t.leadingIco.SetTabIndex(0)
	t.leadingIco.SetRole("button")
	t.absComponent.addResource(t.leadingIco.AddEventListener(string(event.Click), func(this js.Value, args []js.Value) interface{} {
		f(t)
		return nil
	}, false))

	return t
}

func (t *TextField) SetEnabled(b bool) *TextField {
	t.input.SetDisabled(!b)
	t.mdcTextField.RemoveClass("field--disabled")
	if !b {
		t.mdcTextField.AddClass("field--disabled")
	}
	t.invalidate(false)
	return t
}

// SetText updates the text
func (t *TextField) SetText(str string) *TextField {
	t.input.Unwrap().Set("value", str)
	t.invalidate(false)
	return t
}

// Text returns the current text
func (t *TextField) Text() string {
	return t.input.Unwrap().Get("value").String()
}

// TextProperty returns a text property to set or get the text.
func (t *TextField) TextProperty() property.String {
	return t.textProperty
}

// IntProperty returns an int property to set or get the text with integer to string conversions.
func (t *TextField) IntProperty() property.Int {
	return t.intProperty
}

// BindText is a shortcut for TextProperty().Bind() and returning self. Initially the value is read
// from the pointers position and the component is populated. Afterwards only the value at the pointer
// location is updated by the component.
func (t *TextField) BindText(s *string) *TextField {
	t.TextProperty().Bind(s)
	return t
}

// BindInt is a shortcut for IntProperty().Bind() and returning self. Initially the value is read
// from the pointers position and the component is populated. Afterwards only the value at the pointer
// location is updated by the component.
//
// Text values which cannot be parsed as integers, are ignored.
func (t *TextField) BindInt(s *int) *TextField {
	t.IntProperty().Bind(s)
	return t
}

func (t *TextField) SetLabel(str string) *TextField {
	t.mdcTextField.RemoveClass("mdc-text-field--no-label")
	if len(str) == 0 {
		t.mdcTextField.AddClass("mdc-text-field--no-label")
		t.labelNotch.Style().Set("display", "none")
	} else {
		t.labelNotch.Style().Set("display", "block")
	}
	t.label.SetText(str)
	t.invalidate(false)
	return t
}

// Styles changes the container
func (t *TextField) Style(style ...Style) *TextField {
	t.absComponent.style(style...)
	t.inputStyle(style...)
	return t
}

// InputStyle changes the internal field style.
func (t *TextField) inputStyle(styles ...Style) *TextField {
	for _, s := range styles {
		s.applyCSS(t.mdcTextField)
		_ = s
	}
	return t
}

func (t *TextField) SetInputType(in InputType) *TextField {
	t.input.SetAttr("type", string(in))
	return t
}

func (t *TextField) SetRange(min, max int) *TextField {
	t.SetInputType(Range)
	t.input.SetAttr("min", min)
	t.input.SetAttr("max", max)
	return t
}

func (t *TextField) SetHelper(str string) *TextField {
	t.helperText.SetTextContent(str)
	if str == "" {
		t.helperLine.Style().Set("display", "none")
	} else {
		t.helperLine.Style().Set("display", "block")
	}
	return t
}

func (t *TextField) SetHelperPersistent(b bool) *TextField {
	t.helperText.RemoveClass("mdc-text-field-helper-text--persistent")
	if b {
		t.helperText.AddClass("mdc-text-field-helper-text--persistent")
	}
	return t
}

func (t *TextField) SetMaxLength(chars int) *TextField {
	t.characterCount.SetClassName("mdc-text-field-character-counter")
	t.input.SetAttr("maxLength", chars)
	t.invalidate(false)
	return t
}

func (t *TextField) SetInvalid(b bool) *TextField {
	t.SetHelperPersistent(b)
	t.mdcTextField.AddClass("mdc-text-field--invalid")
	if !t.fndTF.IsValid() {
		t.invalidate(true)
	}
	t.fndTF.Unwrap().Set("useNativeValidation", !b)
	t.fndTF.Unwrap().Set("valid", !b)
	return t
}

// Self assigns the receiver to the given reference
func (t *TextField) Self(ref **TextField) *TextField {
	*ref = t
	return t
}

func (t *TextField) SetRequired(b bool) *TextField {
	if b {
		t.input.Unwrap().Set("required", "required")
	} else {
		t.input.Unwrap().Delete("required")
	}
	return t
}

func (t *TextField) attach(v View) {
	t.absComponent.attach(v)
	if t.dirty {
		t.dirty = false
		t.invalidate(true)
	}
}

func (t *TextField) Release() {
	t.fndTF.Release()
	t.absComponent.Release()
}

// AddClickListener registers another click listener
func (t *TextField) AddClickListener(f func(v View)) *TextField {
	t.addEventListener(event.Click, func(v View, params []js.Value) {
		f(v)
	})
	return t
}

// AddKeyUpListener registers another key listener
func (t *TextField) AddKeyUpListener(f func(v View, keyCode int)) *TextField {
	t.addEventListener(event.KeyUp, func(v View, params []js.Value) {
		f(v, params[0].Get("keyCode").Int())
	})
	return t
}

// AddKeyDownListener registers another key listener
func (t *TextField) AddKeyDownListener(f func(v View, keyCode int)) *TextField {
	t.addEventListener(event.KeyDown, func(v View, params []js.Value) {
		f(v, params[0].Get("keyCode").Int())
	})
	return t
}

// AddFocusOutListener registers another key listener
func (t *TextField) AddFocusChangeListener(f func(v View, hasFocus bool)) *TextField {
	t.addEventListener(event.FocusOut, func(v View, params []js.Value) {
		f(v, false)
	})

	t.addEventListener(event.FocusIn, func(v View, params []js.Value) {
		f(v, true)
	})
	return t
}
