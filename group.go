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

type Group struct {
	*absComponent
}

func NewGroup(views ...View) *Group {
	t := &Group{}
	t.absComponent = newComponent(t, "div")
	t.AddViews(views...)
	return t
}

func (t *Group) ClearViews() ViewGroup {
	return t.RemoveAll()
}

func (t *Group) AppendViews(views ...View) ViewGroup {
	return t.AddViews(views...)
}

func (t *Group) AddViews(views ...View) *Group {
	for _, v := range views {
		t.addView(v)
	}
	return t
}

func (t *Group) RemoveAll() *Group {
	t.absComponent.removeAll()
	return t
}

func (t *Group) Style(style ...Style) *Group {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Group) Self(ref **Group) *Group {
	*ref = t
	return t
}
