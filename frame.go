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

type Frame struct {
	v View
	*absComponent
}

func NewFrame() *Frame {
	t := &Frame{}
	t.absComponent = newComponent(t, "div")
	/* //this is center inside, but any flex-grow will not work inside it
	t.node().Style().Set("margin", "0")
	t.node().Style().Set("position", "absolute")
	t.node().Style().Set("top", "50%")
	t.node().Style().Set("left", "50%")
	t.node().Style().Set("transform", "translate(-50%,-50%)")
	*/
	/*
		t.node().Style().Set("max-width", "600")
		t.node().Style().Set("margin-left", "auto")
		t.node().Style().Set("margin-right", "auto")*/
	return t
}

func (t *Frame) SetView(v View) *Frame {
	if t.v != nil {
		t.absComponent.removeView(t.v)
		t.v = nil
	}
	t.absComponent.addView(v)
	t.v = v
	return t
}

func (t *Frame) Style(style ...Style) *Frame {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Frame) Self(ref **Frame) *Frame {
	*ref = t
	return t
}
