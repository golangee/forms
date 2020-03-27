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
