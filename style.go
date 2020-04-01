package wtk

import (
	"github.com/worldiety/wtk/dom"
)

type FontStyle string

const Title FontStyle = "titleText"

// A Style modifies different kinds of visualization of a View.
type Style interface {
	internalStyle
}

type internalStyle interface {
	applyCSS(element dom.Element)
}

type styleFunc func(element dom.Element)

func (s styleFunc) applyCSS(element dom.Element) {
	s(element)
}

func PadLeft(v float64) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingLeft(floatToPx(v))
	})
}

func ForegroundColor(color Color) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetColor(color.String())
	})
}

func Font(name FontStyle) Style {
	return styleFunc(func(element dom.Element) {
		element.AddClass(string(name))
	})
}
