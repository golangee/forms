package dom

import (
	"syscall/js"
)

type absEventTarget struct {
	val js.Value
}

func (t absEventTarget) AddEventListener(typ string, listener func(this js.Value, args []js.Value) interface{}, once bool) js.Func {
	f := js.FuncOf(listener)
	t.val.Call("addEventListener", typ, f, once)
	return f
}
