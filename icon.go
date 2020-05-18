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

import "github.com/golangee/forms/theme/material/icon"

type Icon struct {
	Value icon.Icon
	*absComponent
}

func NewIcon(str icon.Icon) *Icon {
	t := &Icon{}
	t.absComponent = newComponent(t, "i")
	t.node().AddClass("material-icons")
	t.Set(str)
	return t
}

func (t *Icon) Set(str icon.Icon) *Icon {
	t.Value = str
	t.absComponent.elem.SetInnerText(string(str))
	return t
}

func (t *Icon) Style(style ...Style) *Icon {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Icon) Self(ref **Icon) *Icon {
	*ref = t
	return t
}

