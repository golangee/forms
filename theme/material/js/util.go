package js

import (
	"github.com/worldiety/wtk/dom"
	"strings"
	"syscall/js"
)

type MDCName string

const Ripple MDCName = "ripple"

type Foundation struct {
	val js.Value
}

func (f Foundation) Release() {
	f.val.Call("destroy")
}

// Attach invokes new mdc.<name>.MDC<Name> to apply the material web foundation magic
func Attach(name MDCName, element dom.Element) Foundation {
	mdcObject := js.Global().Get("mdc")
	pkg := mdcObject.Get(string(name))
	upperCase := strings.ToUpper(string(name[0:1])) + string(name[1:])
	newMDCClassObj := pkg.Get("MDC" + upperCase)
	return Foundation{newMDCClassObj.New(element.Unwrap())}
}
