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

package js

import (
	"github.com/golangee/forms/dom"
	"strings"
	"syscall/js"
)

type MDCName string

const Ripple MDCName = "ripple"
const TextField MDCName = "textField"
const Dialog MDCName = "dialog"
const Menu MDCName = "menu"
const MenuSurface MDCName = "menuSurface"
const TopAppBar MDCName = "topAppBar"
const Drawer MDCName = "drawer"
const List MDCName = "list"
const DataTable MDCName = "dataTable"
const Checkbox MDCName = "checkbox"
const FormField MDCName = "formField"
const Select MDCName = "select"
const TabBar MDCName = "tabBar"
const TabScroller MDCName = "tabScroller"
const LinearProgress MDCName = "linearProgress"
const Snackbar MDCName = "snackbar"

type Foundation struct {
	val js.Value
}

func (f Foundation) Release() {
	if f.IsValid() {
		//log.Println("destroyed ", f.val)
		f.val.Call("destroy")
	}
}
func (f Foundation) Unwrap() js.Value {
	return f.val
}

func (f Foundation) IsValid() bool {
	return !f.val.IsUndefined() && !f.val.IsNull()
}

// Attach invokes new mdc.<name>.MDC<Name> to apply the material web foundation magic
func Attach(name MDCName, element dom.Element) Foundation {
	mdcObject := js.Global().Get("mdc")
	pkg := mdcObject.Get(string(name))
	upperCase := strings.ToUpper(string(name[0:1])) + string(name[1:])
	newMDCClassObj := pkg.Get("MDC" + upperCase)
	return Foundation{newMDCClassObj.New(element.Unwrap())}
}
