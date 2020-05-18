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

import (
	h "github.com/golangee/forms/dom"
)

type CircularProgress struct {
	*absComponent
}

func NewCircularProgress() *CircularProgress {
	t := &CircularProgress{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("text-align","center")
	h.Wrap(t.node(),
		h.Svg(h.Class("wtk-mdc-circular-progress"), h.ViewBox("25 25 50 50"),
			h.Circle(h.Class("wtk-mdc-circular-progress__path"), h.Cx("50"), h.Cy("50"), h.R("20"), h.Fill("none"), h.StrokeWidth("4"), h.StrokeMiterlimit("10")),
		),

	)
	return t
}

func (t *CircularProgress) Style(style ...Style) *CircularProgress {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *CircularProgress) Self(ref **CircularProgress) *CircularProgress {
	*ref = t
	return t
}
