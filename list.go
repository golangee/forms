package wtk

import (
	"github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/theme/material/icon"
	js2 "github.com/worldiety/wtk/theme/material/js"
	"syscall/js"
)

type List struct {
	*absComponent
	fnd           js2.Foundation
	selectionList bool
}

func NewList() *List {
	t := &List{}
	t.absComponent = newComponent(t, "nav")
	t.node().AddClass("mdc-list")
	t.fnd = js2.Attach(js2.List, t.node())
	t.addResource(t.fnd)
	return t
}

func NewSelectionList() *List {
	t := &List{}
	t.absComponent = newComponent(t, "nav")
	t.node().AddClass("mdc-list").SetRole("listbox")
	t.fnd = js2.Attach(js2.List, t.node())
	t.fnd.Unwrap().Set("singleSelection", true)
	t.selectionList = true
	t.addResource(t.fnd)
	return t
}

func (t *List) AddSelectListener(action func(idx int)) *List {
	t.addResource(t.node().AddEventListener("MDCList:action", func(this js.Value, args []js.Value) interface{} {
		// args[0] contains isTrusted and not the index, as documented
		selectedIndex := t.SelectedIndex()
		action(selectedIndex)
		return nil
	}, false))
	return t
}

func (t *List) SelectedIndex() int {
	if !t.selectionList {
		panic("not a selection list")
	}
	return t.fnd.Unwrap().Get("selectedIndex").Int()
}

func (t *List) SetSelectedIndex(idx int) *List {
	if !t.selectionList {
		panic("not a selection list")
	}
	t.fnd.Unwrap().Set("selectedIndex", idx)
	// another bug here: this does not fire MDCList:action
	t.node().Unwrap().Call("dispatchEvent", js.Global().Get("Event").New("MDCList:action"))
	return t
}

func (t *List) AddItems(items ...LstItem) *List {
	anySelected := false
	for _, item := range items {
		if t.selectionList {
			item.node().SetRole("option")
		}
		if item.isSelected() {
			anySelected = true
		}
		t.addView(item)
		t.addResource(js2.Attach(js2.Ripple, item.node()))
	}
	// a quickfix to reset tabindex to the selected
	if anySelected {
		for _, item := range items {
			item.node().Unwrap().Set("tabIndex", -1)
			if item.isSelected() {
				item.node().Unwrap().Set("tabIndex", 0)
			}
		}
	}

	return t
}

func (t *List) Style(style ...Style) *List {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *List) Self(ref **List) *List {
	*ref = t
	return t
}

type LstItem interface {
	myListItem()
	isSelected() bool
	View
}

type ListSeparator struct {
	*absComponent
}

func NewListSeparator() *ListSeparator {
	t := &ListSeparator{}
	t.absComponent = newComponent(t, "hr")
	t.node().AddClass("mdc-list-divider")
	return t
}

func (t *ListSeparator) myListItem() {
}

func (t *ListSeparator) isSelected() bool {
	return false
}

type ListHeader struct {
	*absComponent
}

func NewListHeader(caption string) *ListHeader {
	t := &ListHeader{}
	t.absComponent = newComponent(t, "h6")
	t.node().AddClass("mdc-list-group__subheader")
	t.node().SetText(caption)
	return t
}

func (t *ListHeader) SetText(s string) *ListHeader {
	t.node().SetText(s)
	return t
}

func (t *ListHeader) myListItem() {
}

func (t *ListHeader) isSelected() bool {
	return false
}

type ListItem struct {
	*absComponent
	span     dom.Element
	ico      dom.Element
	trailing dom.Element
	selected bool
}

func (t *ListItem) myListItem() {
}

func NewListItem(text string) *ListItem {
	t := &ListItem{}
	t.absComponent = newComponent(t, "a")
	t.node().AddClass("mdc-list-item")
	t.node().Unwrap().Set("tabIndex", 0)
	t.span = dom.CreateElement("span").AddClass("mdc-list-item__text").SetText(text)
	t.ico = dom.CreateElement("i").SetClassName("material-icons mdc-list-item__graphic")
	t.ico.Style().Set("display", "none")
	t.trailing = dom.CreateElement("div").AddClass("mdc-list-item__meta")
	t.trailing.Style().Set("display", "none")
	t.node().AppendChild(t.ico)
	t.node().AppendChild(t.span)
	t.node().AppendChild(t.trailing)

	return t
}

func (t *ListItem) SetText(s string) *ListItem {
	t.span.SetText(s)
	return t
}

func (t *ListItem) SetIcon(ico icon.Icon) *ListItem {
	if len(ico) > 0 {
		t.ico.Style().Set("display", "inherit")
	} else {
		t.ico.Style().Set("display", "none")
	}
	t.ico.SetText(string(ico))
	return t
}

func (t *ListItem) SetSelected(b bool) *ListItem {
	t.selected = b
	if b {
		t.node().AddClass("mdc-list-item--selected")
	} else {
		t.node().RemoveClass("mdc-list-item--selected")
	}
	return t
}

func (t *ListItem) AddClickListener(action func(v View)) *ListItem {
	t.addResource(t.node().AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
		action(t)
		return nil
	}, false))
	return t
}

func (t *ListItem) SetTrailingView(v View) *ListItem {
	if v != nil {
		t.trailing.Style().Set("display", "inherit")
	} else {
		t.trailing.Style().Set("display", "none")
		return t
	}
	v.attach(t)
	t.addResource(v) // todo: this should be done by attach?
	t.trailing.SetText("")
	t.trailing.AppendChild(v.node())
	return t
}

func (t *ListItem) isSelected() bool {
	return t.selected
}
