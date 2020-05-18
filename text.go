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

type Text struct {
	Value string
	*absComponent
}

func NewText(str string) *Text {
	t := &Text{}
	t.absComponent = newComponent(t, "span")
	t.Set(str)
	return t
}

func (t *Text) Set(str string) *Text {
	t.Value = str
	t.absComponent.elem.SetInnerText(str)
	return t
}

func (t *Text) Style(style ...Style) *Text {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Text) Self(ref **Text) *Text {
	*ref = t
	return t
}
