package dom

import "syscall/js"

// The Location interface represents the location (URL) of the object it is linked to. Changes done on it are reflected on the object it relates to. Both the Document and Window interface have such a linked Location, accessible via Document.location and Window.location respectively.
type Location struct {
	val js.Value
}

// Is a DOMString containing a '#' followed by the fragment identifier of the URL.
func (l Location) Hash() string {
	return l.val.Get("hash").String()
}
