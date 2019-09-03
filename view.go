package gowtk

// A View is the common interface to represent a part in the view graph.
type View interface {
	// Children returns a (likely the backing) slice with all appended views so far. Do not modify.
	Children() []View

	// Attach prepares this view to be added to the actual view graph
	Attach()

	// Detach removes this view from the actual view graph
	Detach()
}
