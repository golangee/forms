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

func (s Element) SetClassName(str string) Element {
	s.val.Set("className", str)
	return s
}

func (s Element) AddClass(v string) Element {
	s.val.Get("classList").Call("add", v)
	return s
}

func (s Element) SetDisabled(b bool) Element {
	s.val.Set("disabled", b)
	return s
}
