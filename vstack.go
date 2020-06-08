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

var _ View = (*VStack)(nil)

// VStack is actually a vertical grid, with just a single column.
type VStack struct {
	*Grid
}

// NewVStack creates a new vertical stack
func NewVStack(views ...View) *VStack {
	t := &VStack{}
	t.Grid = NewGrid()
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
		//v.node().Style().Set("display", "block") // TODO this is wrong and destroy nested things like grid
		t.Grid.AddView(v, GridLayoutParams{})
	}
	return t
}

func (t *VStack) RemoveAll() *VStack {
	t.Grid.RemoveAll()
	return t
}

func (t *VStack) Style(style ...Style) *VStack {
	t.Grid.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *VStack) Self(ref **VStack) *VStack {
	*ref = t
	return t
}
