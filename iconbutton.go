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
	"github.com/golangee/forms/event"
	"github.com/golangee/forms/theme/material/icon"
	"github.com/golangee/forms/theme/material/js"
	js2 "syscall/js"
)

// An IconButton is a rounded button with an icon or a single character on it
type IconButton struct {
	Value icon.Icon
	*absComponent
	isIcon bool
}

// NewIconButton creates a new icon
func NewIconButton(str icon.Icon) *IconButton {
	t := &IconButton{}
	t.absComponent = newComponent(t, "button")
	t.isIcon = true
	t.node().AddClass("mdc-icon-button").AddClass("material-icons")
	t.node().Style().Set("border-radius", "50%") // bugfix so that background colors looks nice
	t.Set(str)

	ripple := js.Attach(js.Ripple, t.node())
	ripple.Unwrap().Set("unbounded", true)
	t.absComponent.addResource(ripple)

	return t
}

// Set updates the material icon
func (t *IconButton) Set(str icon.Icon) *IconButton {
	if !t.isIcon {
		t.isIcon = true
		t.node().AddClass("material-icons")
	}

	t.Value = str
	t.absComponent.elem.SetInnerText(string(str))
	return t
}

// SetChar violates the material spec and allows to set a unicode character codepoint instead of an icon
func (t *IconButton) SetChar(r rune) *IconButton {
	t.node().RemoveClass("material-icons")
	t.absComponent.elem.SetInnerText(string(r))
	return t
}

// Style applies generic style attributes.
func (t *IconButton) Style(style ...Style) *IconButton {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *IconButton) Self(ref **IconButton) *IconButton {
	*ref = t
	return t
}

// AddClickListener registers another click listener
func (t *IconButton) AddClickListener(f func(v View)) *IconButton {
	t.addEventListener(event.Click, func(v View, params []js2.Value) {
		f(v)
	})
	return t
}

