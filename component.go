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
	"context"
	"github.com/golangee/forms/dom"
	"github.com/golangee/forms/event"
	"log"
	"reflect"
	"syscall/js"
)

// absComponent contains the basic implementation of a View
type absComponent struct {
	this          View // this is a pointer to the actual inheriting view
	children      []View
	par           View // nil, if not attached
	elem          dom.Element
	resources     []Resource
	scope         context.Context
	cancelFunc    context.CancelFunc
	defaultStyles []Style
	mediaMatcher  *MediaMatcher // usually nil, only non-nil if a conditional style has been set.
}

func newComponent(self View, tag string) *absComponent {
	b := &absComponent{
		this: self,
		elem: dom.CreateElement(tag),
	}
	b.scope, b.cancelFunc = context.WithCancel(context.Background())
	return b
}

func (b *absComponent) Scope() context.Context {
	return b.scope
}

func (b *absComponent) addResource(r Resource) {
	b.resources = append(b.resources, r)
}

func (b *absComponent) style(styles ...Style) *absComponent {
	for _, s := range styles {
		s.applyCSS(b.elem)
	}
	b.defaultStyles = styles
	return b
}

func (b *absComponent) styleFor(criteria MediaCriteria, styles ...Style) *absComponent {
	if b.mediaMatcher == nil {
		b.mediaMatcher = NewMediaMatcher(b.this, func(view View) {
			// apply all default styles, if nothing matches
			for _, s := range b.defaultStyles {
				s.applyCSS(b.elem)
			}
		})
	}

	b.mediaMatcher.Add(criteria, func(view View) {
		// apply the matched special styles only the first time a criteria is met
		for _, s := range styles {
			s.applyCSS(b.elem)
		}
	})

	b.mediaMatcher.Check()

	return b
}

func (b *absComponent) isAttached() bool {
	return b.par != nil
}

func (b *absComponent) attach(parent View) {
	assertNotAttached(b)
	b.par = parent
}

func (b *absComponent) detach() {
	assertAttached(b)
	b.par.node().RemoveChild(b.elem)
	b.par = nil
}

func (b *absComponent) node() dom.Element {
	return b.elem
}

func (b *absComponent) addView(child View) {
	if child == nil {
		panic("adding nil view not allowed")
	}
	b.elem.AppendChild(child.node())
	child.attach(b)
	b.children = append(b.children, child)
}

func (b *absComponent) removeView(child View) {
	for i, c := range b.children {
		if child == c {
			b.removeViewAt(i)
			return
		}
	}
}

func (b *absComponent) removeAll() {
	for _, view := range b.children {
		if view == nil {
			panic("illegal state: child view is nil")
		}
		view.detach()
	}
	b.children = nil
	b.elem.SetTextContent("")
}

// TODO unclear how to handle the Release: we do not own the view so destroying is probably wrong however without relasing, we will leak a lot of stuff
// Option a: add/set View takes ownership and remove releases it. Introduce a public detach method, if you need something else
// Option b: register at window, if detached and if the window gets a new view to show, release all detached view (ui state change)
func (b *absComponent) removeViewAt(i int) View {
	a := b.children
	child := a[i]
	child.detach()
	//b.elem.RemoveChild(child.node()) //TODO detach already removes the child from the dom?
	if i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	a[len(a)-1] = nil
	a = a[:len(a)-1]

	b.children = a
	return child
}

func (b *absComponent) parent() View {
	return b.par
}

func (b *absComponent) Release() {
	//log.Println("release of ", reflect.TypeOf(b.this).String())
	for _, child := range b.children {
		child.Release()
	}
	// we do not remove the child nodes, because the root node can just be GCed, which is probably more efficient
	for _, res := range b.resources {
		res.Release()
	}

	if b.cancelFunc != nil {
		b.cancelFunc()
	}
}

func (b *absComponent) addEventListener(t event.Type, f func(v View, params []js.Value)) {
	ref := b.this.node().AddEventListener(string(t), func(_ js.Value, params []js.Value) interface{} {
		params[0].Call("stopPropagation")
		f(b.this, params)
		return nil
	}, false)
	b.resources = append(b.resources, ref)
}

func (b *absComponent) Context() Context {
	if b.parent() == nil {
		log.Println("here is " + reflect.TypeOf(b.this).String() + " and i'm not attached, returning mock context")

		mock := &myContext{r: NewRouter()}
		return mock
	}
	return b.parent().Context()
}
