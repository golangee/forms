package gowtk

// A VBox layouts its children from top to bottom
type VBox struct {
	Component
}

// NewVBox creates a new instance and returns the pointer to it
func NewVBox(views ...View) *VBox {
	return &VBox{}
}
