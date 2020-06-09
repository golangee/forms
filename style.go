// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package forms

import (
	"github.com/golangee/forms/dom"
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
const DrawerTitle FontStyle = "mdc-drawer__title"
const DrawerSubTitle FontStyle = "mdc-drawer__subtitle"

type TypeWeight string

const WeightNormal TypeWeight = "normal"
const WeightLighter TypeWeight = "lighter"
const WeightBolder TypeWeight = "500"

var DefaultPadding = Pixel(8)

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

func PadLeft(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingLeft(string(scalar))
	})
}

func PadRight(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingRight(string(scalar))
	})
}

func PadTop(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingTop(string(scalar))
	})
}

func PadBottom(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPaddingBottom(string(scalar))
	})
}

func Padding() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetPadding(string(DefaultPadding))
	})
}

func Repel() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginBottom("1em")
	})
}

func Margin() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMargin(string(DefaultPadding))
	})
}

func MarginTop(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginTop(string(scalar))
	})
}

func MarginBottom(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginBottom(string(scalar))
	})
}

func MarginLeft(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginLeft(string(scalar))
	})
}

func MarginRight(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginRight(string(scalar))
	})
}

func MaxWidth(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("max-width", string(scalar))
	})
}

func ForegroundColor(color Color) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetColor(color.String())
	})
}

func BackgroundColor(color Color) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetBackgroundColor(color.String())
	})
}

func Font(name FontStyle) Style {
	return styleFunc(func(element dom.Element) {
		element.AddClass(string(name))
	})
}

func FontWeight(weight TypeWeight) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("font-weight", string(weight))
	})
}

// BorderRadius sets all 4 edges to the same radius
func BorderRadius(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("border-radius", string(scalar))
	})
}
