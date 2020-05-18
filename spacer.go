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

type Spacer struct {
	*absComponent
}

func NewSpacer() *Spacer {
	t := &Spacer{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("flex-grow", "1")
	return t
}

func (t *Spacer) Style(style ...Style) *Spacer {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Spacer) Self(ref **Spacer) *Spacer {
	*ref = t
	return t
}
