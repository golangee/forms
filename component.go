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
	this       View // this is a pointer to the actual inheriting view
	children   []View
	par        View // nil, if not attached
	elem       dom.Element
	resources  []Resource
	scope      context.Context
	cancelFunc context.CancelFunc
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

func (b *absComponent) addResource(r Resource) *absComponent {
	b.resources = append(b.resources, r)
	return b
}

func (b *absComponent) style(styles ...Style) *absComponent {
	for _, s := range styles {
		s.applyCSS(b.elem)
	}
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
		view.detach()
	}
	b.children = nil
	b.elem.SetTextContent("")
}

func (b *absComponent) removeViewAt(i int) View {
	a := b.children
	child := a[i]
	child.detach()
	b.elem.RemoveChild(child.node())
	if i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	a[len(a)-1] = nil
	a = a[:len(a)-1]
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

func (b *absComponent) addEventListener(t event.Type, f func(v View)) {
	ref := b.this.node().AddEventListener(string(t), func(_ js.Value, _ []js.Value) interface{} {
		f(b.this)
		return nil
	}, false)
	b.resources = append(b.resources, ref)
}

func (b *absComponent) Context() Context {
	if b.parent() == nil {
		log.Println("here is " + reflect.TypeOf(b.this).String() + " and i'm not attached")
		return nil
	}
	return b.parent().Context()
}
