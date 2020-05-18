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

//+build wasm,js

package dom

import "syscall/js"

//The Window interface represents a window containing a DOM document; the document property points to the DOM document loaded in that window. A window for a given document can be obtained using the document.defaultView property.
//
//A global variable, window, representing the window in which the script is running, is exposed to JavaScript code.
//
//The Window interface is home to a variety of functions, namespaces, objects, and constructors which are not necessarily directly associated with the concept of a user interface window. However, the Window interface is a suitable place to include these items that need to be globally available. Many of these are documented in the JavaScript Reference and the DOM Reference.
//
//In a tabbed browser, each tab is represented by its own Window object; the global window seen by JavaScript code running within a given tab always represents the tab in which the code is running. That said, even in a tabbed browser, some properties and methods still apply to the overall window that contains the tab, such as resizeTo() and innerHeight. Generally, anything that can't reasonably pertain to a tab pertains to the window instead.
type Window struct {
	val js.Value
	absEventTarget
}

// Window returns the window global
func GetWindow() Window {
	v := js.Global().Get("window")
	return Window{v, absEventTarget{v}}
}

// Document returns a reference to the document contained in the window.
func (w Window) Document() Document {
	v := w.val.Get("document")
	return Document{v, absNode{v}}
}

func (w Window) Location() Location {
	return Location{w.val.Get("location")}
}
func (w Window) Unwrap() js.Value {
	return w.val
}

func (w Window) MatchesMedia(criteria string) bool {
	return w.val.Call("matchMedia", criteria).Get("matches").Bool()
}

func CreateElement(name string) Element {
	return GetWindow().Document().CreateElement(name)
}

func CreateElementNS(ns string, name string) Element {
	return GetWindow().Document().CreateElementNS(ns, name)
}
