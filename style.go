package wtk

import (
	"github.com/worldiety/wtk/dom"
)

// A Style modifies different kinds of visualization of a View.
type Style interface {
	internalStyle
}

type internalStyle interface {
	apply(element dom.Element)
}

type styleFunc func(element dom.Element)

func (s styleFunc) apply(element dom.Element) {
	s(element)
}

func PadTop(v float64) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingLeft(floatToPx(v))
	})
}

func ForegroundColor(color Color)Style{
	return styleFunc(func(element dom.Element) {
		element.Style().SetColor(color.String())
	})
}