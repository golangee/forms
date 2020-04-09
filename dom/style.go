package dom

import "syscall/js"

type Style struct {
	val js.Value
}

func (s Style) SetColor(v string) {
	s.val.Set("color", v)
}

func (s Style) SetPaddingLeft(v string) {
	s.val.Set("padding-left", v)
}

func (s Style) SetPaddingTop(v string) {
	s.val.Set("padding-top", v)
}

func (s Style) SetPaddingBottom(v string) {
	s.val.Set("padding-bottom", v)
}

func (s Style) SetPadding(v string) {
	s.val.Set("padding", v)
}

func (s Style) SetMargin(v string) {
	s.val.Set("margin", v)
}

func (s Style) SetMarginBottom(v string) {
	s.val.Set("margin-bottom", v)
}

func (s Style) SetMarginTop(v string) {
	s.val.Set("margin-top", v)
}

func (s Style) Unwrap() js.Value {
	return s.val
}

func (s Style) Set(k, v string) {
	s.val.Set(k, v)
}

func (s Style) SetProperty(k, v string) {
	s.val.Call("setProperty", k, v)
}
