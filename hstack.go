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

type HStack struct {
	*absComponent
}

func NewHStack() *HStack {
	t := &HStack{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("display", "flex")
	return t
}

func (t *HStack) AddViews(views ...View) *HStack {
	for _, v := range views {
		t.addView(v)
	}
	return t
}

func (t *HStack) Style(style ...Style) *HStack {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *HStack) Self(ref **HStack) *HStack {
	*ref = t
	return t
}
