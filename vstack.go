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

type VStack struct {
	*absComponent
}

func NewVStack(views ...View) *VStack {
	t := &VStack{}
	t.absComponent = newComponent(t, "div")
	t.AddViews(views...)
	return t
}

func (t *VStack) ClearViews() ViewGroup {
	return t.RemoveAll()
}

func (t *VStack) AppendViews(views ...View) ViewGroup {
	return t.AddViews(views...)
}

func (t *VStack) AddViews(views ...View) *VStack {
	for _, v := range views {
		v.node().Style().Set("display", "block")
		t.addView(v)
	}
	return t
}

func (t *VStack) RemoveAll() *VStack {
	t.absComponent.removeAll()
	return t
}

func (t *VStack) Style(style ...Style) *VStack {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *VStack) Self(ref **VStack) *VStack {
	*ref = t
	return t
}
