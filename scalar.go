package forms

import (
	"github.com/golangee/forms/dom"
	"strconv"
)

type Scalar string

type scalarSlice []Scalar

func (s scalarSlice) toStrings() []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = string(v)
	}
	return res
}

// Auto may have different meanings, depending on the used context.
//  Grid
//    - if possible, take mostly the required space (fit-content), however
//    - if not reasonable (like others have fixed sizes or everything is auto), stretch to availabel space
func Auto() Scalar {
	return Scalar("auto")
}

// Fraction is used for Grid containers to define how a fraction of the remaining available space is distributed.
func Fraction(f int) Scalar {
	return Scalar(strconv.Itoa(f) + "fr")
}

// Cover is only valid for the background-size attribute
func Cover() Scalar {
	return Scalar("cover")
}

// Contain is only valid for the background-size attribute
func Contain() Scalar {
	return Scalar("contain")
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

// Pixel is in a kind of DIP, density independent pixel. So your value is scaled with the display density to
// something reasonable.
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
