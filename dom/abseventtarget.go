package dom

import (
	"log"
	"syscall/js"
)

type absEventTarget struct {
	val js.Value
}

func (t absEventTarget) AddEventListener(typ string, listener func(value js.Value), once bool) js.Func {
	log.Println("addEventListener added")

	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		log.Println("addEventListener called: ", this, this.Get("id"))
		listener(args[0])
		return nil
	})
	t.val.Call("addEventListener", typ, f, once)
	return f
}
