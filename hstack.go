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

// HStack is a horizontal grid with just a single row.
type HStack struct {
	*Grid
}

// NewHStack creates a new horizontal view stack.
func NewHStack(views ...View) *HStack {
	t := &HStack{}
	t.Grid = NewGrid()
	t.AddViews(views...)
	return t
}

func (t *HStack) ClearViews() ViewGroup {
	return t.RemoveAll()
}

func (t *HStack) AppendViews(views ...View) ViewGroup {
	return t.AddViews(views...)
}

func (t *HStack) AddViews(views ...View) *HStack {
	for _, v := range views {
		v.node().Style().Set("grid-row", "1") // unclear, if this is the best way
		t.addView(v)
	}
	return t
}

func (t *HStack) RemoveAll() *HStack {
	t.Grid.removeAll()
	return t
}

// Style applies generic style attributes.
func (t *HStack) Style(style ...Style) *HStack {
	t.Grid.style(style...)
	return t
}

// StyleFor applies the given styles only if the criteria is met.
func (t *HStack) StyleFor(criteria MediaCriteria, style ...Style) *HStack {
	t.absComponent.styleFor(criteria, style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *HStack) Self(ref **HStack) *HStack {
	*ref = t
	return t
}
