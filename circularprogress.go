package wtk

import (
	h "github.com/worldiety/wtk/dom"
)

type CircularProgress struct {
	*absComponent
}

func NewCircularProgress() *CircularProgress {
	t := &CircularProgress{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("text-align","center")
	h.Wrap(t.node(),
		h.Svg(h.Class("wtk-mdc-circular-progress"), h.ViewBox("25 25 50 50"),
			h.Circle(h.Class("wtk-mdc-circular-progress__path"), h.Cx("50"), h.Cy("50"), h.R("20"), h.Fill("none"), h.StrokeWidth("4"), h.StrokeMiterlimit("10")),
		),

	)
	return t
}

func (t *CircularProgress) Style(style ...Style) *CircularProgress {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *CircularProgress) Self(ref **CircularProgress) *CircularProgress {
	*ref = t
	return t
}
