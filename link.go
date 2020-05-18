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

type LinkTarget string

const TargetBlank LinkTarget = "_blank"
const TargetSelf LinkTarget = "_self"
const TargetParent LinkTarget = "_parent"
const TargetTop LinkTarget = "_top"

type Link struct {
	*absComponent
}

func NewLink(caption string, href string) *Link {
	t := &Link{}
	t.absComponent = newComponent(t, "a")
	t.SetCaption(caption)
	t.SetRef(href)
	return t
}

func (t *Link) SetRef(r string) *Link {
	t.elem.SetHref(r)
	return t
}

func (t *Link) SetCaption(str string) *Link {
	t.absComponent.elem.SetText(str)
	return t
}

func (t *Link) Style(style ...Style) *Link {
	t.absComponent.style(style...)
	return t
}

func (t *Link) SetTarget(target LinkTarget) *Link {
	t.node().Unwrap().Set("target", string(target))
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Link) Self(ref **Link) *Link {
	*ref = t
	return t
}
