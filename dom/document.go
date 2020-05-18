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

// The Document interface represents any web page loaded in the browser and serves as an entry point into the web page's content, which is the DOM tree. The DOM tree includes elements such as <body> and <table>, among many others. It provides functionality globally to the document, like how to obtain the page's URL and create new elements in the document.
// The Document interface describes the common properties and methods for any kind of document. Depending on the document's type (e.g. HTML, XML, SVG, …), a larger API is available: HTML documents, served with the "text/html" content type, also implement the HTMLDocument interface, whereas XML and SVG documents implement the XMLDocument interface.
type Document struct {
	val js.Value
	absNode
}

// The Document method getElementById() returns an Element object representing the element whose id property matches the specified string. Since element IDs are required to be unique if specified, they're a useful way to get access to a specific element quickly.
//
// If you need to get access to an element which doesn't have an ID, you can use querySelector() to find the element using any selector.
func (d Document) GetElementById(id string) Element {
	v := d.val.Call("getElementById", id)
	return Element{v, absNode{v}, absEventTarget{v}}
}

// In an HTML document, the document.createElement() method creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
func (d Document) CreateElement(name string) Element {
	v := d.val.Call("createElement", name)
	return Element{v, absNode{v}, absEventTarget{v}}
}

// In an HTML document, the document.createElement() method creates the HTML element specified by tagName, or an HTMLUnknownElement if tagName isn't recognized.
func (d Document) CreateElementNS(ns string, name string) Element {
	v := d.val.Call("createElementNS", ns, name)
	return Element{v, absNode{v}, absEventTarget{v}}
}

func (d Document) Body() Element {
	body := d.val.Get("body")
	return Element{body, absNode{body}, absEventTarget{body}}
}

func (d Document) DocumentElement() Element {
	body := d.val.Get("documentElement")
	return Element{body, absNode{body}, absEventTarget{body}}
}
