package wtk

import (
	"github.com/worldiety/wtk/dom"
	"syscall/js"
	"time"
)

func Post(d time.Duration, f func()) {
	var wrapper js.Func
	wrapper = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		f()
		wrapper.Release()
		return nil
	})
	dom.GetWindow().Unwrap().Call("setTimeout", wrapper, d.Milliseconds())
}
