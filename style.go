package wtk

import (
	"github.com/worldiety/wtk/dom"
	"strconv"
)

type FontStyle string

const Headline1 FontStyle = "mdc-typography--headline1"
const Headline2 FontStyle = "mdc-typography--headline2"
const Headline3 FontStyle = "mdc-typography--headline3"
const Headline4 FontStyle = "mdc-typography--headline4"
const Headline5 FontStyle = "mdc-typography--headline5"
const Headline6 FontStyle = "mdc-typography--headline6"
const Body FontStyle = "mdc-typography--body1"
const Body2 FontStyle = "mdc-typography--body2"
const Subtitle1 FontStyle = "mdc-typography--subtitle1"
const Subtitle2 FontStyle = "mdc-typography--subtitle1"
const Caption FontStyle = "mdc-typography--caption"
const Btn FontStyle = "mdc-typography--button"
const Overline FontStyle = "mdc-typography--overline"

var DefaultPadding = 8.0

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

func PadTop(v float64) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingTop(floatToPx(v))
	})
}

func PadBottom(v float64) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingBottom(floatToPx(v))
	})
}

func Padding() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPadding(floatToPx(DefaultPadding))
	})
}

func Repel() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginBottom("1em")
	})
}

func Margin() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMargin(floatToPx(DefaultPadding))
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

type Scalar string

func Percent(i int) Scalar {
	return Scalar(strconv.Itoa(i) + "%")
}

func Width(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("width", string(scalar))
	})
}
