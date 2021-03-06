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
	"github.com/golangee/forms/event"
	"github.com/golangee/forms/theme/material/icon"
	"github.com/golangee/forms/theme/material/js"
	js2 "syscall/js"
)

type ButtonStyleKind string

const Default ButtonStyleKind = "mdc-button"
const Raised ButtonStyleKind = "mdc-button--raised"
const Unelevated ButtonStyleKind = "mdc-button--unelevated"
const Outlined ButtonStyleKind = "mdc-button--outlined"
const Dlg ButtonStyleKind = "mdc-dialog__button"

type Alignment string

const Leading Alignment = "leading"
const Trailing Alignment = "trailing"

type Button struct {
	Text  string
	label dom.Element
	*absComponent
	leadingIcon dom.Element
}

func NewButton(text string) *Button {
	t := &Button{}
	t.absComponent = newComponent(t, "button")
	t.node().AddClass("mdc-button")
	t.node().AppendChild(dom.CreateElement("div").AddClass("mdc-button__ripple"))
	t.label = dom.CreateElement("span").AddClass("mdc-button__label")

	t.node().AppendChild(t.label)

	t.SetText(text)
	t.absComponent.addResource(js.Attach(js.Ripple, t.node()))
	return t
}

func (t *Button) AddIcon(icon icon.Icon, alignment Alignment) *Button {
	switch alignment {
	case Leading:
		if !t.leadingIcon.IsValid() {
			t.leadingIcon = dom.CreateElement("i").SetClassName("material-icons mdc-button__icon").SetInnerText(string(icon))
			t.node().InsertBefore(t.leadingIcon, t.label)
		}
		t.leadingIcon.SetInnerText(string(icon))
	case Trailing:
		t.node().AppendChild(dom.CreateElement("i").SetClassName("material-icons mdc-button__icon").SetInnerText(string(icon)))
	default:
		panic("unsupported alignment: " + alignment)
	}
	return t
}

func (t *Button) SetLeadingIcon(icon icon.Icon) *Button {
	t.AddIcon(icon, Leading)
	return t
}

func (t *Button) SetStyleKind(s ButtonStyleKind) *Button {
	t.node().SetClassName("mdc-button " + string(s))
	return t
}

func (t *Button) SetEnabled(b bool) *Button {
	t.node().SetDisabled(!b)
	return t
}

func (t *Button) SetText(str string) *Button {
	t.Text = str
	t.label.SetInnerText(str)
	return t
}

func (t *Button) Style(style ...Style) *Button {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Button) Self(ref **Button) *Button {
	*ref = t
	return t
}

// AddClickListener registers another click listener
func (t *Button) AddClickListener(f func(v View)) *Button {
	t.addEventListener(event.Click, func(v View, params []js2.Value) {
		f(v)
	})
	return t
}
