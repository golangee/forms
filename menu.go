package wtk

import (
	"github.com/worldiety/wtk/dom"
	js2 "github.com/worldiety/wtk/theme/material/js"
	"log"
	"syscall/js"
)

type Menu struct {
	*absComponent
	menu           dom.Element
	items          []*MenuItem
	fnd            js2.Foundation
	mFnd           js2.Foundation
	releaseOnClose bool
}

func NewMenu() *Menu {
	t := &Menu{}
	t.absComponent = newComponent(t, "div")
	t.node().SetClassName("mdc-menu mdc-menu-surface")

	t.menu = dom.CreateElement("ul").AddClass("mdc-list").SetRole("menu").SetAriaHidden(true).SetAriaOrientation("vertical").SetTabIndex(-1)
	t.node().AppendChild(t.menu)

	t.addResource(t.node().AddEventListener("MDCMenu:selected", func(this js.Value, args []js.Value) interface{} {
		idx := args[0].Get("detail").Get("index").Int()
		if idx < len(t.items) {
			cb := t.items[idx].f
			if cb != nil {
				cb(t.items[idx])
			}
		}

		return nil
	}, false))

	t.addResource(t.node().AddEventListener("MDCMenuSurface:closed", func(this js.Value, args []js.Value) interface{} {
		t.Release()
		//log.Println("on closed")
		return nil
	}, false))

	t.mFnd = js2.Attach(js2.Menu, t.node())
	t.absComponent.addResource(t.mFnd)

	t.fnd = js2.Attach(js2.MenuSurface, t.node())
	t.absComponent.addResource(t.fnd)

	dom.GetWindow().Document().Body().AppendChild(t.node())
	return t
}

// ShowMenu is a shortcut for NewMenu().ReleaseOnClose()...Show()
func ShowMenu(anchor View, items ...*MenuItem) *Menu {
	menu := NewMenu()
	for _, item := range items {
		menu.AddItem(item)
	}
	menu.Show(anchor)
	return menu
}

func (t *Menu) ReleaseOnClose() *Menu {
	t.releaseOnClose = true
	return t
}

func (t *Menu) Show(anchor View) {
	wnd := getWindow(anchor)
	if wnd == nil {
		log.Println("cannot show dialog, anchor is gone")
		return
	}

	/*this.anchor.getElement().style.position = "absolute";
	this.anchor.getElement().style.left = (component.getElement().getBoundingClientRect().left + window.scrollX) + "px";
	if (below) {
		this.anchor.getElement().style.top = (component.getElement().getBoundingClientRect().bottom + window.scrollY) + "px";
	} else {
		this.anchor.getElement().style.top = (component.getElement().getBoundingClientRect().top + window.scrollY) + "px";
	}*/

	//t.fnd.Unwrap().Call("setAbsolutePosition")
	//t.fnd.Unwrap().Call("setAnchorCorner", anchor.node().Unwrap())
	t.fnd.Unwrap().Call("setIsHoisted", true)
	t.fnd.Unwrap().Call("setMenuSurfaceAnchorElement", anchor.node().Unwrap())
	t.fnd.Unwrap().Call("open")

}

func (t *Menu) Style(style ...Style) *Menu {
	t.absComponent.style(style...)
	return t
}

func (t *Menu) AddItem(item *MenuItem) *Menu {
	t.items = append(t.items, item)
	t.menu.AppendChild(item.elem)
	t.absComponent.addResource(item)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Menu) Self(ref **Menu) *Menu {
	*ref = t
	return t
}

func (t *Menu) Release() {
	//log.Println("menu released")

	t.absComponent.Release()
	dom.GetWindow().Document().Body().RemoveChild(t.node())

}

type MenuItem struct {
	parent     *Menu
	elem       dom.Element
	f          func(menu *MenuItem)
	fnd        js2.Foundation
	selectable bool
}

func NewMenuItem(caption string, action func(menu *MenuItem)) *MenuItem {
	elem := dom.CreateElement("li").AddClass("mdc-list-item").SetRole("menuitem")
	item := &MenuItem{
		elem: elem,
		f:    action,
	}
	item.SetCaption(caption)
	item.fnd = js2.Attach(js2.Ripple, elem)
	return item
}

func NewMenuDivider() *MenuItem {
	elem := dom.CreateElement("li").AddClass("mdc-list-divider").SetRole("separator")
	m := &MenuItem{elem: elem}
	return m
}

func (m *MenuItem) SetCaption(caption string) *MenuItem {
	m.elem.SetTextContent("")
	if m.selectable {
		m.elem.AppendChild(dom.CreateElement("span").SetClassName("mdc-list-item__graphic mdc-menu__selection-group-icon"))
	}
	m.elem.AppendChild(dom.CreateElement("span").AddClass("mdc-list-item__text").SetText(caption))
	return m
}

func (m *MenuItem) SetEnabled(b bool) *MenuItem {
	if b {
		m.elem.RemoveClass("mdc-list-item--disabled")
	} else {
		m.elem.AddClass("mdc-list-item--disabled")
	}
	return m
}

func (m *MenuItem) SetSelectListener(f func(menu *MenuItem)) *MenuItem {
	m.f = f
	return m
}

func (m *MenuItem) Menu() *Menu {
	return m.parent
}

func (m *MenuItem) Release() {
	m.fnd.Release()
}
