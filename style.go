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
const DrawerTitle FontStyle = "mdc-drawer__title"
const DrawerSubTitle FontStyle = "mdc-drawer__subtitle"

var DefaultPadding = 8

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
		element.Style().SetPadding(intToPx(DefaultPadding))
	})
}

func Repel() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginBottom("1em")
	})
}

func Margin() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMargin(intToPx(DefaultPadding))
	})
}

func CenterHorizontal() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("margin-left", "auto")
		element.Style().Set("margin-right", "auto")
	})
}

func AlignBottom() Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("margin-bottom", "0px")
	})
}

func MarginTop(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().SetMarginTop(string(scalar))
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

type Scalar string

type scalarSlice []Scalar

func (s scalarSlice) toStrings() []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = string(v)
	}
	return res
}

func Auto() Scalar {
	return Scalar("auto")
}

func Percent(i int) Scalar {
	return Scalar(strconv.Itoa(i) + "%")
}

func PercentViewPortHeight(i int) Scalar {
	return Scalar(strconv.Itoa(i) + "vh")
}

func PercentViewPortWidth(i int) Scalar {
	return Scalar(strconv.Itoa(i) + "vw")
}

func Pixel(i int) Scalar {
	return Scalar(strconv.Itoa(i) + "px")
}

func Width(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("width", string(scalar))
	})
}

func MinWidth(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("min-width", string(scalar))
	})
}

func Height(scalar Scalar) Style {
	return styleFunc(func(element dom.Element) {
		element.Style().Set("height", string(scalar))
	})
}
