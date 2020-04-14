package dom

import "syscall/js"

// Element is the most general base class from which all objects in a Document inherit. It only has methods and properties common to all kinds of elements. More specific classes inherit from Element. For example, the HTMLElement interface is the base interface for HTML elements, while the SVGElement interface is the basis for all SVG elements. Most functionality is specified further down the class hierarchy.&
//
// Languages outside the realm of the Web platform, like XUL through the XULElement interface, also implement Element.
type Element struct {
	val js.Value
	absNode
	absEventTarget
}

func NewElement(v js.Value) Element {
	return Element{v, absNode{v}, absEventTarget{v}}
}

func (s Element) SetClassName(str string) Element {
	s.val.Set("className", str)
	return s
}

func (s Element) AddClass(v string) Element {
	s.val.Get("classList").Call("add", v)
	return s
}

func (s Element) HasClass(v string) bool {
	return s.val.Get("classList").Call("contains", v).Bool()
}

func (s Element) RemoveClass(v string) Element {
	s.val.Get("classList").Call("remove", v)
	return s
}

func (s Element) SetDisabled(b bool) Element {
	s.val.Set("disabled", b)
	return s
}

func (s Element) SetId(id string) Element {
	s.val.Set("id", id)
	return s
}

func (s Element) SetRole(r string) Element {
	s.val.Set("role", r)
	return s
}

func (s Element) SetScope(r string) Element {
	s.val.Set("scope", r)
	return s
}

func (s Element) SetAriaHidden(b bool) Element {
	s.val.Set("aria-hidden", b)
	return s
}

func (s Element) SetTabIndex(i int) Element {
	s.val.Set("tabindex", i)
	return s
}

func (s Element) SetAriaOrientation(o string) Element {
	s.val.Set("aria-orientation", o)
	return s
}

func (s Element) SetAriaLabel(l string) Element {
	s.val.Set("aria-label", l)
	return s
}

func (s Element) Id() string {
	return s.val.Get("id").String()
}
func (s Element) SetFor(id string) Element {
	s.val.Set("htmlFor", id)
	return s
}

func (s Element) SetType(t string) Element {
	s.val.Set("type", t)
	return s
}

func (s Element) SetText(v string) Element {
	s.absNode.SetTextContent(v)
	return s
}

func (s Element) SetHref(v string) Element {
	s.absNode.Unwrap().Set("href", v)
	return s
}

func (s Element) SetAttrNS(ns string, key, val string) Element {
	s.absNode.Unwrap().Call("setAttributeNS", ns, key, val)
	return s
}

func (s Element) SetAttr(key, val string) Element {
	s.absNode.Unwrap().Call("setAttribute", key, val)
	return s
}

func (s Element) FirstChild() Element {
	return NewElement(s.Unwrap().Get("firstChild"))
}
