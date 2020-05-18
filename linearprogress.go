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
	js2 "github.com/golangee/forms/theme/material/js"
)

type LinearProgress struct {
	*absComponent
	fnd js2.Foundation
}

func NewLinearProgress() *LinearProgress {
	t := &LinearProgress{}
	t.absComponent = newComponent(t, "div")
	h.Wrap(t.node(), h.Role("progressbar"), h.Class("mdc-linear-progress"),
		h.Div(h.Class("mdc-linear-progress__buffering-dots")),
		h.Div(h.Class("mdc-linear-progress__buffer")),
		h.Div(h.Class("mdc-linear-progress__bar", "mdc-linear-progress__primary-bar"),
			h.Span(h.Class("mdc-linear-progress__bar-inner"),
			),
		),
		h.Div(h.Class("mdc-linear-progress__bar", "mdc-linear-progress__secondary-bar"),
			h.Span(h.Class("mdc-linear-progress__bar-inner")),
		),
	)
	t.fnd = js2.Attach(js2.LinearProgress, t.node())
	t.addResource(t.fnd)
	return t
}

func (t *LinearProgress) SetIndeterminate(b bool) *LinearProgress {
	t.fnd.Unwrap().Set("determinate", !b)
	return t
}

func (t *LinearProgress) SetProgress(f float64) *LinearProgress {
	t.fnd.Unwrap().Set("progress", f)
	return t
}

func (t *LinearProgress) SetSecondaryProgress(f float64) *LinearProgress {
	t.fnd.Unwrap().Set("buffer", f)
	return t
}

func (t *LinearProgress) Style(style ...Style) *LinearProgress {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *LinearProgress) Self(ref **LinearProgress) *LinearProgress {
	*ref = t
	return t
}
