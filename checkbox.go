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
	js2 "github.com/golangee/forms/theme/material/js"
	"syscall/js"
)

type Checkbox struct {
	label       dom.Element
	input       dom.Element
	fndCheckbox js2.Foundation
	*absComponent
}

func NewCheckbox() *Checkbox {
	t := &Checkbox{}
	t.absComponent = newComponent(t, "div") // we need another root, which can be a block element instead of flex, otherwise the label alignment breaks
	formFieldRoot := dom.CreateElement("div").SetClassName("mdc-form-field mdc-checkbox--touch")

	mdcCheckbox := dom.CreateElement("div").AddClass("mdc-checkbox")
	t.input = dom.CreateElement("input").SetType("checkbox").SetClassName("mdc-checkbox__native-control").SetId(nextId())

	mdcCheckbox.AppendChild(t.input)
	mdcCheckbox.AppendChild(createPath())
	mdcCheckbox.AppendChild(dom.CreateElement("div").SetClassName("mdc-checkbox__ripple"))

	formFieldRoot.AppendChild(mdcCheckbox)
	t.label = dom.CreateElement("label").SetFor(t.input.Id())
	formFieldRoot.AppendChild(t.label)
	t.node().AppendChild(formFieldRoot)

	t.fndCheckbox = js2.Attach(js2.Checkbox, mdcCheckbox)
	fndFormField := js2.Attach(js2.FormField, t.node())
	fndFormField.Unwrap().Set("input", t.fndCheckbox.Unwrap())

	t.addResource(t.fndCheckbox)
	t.addResource(fndFormField)

	return t
}

func (t *Checkbox) AddChangeListener(f func(v *Checkbox)) *Checkbox {
	t.addResource(t.node().AddEventListener("change", func(this js.Value, args []js.Value) interface{} {
		f(t)
		return nil
	}, true))
	return t
}

func (t *Checkbox) SetEnabled(b bool) *Checkbox {
	t.fndCheckbox.Unwrap().Set("disabled", !b)
	return t
}

func (t *Checkbox) SetIndeterminate(b bool) *Checkbox {
	t.fndCheckbox.Unwrap().Set("indeterminate", b)
	return t
}

func (t *Checkbox) Indeterminate() bool {
	return t.fndCheckbox.Unwrap().Get("indeterminate").Bool()
}

func (t *Checkbox) SetChecked(b bool) *Checkbox {
	t.fndCheckbox.Unwrap().Set("checked", b)
	return t
}

func (t *Checkbox) Checked() bool {
	return t.fndCheckbox.Unwrap().Get("checked").Bool()
}

func (t *Checkbox) SetText(s string) *Checkbox {
	t.label.SetTextContent(s)
	return t
}

func (t *Checkbox) Style(style ...Style) *Checkbox {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Checkbox) Self(ref **Checkbox) *Checkbox {
	*ref = t
	return t
}

func createPath() dom.Element {
	nsSVG := "http://www.w3.org/2000/svg"
	bg := dom.CreateElement("div").AddClass("mdc-checkbox__background")
	svg := dom.CreateElementNS(nsSVG, "svg").AddClass("mdc-checkbox__checkmark")
	svg.SetAttr("viewBox", "0 0 24 24")
	path := dom.CreateElementNS(nsSVG, "path").AddClass("mdc-checkbox__checkmark-path")
	path.SetAttr("fill", "none")
	path.SetAttr("d", "M1.73,12.91 8.1,19.28 22.79,4.59")

	svg.AppendChild(path)
	bg.AppendChild(svg)

	bg.AppendChild(dom.CreateElement("div").SetClassName("mdc-checkbox__mixedmark"))
	return bg
}
