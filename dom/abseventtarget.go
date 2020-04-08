package dom

import (
	"syscall/js"
)

type Func struct {
	js.Func
	val  js.Value
	typ  string
	once bool
}

func (f Func) Release() {
	f.val.Call("removeEventListener", f.typ, f.Func, f.once)
	f.Func.Release()
}

type absEventTarget struct {
	val js.Value
}

func (t absEventTarget) AddEventListener(typ string, listener func(this js.Value, args []js.Value) interface{}, once bool) Func {
	f := js.FuncOf(listener)
	t.val.Call("addEventListener", typ, f, once)
	return Func{
		Func: f,
		val:  t.val,
		typ:  typ,
		once: once,
	}
}
