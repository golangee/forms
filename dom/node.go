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

import "syscall/js"

// Node is an interface from which various types of DOM API objects inherit. This allows these types to be treated similarly; for example, inheriting the same set of methods or being tested in the same way.
//
// All of the following interfaces inherit the absNode interface's methods and properties: Document, Element, Attr, CharacterData (which Text, Comment, and CDATASection inherit), ProcessingInstruction, DocumentFragment, DocumentType, Notation, Entity, EntityReference
//
// These interfaces may return null in certain cases where the methods and properties are not relevant. They may throw an exception — for example when adding children to a node type for which no children can exist.
type Node interface {
	TextContent() string
	SetTextContent(v string)
	AppendChild(aChild Node) Node
	Unwrap() js.Value
}
