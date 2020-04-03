package wtk

import (
	"github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/event"
	"syscall/js"
)

// absComponent contains the basic implementation of a View
type absComponent struct {
	this      View // this is a pointer to the actual inheriting view
	children  []View
	par       View // nil, if not attached
	elem      dom.Element
	resources []Resource
}

func newComponent(self View, tag string) *absComponent {
	return &absComponent{
		this: self,
		elem: dom.CreateElement(tag),
	}
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
	for _, child := range b.children {
		child.Release()
	}
	// we do not remove the child nodes, because the root node can just be GCed, which is probably more efficient
	for _, res := range b.resources {
		res.Release()
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
		return nil
	}
	return b.parent().Context()
}
