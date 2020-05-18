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
