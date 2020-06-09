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

// Card represents the material card component, see https://material.io/components/cards.
type Card struct {
	*absComponent
}

// NewCard creates a new card component.
func NewCard(views ...View) *Card {
	t := &Card{}
	t.absComponent = newComponent(t, "div")
	t.node().AddClass("mdc-card")
	t.node().Style().Set("box-sizing", "border-box")
	t.node().Style().Set("overflow","auto")
	t.AddViews(views...)
	return t
}

func (t *Card) ClearViews() ViewGroup {
	return t.RemoveAll()
}

func (t *Card) AppendViews(views ...View) ViewGroup {
	return t.AddViews(views...)
}

func (t *Card) AddViews(views ...View) *Card {
	for _, v := range views {
		t.addView(v)
	}
	return t
}

func (t *Card) RemoveAll() *Card {
	t.absComponent.removeAll()
	return t
}

func (t *Card) Style(style ...Style) *Card {
	t.absComponent.style(style...)
	return t
}

// StyleFor applies the given styles only if the criteria is met.
func (t *Card) StyleFor(criteria MediaCriteria, style ...Style) *Card {
	t.absComponent.styleFor(criteria, style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Card) Self(ref **Card) *Card {
	*ref = t
	return t
}
