package gowtk

type Padding struct {
	r, t, l, b int
}

var DefaultPadding = Padding{}

type Alignment int

const Leading Alignment = 0

// A VBox layouts its children from top to bottom
type VBox struct {
	Component
}

func (b *VBox) SetAlignment(alignment Alignment) *VBox {
	return b
}

// NewVBox creates a new instance and returns the pointer to it
func NewVBox(views ...View) *VBox {
	return &VBox{}
}

func (b *VBox) SetPadding(padding Padding) *VBox {
	return b
}
