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

func (s Style) Unwrap() js.Value {
	return s.val
}
