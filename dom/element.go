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

func (s Element) AddClass(v string) {
	s.val.Get("classList").Call("add", v)
}

