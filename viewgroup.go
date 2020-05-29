package forms

// ViewGroup is a generic interface which all view groups should support. Sadly go does not
// allow covariant return types in interfaces, so we have to define another signature set.
// See also https://github.com/golang/go/issues/30602.
type ViewGroup interface {
	View
	ClearViews() ViewGroup
	AppendViews(views ...View) ViewGroup
}
