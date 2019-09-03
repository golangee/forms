package gowtk

var _ View = (*HBox)(nil)

// A HBox layouts its children from left to right
type HBox struct {
	Component
}

// NewHBox creates a new instance and returns the pointer to it
func NewHBox(views ...View) *HBox {
	return &HBox{}
}
