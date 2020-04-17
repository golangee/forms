package wtk

import (
	"github.com/worldiety/wtk/dom"
	js2 "github.com/worldiety/wtk/theme/material/js"
	"log"
	"syscall/js"
)

const drawerModeModal = "modal"
const drawerModePermanent = "permanent"

type Drawer struct {
	*absComponent
	drawer      dom.Element
	mainContent dom.Element
	drawerFnd   js2.Foundation
	topAppBar   *TopAppBar
	navList     dom.Element
	maxWidth    string
	mode        string
}

func NewDrawer(bar *TopAppBar, header View, menu *List, content View) *Drawer {
	t := &Drawer{}
	t.topAppBar = bar
	t.topAppBar.style.releaseFoundation()
	t.topAppBar.style.initFoundation()
	t.maxWidth = "900px"

	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("display", "flex") // Removes top gap between modal drawer and viewport
	t.node().Style().Set("min-height", "100vh") // Ensures that permanent drawer extends to bottom of viewport
	t.node().Style().Set("margin", "0")     // Removes top gap between top app bar and viewport

	t.node().AppendChild(dom.CreateElement("div").SetClassName("behindSideMenu"))

	drawerCnt := dom.CreateElement("div").SetClassName("mdc-drawer-app-content")
	drawerCnt.AppendChild(bar.node())
	t.addResource(bar)
	bar.attach(t)
	t.node().AppendChild(drawerCnt)

	t.drawer = dom.CreateElement("aside").SetClassName("wtk-drawer mdc-drawer")
	drawerHeader := dom.CreateElement("div").SetClassName("mdc-drawer__header")
	//	drawerHeader.SetInnerHTML(`<h3 class="mdc-drawer__title">Mail</h3>
	//   <h6 class="mdc-drawer__subtitle">email@material.io</h6>`)
	drawerHeader.AppendChild(header.node())
	header.attach(t)

	drawerContent := dom.CreateElement("div").SetClassName("mdc-drawer__content")

	t.navList = menu.node()
	menu.attach(t)
	drawerContent.AppendChild(t.navList)

	t.drawer.AppendChild(drawerHeader)
	t.drawer.AppendChild(drawerContent)

	t.node().AppendChild(t.drawer)

	t.node().AppendChild(dom.CreateElement("div").SetClassName("mdc-drawer-scrim"))

	t.mainContent = dom.CreateElement("div").SetClassName("main-content")
	fixAdjust := dom.CreateElement("div").SetClassName("mdc-top-app-bar--fixed-adjust")
	t.mainContent.AppendChild(fixAdjust)
	fixAdjust.AppendChild(content.node())
	t.addResource(content)
	content.attach(t)
	t.node().AppendChild(t.mainContent)

	t.initResponsiveLogic()
	return t
}

func (t *Drawer) initModalDrawer() js2.Foundation {
	t.drawer.AddClass("mdc-drawer--modal")
	fnd := js2.Attach(js2.Drawer, t.drawer)
	fnd.Unwrap().Set("open", false)

	//t.topAppBar.style.fnd.Unwrap().Call("setScrollTarget", t.mainContent.Unwrap()) //this breaks the scrolling behavior

	t.addResource(t.topAppBar.node().AddEventListener("MDCTopAppBar:nav", func(this js.Value, args []js.Value) interface{} {
		fnd.Unwrap().Set("open", !fnd.Unwrap().Get("open").Bool())
		return nil
	}, false))

	t.addResource(t.drawer.AddEventListener("MDCDrawer:closed", func(this js.Value, args []js.Value) interface{} {
		log.Println("drawer is closed")
		return nil
	}, false))

	t.addResource(t.navList.AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
		fnd.Unwrap().Set("open", false)
		return nil
	}, false))
	t.mode = drawerModeModal
	return fnd
}

func (t *Drawer) initPermanentDrawer() js2.Foundation {
	t.drawer.RemoveClass("mdc-drawer--modal")
	fnd := js2.Attach(js2.List, t.navList)
	fnd.Unwrap().Set("wrapFocus", true)
	t.mode = drawerModePermanent
	return fnd
}

func (t *Drawer) matchMaxCriteria() string {
	return "(max-width: " + t.maxWidth + ")"
}

func (t *Drawer) matchMinCriteria() string {
	return "(min-width: " + t.maxWidth + ")"
}

func (t *Drawer) initResponsiveLogic() {
	var fnd js2.Foundation
	if dom.GetWindow().MatchesMedia(t.matchMaxCriteria()) {
		fnd = t.initModalDrawer()
	} else {
		fnd = t.initPermanentDrawer()
	}

	t.addResource(dom.GetWindow().AddEventListener("resize", func(this js.Value, args []js.Value) interface{} {
		if dom.GetWindow().MatchesMedia(t.matchMaxCriteria()) && t.mode == drawerModePermanent {
			fnd.Release()
			fnd = t.initModalDrawer()
		} else if dom.GetWindow().MatchesMedia(t.matchMinCriteria()) && t.mode == drawerModeModal {
			fnd.Release()
			fnd = t.initPermanentDrawer()
		}
		return nil
	}, false))
}

func (t *Drawer) Style(style ...Style) *Drawer {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Drawer) Self(ref **Drawer) *Drawer {
	*ref = t
	return t
}

func (t *Drawer) Release() {
	t.absComponent.Release()
	t.drawerFnd.Release()
}
