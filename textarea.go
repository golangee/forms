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
	"github.com/golangee/forms/dom"
	"github.com/golangee/forms/theme/material/js"
)

type TextArea struct {
	*absComponent
	layoutCtr *taDefault
}

func NewTextArea() *TextArea {
	t := &TextArea{}
	t.absComponent = newComponent(t, "div")
	t.layoutCtr = newTaDefault(t.node())
	t.absComponent.addResource(t.layoutCtr)
	return t
}

func (t *TextArea) SetEnabled(b bool) *TextArea {
	t.layoutCtr.setEnabled(b)
	return t
}

func (t *TextArea) SetText(str string) *TextArea {
	t.layoutCtr.setInput(str)
	return t
}

func (t *TextArea) SetLabel(str string) *TextArea {
	t.layoutCtr.setLabel(str)
	return t
}

// Styles changes the container
func (t *TextArea) Style(style ...Style) *TextArea {
	t.absComponent.style(style...)
	return t
}

// Styles changes the
func (t *TextArea) InputStyle(styles ...Style) *TextArea {
	for _, s := range styles {
		s.applyCSS(t.layoutCtr.mdcTextField())
	}
	return t
}

func (t *TextArea) SetHelper(str string) *TextArea {
	t.layoutCtr.setHelper(str)
	return t
}

func (t *TextArea) SetMaxLength(chars int) *TextArea {
	t.layoutCtr.setMaxLength(chars)
	return t
}

/*
func (t *TextArea) SetRows(r int) *TextArea {
	t.layoutCtr.mdcTextField().Unwrap().Set("row", r)
	return t
}

func (t *TextArea) SetColumns(c int) *TextArea {
	t.layoutCtr.mdcTextField().Unwrap().Set("cols", c)
	return t
}*/

func (t *TextArea) SetInvalid(b bool) *TextArea {
	t.layoutCtr.foundation.Unwrap().Set("useNativeValidation", !b)
	t.layoutCtr.foundation.Unwrap().Set("valid", !b)
	if b {
		t.layoutCtr.mdcTextField().AddClass("mdc-text-field--invalid")
	} else {
		t.layoutCtr.mdcTextField().RemoveClass("mdc-text-field--invalid")
	}
	return t
}

// Self assigns the receiver to the given reference
func (t *TextArea) Self(ref **TextArea) *TextArea {
	*ref = t
	return t
}

func (t *TextArea) SetRequired(b bool) *TextArea {
	if b {
		t.layoutCtr.inputField().Unwrap().Set("required", "required")
	} else {
		t.layoutCtr.mdcTextField().Unwrap().Delete("required")
	}
	return t
}

type textAreaLayoutController interface {
	setInput(val string)
	setLabel(val string)
	setEnabled(b bool)
	setHelper(str string)
	setMaxLength(c int)
	mdcTextField() dom.Element
	inputField() dom.Element
	Release()
}

//  <label class="mdc-text-field mdc-text-field--textarea">
//  <div class="mdc-text-field-character-counter">0 / 140</div>
//  <textarea class="mdc-text-field__input" aria-labelledby="my-label-id" rows="8" cols="40" maxlength="140"></textarea>
//  <div class="mdc-notched-outline">
//    <div class="mdc-notched-outline__leading"></div>
//    <div class="mdc-notched-outline__notch">
//      <span class="mdc-floating-label" id="my-label-id">Textarea Label</span>
//    </div>
//    <div class="mdc-notched-outline__trailing"></div>
//  </div>
//  </label>
type taDefault struct {
	container    dom.Element
	div          dom.Element
	input        dom.Element
	outline      dom.Element
	outlineNotch dom.Element
	label        dom.Element
	helperDiv    dom.Element
	helperText   dom.Element
	valueLabel   string
	valueInput   string
	valueHelper  string
	maxLen       int
	foundation   js.Foundation
	fndContainer js.Foundation
}

func newTaDefault(parentDiv dom.Element) *taDefault {
	t := &taDefault{container: parentDiv}

	t.div = dom.CreateElement("div")
	t.input = dom.CreateElement("textarea").AddClass("mdc-text-field__input")
	t.input.SetId(nextId())
	t.label = dom.CreateElement("label").SetFor(t.input.Id()).AddClass("mdc-floating-label")

	t.outline = dom.CreateElement("div").AddClass("mdc-notched-outline")
	t.outline.AppendChild(dom.CreateElement("div").AddClass("mdc-notched-outline__leading"))
	t.outlineNotch = dom.CreateElement("div").AddClass("mdc-notched-outline__notch")
	t.outline.AppendChild(t.outlineNotch)
	t.outline.AppendChild(dom.CreateElement("div").AddClass("mdc-notched-outline__trailing"))

	t.helperText = dom.CreateElement("div").SetClassName("mdc-text-field-helper-text mdc-text-field-helper-text--persistent mdc-text-field-helper-text--validation-msg")
	t.helperDiv = dom.CreateElement("div").AddClass("mdc-text-field-helper-line")
	t.helperDiv.AppendChild(t.helperText)

	t.apply()
	return t
}

func (t *taDefault) apply() {
	t.foundation.Release()

	t.container.SetClassName("text-field-container")
	t.container.SetTextContent("")
	t.div.SetClassName("mdc-text-field mdc-text-field--textarea")

	t.setInput(t.valueInput)
	t.div.AppendChild(t.input)

	t.outlineNotch.SetTextContent("")
	if len(t.valueLabel) > 0 {
		t.label.SetTextContent(t.valueLabel)
		t.outlineNotch.AppendChild(t.label)
	} else {
		t.div.AddClass("mdc-text-field--no-label")
	}

	t.container.AppendChild(t.div)

	t.div.AppendChild(t.outline)

	if len(t.valueHelper) > 0 {
		t.helperText.SetTextContent(t.valueHelper)
		t.container.AppendChild(t.helperDiv)
	}

	if t.maxLen != 0 {
		t.input.Unwrap().Set("maxLength", t.maxLen)
		t.container.AppendChild(dom.CreateElement("div").SetClassName("mdc-text-field-character-counter"))
	}

	t.foundation = js.Attach(js.TextField, t.div)         // otherwise the focus border does not work
	t.fndContainer = js.Attach(js.TextField, t.container) // otherwise the character count does not work
}

// setInput fixes FOUC label
func (t *taDefault) setInput(val string) {
	t.label.SetClassName("mdc-floating-label")
	if len(val) > 0 {
		t.label.AddClass("mdc-floating-label--float-above")
	}
	t.input.Unwrap().Set("value", val)
	t.valueInput = val
}

func (t *taDefault) setLabel(val string) {
	t.valueLabel = val
	t.apply()
}

func (t *taDefault) Release() {
	t.foundation.Release()
	t.fndContainer.Release()
}

func (t *taDefault) setEnabled(b bool) {
	t.input.SetDisabled(!b)
	if b {
		t.div.RemoveClass("mdc-text-field--disabled")
	} else {
		t.div.AddClass("mdc-text-field--disabled")
	}
}

func (t *taDefault) setHelper(str string) {
	t.valueHelper = str
	t.apply()
}

func (t *taDefault) mdcTextField() dom.Element {
	return t.div
}

func (t *taDefault) inputField() dom.Element {
	return t.input
}

func (t *taDefault) setMaxLength(c int) {
	t.maxLen = c
	t.apply()
}
