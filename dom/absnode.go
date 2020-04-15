package dom

import "syscall/js"

type absNode struct {
	val js.Value
}

func (n absNode) Style() Style {
	return Style{val: n.val.Get("style")}
}

// The TextContent property of the absNode interface represents the text content of the node and its descendants.
func (n absNode) TextContent() string {
	return n.val.Call("textContent").String()
}

// The SetTextContent property of the absNode interface represents the text content of the node and its descendants.
func (n absNode) SetTextContent(v string) {
	n.val.Set("textContent", v)
}

func (n absNode) Clear() {
	n.SetTextContent("")
}

// The Element property innerHTML gets or sets the HTML or XML markup contained within the element.
//
//Note: If a <div>, <span>, or <noembed> node has a child text node that includes the characters (&), (<), or (>), innerHTML returns these characters as the HTML entities "&amp;", "&lt;" and "&gt;" respectively. Use Node.textContent to get a raw copy of these text nodes' contents.
//To insert the HTML into the document rather than replace the contents of an element, use the method insertAdjacentHTML().
func (n absNode) InnerHTML() string {
	return n.val.Call("innerHTML").String()
}

// The Element property innerHTML gets or sets the HTML or XML markup contained within the element.
//
//Note: If a <div>, <span>, or <noembed> node has a child text node that includes the characters (&), (<), or (>), innerHTML returns these characters as the HTML entities "&amp;", "&lt;" and "&gt;" respectively. Use Node.textContent to get a raw copy of these text nodes' contents.
//To insert the HTML into the document rather than replace the contents of an element, use the method insertAdjacentHTML().
func (n absNode) SetInnerHTML(v string) {
	n.val.Set("innerHTML", v)
}

func (n absNode) SetInnerText(v string) absNode {
	n.val.Set("innerText", v)
	return n
}

// The absNode.appendChild() method adds a node to the end of the list of children of a specified parent node. If the given child is a reference to an existing node in the document, appendChild() moves it from its current position to the new position (there is no requirement to remove the node from its parent node before appending it to some other node).
//
//This means that a node can't be in two points of the document simultaneously. So if the node already has a parent, the node is first removed, then appended at the new position. The absNode.cloneNode() method can be used to make a copy of the node before appending it under the new parent. Note that the copies made with cloneNode will not be automatically kept in sync.
//
//If the given child is a DocumentFragment, the entire contents of the DocumentFragment are moved into the child list of the specified parent node.
func (n absNode) AppendChild(aChild Node) Node {
	return absNode{n.val.Call("appendChild", aChild.Unwrap())}
}

func (n absNode) InsertBefore(newNode, refNode Node) Node {
	return absNode{n.val.Call("insertBefore", newNode.Unwrap(), refNode.Unwrap())}
}

func (n absNode) RemoveChild(aChild Node) Node {
	return absNode{n.val.Call("removeChild", aChild.Unwrap())}
}

func (n absNode) Unwrap() js.Value {
	return n.val
}

func (n absNode) NextSibling() Node {
	return absNode{n.val.Get("nextSibling")}
}
