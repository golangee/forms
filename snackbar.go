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
	"log"
	"syscall/js"
	"time"
)

type Snackbar struct {
	*absComponent
	label                h.Element
	btnLabel             h.Element
	btn                  h.Element
	fnd                  js2.Foundation
	action               func(v View)
	callActionAfterClose bool
	timeout              time.Duration
}

func NewSnackbar(text string, actionLabel string) *Snackbar {
	t := &Snackbar{}
	t.absComponent = newComponent(t, "div")
	h.Wrap(t.node(), h.Class("mdc-snackbar"),
		h.Div(h.Class("mdc-snackbar__surface"),
			h.Div(h.Class("mdc-snackbar__label"), h.Role("status")).Self(&t.label),
			h.Div(h.Class("mdc-snackbar__actions"),
				h.Button(h.Class("mdc-button"), h.Class("mdc-snackbar__action"), h.Type("button"),
					h.Div(h.Class("mdc-button__ripple")),
					h.Span(h.Class("mdc-button__label")).Self(&t.btnLabel),
				).Self(&t.btn),
			),
		),
	)
	t.SetText(text)
	t.SetActionLabel(actionLabel)
	t.addResource(t.btn.AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
		t.callActionAfterClose = true
		return nil
	}, false))

	t.addResource(t.node().AddEventListener("MDCSnackbar:closed", func(this js.Value, args []js.Value) interface{} {
		if t.callActionAfterClose && t.action != nil {
			t.action(t)
		}
		t.destroy()
		return nil
	}, false))
	return t
}

func (t *Snackbar) SetAction(action func(v View)) *Snackbar {
	t.action = action
	return t
}

func (t *Snackbar) SetActionLabel(str string) *Snackbar {
	t.btnLabel.SetText(str)
	return t
}

func (t *Snackbar) SetText(str string) *Snackbar {
	t.label.SetTextContent(str)
	return t
}

// SetTimeout value must be between 4000 and 10000 (or -1 to disable the timeout completely) or an error will be thrown.
// Defaults is 5 seconds.
func (t *Snackbar) SetTimeout(d time.Duration) *Snackbar {
	t.timeout = d
	return t
}

func (t *Snackbar) Show(v View) *Snackbar {
	t.callActionAfterClose = false
	wnd := getWindow(v)
	if wnd == nil {
		log.Println("cannot show snackbar, view not attached")
		return t
	}
	t.fnd = js2.Attach(js2.Snackbar, t.node())
	if t.timeout != 0 {
		d := t.timeout.Milliseconds()
		if d < 4000 {
			d = -1
		}
		t.fnd.Unwrap().Set("timeoutMs", d)
	}
	wnd.AddView(t)
	t.fnd.Unwrap().Call("open")
	return t
}

func (t *Snackbar) destroy() {
	t.callActionAfterClose = false
	wnd := getWindow(t)
	t.Close()
	t.fnd.Release()
	if wnd == nil {
		return
	}
	wnd.RemoveView(t)
}

func (t *Snackbar) Release() {
	t.destroy()
	t.absComponent.Release()
}

func (t *Snackbar) Close() *Snackbar {
	t.fnd.Unwrap().Call("close")
	return t
}

func (t *Snackbar) Style(style ...Style) *Snackbar {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Snackbar) Self(ref **Snackbar) *Snackbar {
	*ref = t
	return t
}
