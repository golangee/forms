package wtk

import "github.com/worldiety/wtk/dom"

type Drawer struct {
	*absComponent
}

func NewDrawer(bar *TopAppBar, content View) *Drawer {
	t := &Drawer{}
	t.absComponent = newComponent(t, "div")
	t.node().Style().Set("display", "flex") // Removes top gap between modal drawer and viewport
	t.node().Style().Set("height", "100vh") // Ensures that permanent drawer extends to bottom of viewport
	t.node().Style().Set("margin", "0")     // Removes top gap between top app bar and viewport

	t.node().AppendChild(dom.CreateElement("div").SetClassName("behindSideMenu"))

	aside := dom.CreateElement("aside").SetClassName("mdc-drawer")
	aside.SetInnerHTML(`
      <div class="mdc-drawer__content">
        <nav class="mdc-list">
          <a class="mdc-list-item mdc-list-item--selected" href="#" aria-selected="true" tabindex="0">
            <i class="material-icons mdc-list-item__graphic" aria-hidden="true">inbox</i>
            <span class="mdc-list-item__text">Inbox</span>
          </a>
          <a class="mdc-list-item" href="#">
            <i class="material-icons mdc-list-item__graphic" aria-hidden="true">send</i>
            <span class="mdc-list-item__text">Outgoing</span>
          </a>
          <a class="mdc-list-item" href="#">
            <i class="material-icons mdc-list-item__graphic" aria-hidden="true">drafts</i>
            <span class="mdc-list-item__text">Drafts</span>
          </a>
        </nav>
      </div>
`)

	t.node().AppendChild(aside)

	t.node().AppendChild(dom.CreateElement("div").SetClassName("mdc-drawer-scrim"))

	drawerCnt := dom.CreateElement("div").SetClassName("mdc-drawer-app-content")
	drawerCnt.AppendChild(bar.node())
	t.addResource(bar)
	// no attach?
	t.node().AppendChild(drawerCnt)

	mainCnt := dom.CreateElement("div").SetClassName("main-content")
	fixAdjust := dom.CreateElement("div").SetClassName("mdc-top-app-bar--fixed-adjust")
	mainCnt.AppendChild(fixAdjust)
	fixAdjust.AppendChild(content.node())
	t.addResource(content)
	// no attach?
	t.node().AppendChild(mainCnt)

	return t
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
