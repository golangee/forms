package wtk

import (
	"github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/theme/material/icon"
	js2 "github.com/worldiety/wtk/theme/material/js"
	"syscall/js"
)

type TopAppBar struct {
	*absComponent
	style *topAppBarStyleStandard
}

func NewTopAppBar() *TopAppBar {
	t := &TopAppBar{}
	t.absComponent = newComponent(t, "header")
	t.style = newTopAppBarStyleStandard(t)
	return t
}

func (t *TopAppBar) Style(style ...Style) *TopAppBar {
	t.absComponent.style(style...)
	return t
}

func (t *TopAppBar) SetNavigation(i icon.Icon, action func(view View)) *TopAppBar {
	t.style.setNavigationIcon(i, action)
	return t
}

func (t *TopAppBar) SetTitle(str string) *TopAppBar {
	t.style.setTitle(str)
	return t
}

func (t *TopAppBar) AddActions(items ...*IconItem) *TopAppBar {
	t.style.addActions(items...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *TopAppBar) Self(ref **TopAppBar) *TopAppBar {
	*ref = t
	return t
}

// topAppBarStyleStandard is only fixed, because the standard is broken with scrolling. Dunno why
type topAppBarStyleStandard struct {
	header           dom.Element
	row              dom.Element
	left             dom.Element
	right            dom.Element
	navigationIcon   icon.Icon
	navigationAction func(v View)
	title            string
	parent           *TopAppBar
	actions          []*IconItem
	fnd              js2.Foundation
}

//  <header class="mdc-top-app-bar">
//  <div class="mdc-top-app-bar__row">
//    <section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-start">
//      <button class="mdc-icon-button material-icons mdc-top-app-bar__navigation-icon--unbounded">menu</button><span class="mdc-top-app-bar__title">San Francisco, CA</span> </section>
//    <section class="mdc-top-app-bar__section mdc-top-app-bar__section--align-end">
//      <button class="mdc-icon-button material-icons mdc-top-app-bar__action-item--unbounded" aria-label="Download">file_download</button>
//      <button class="mdc-icon-button material-icons mdc-top-app-bar__action-item--unbounded" aria-label="Print this page">print</button>
//      <button class="mdc-icon-button material-icons mdc-top-app-bar__action-item--unbounded" aria-label="Bookmark this page">bookmark</button>
//    </section>
//  </div>
//  </header>
func newTopAppBarStyleStandard(parent *TopAppBar) *topAppBarStyleStandard {
	t := &topAppBarStyleStandard{parent: parent}
	t.header = t.parent.node().SetClassName("mdc-top-app-bar mdc-top-app-bar--fixed material-icons")
	t.row = dom.CreateElement("div").AddClass("mdc-top-app-bar__row")
	t.header.AppendChild(t.row)
	t.left = dom.CreateElement("section").SetClassName("mdc-top-app-bar__section mdc-top-app-bar__section--align-start")
	t.right = dom.CreateElement("section").SetClassName("mdc-top-app-bar__section mdc-top-app-bar__section--align-end")

	t.row.AppendChild(t.left)
	t.row.AppendChild(t.right)

	return t
}

func (t *topAppBarStyleStandard) rebuild() {
	t.fnd.Release()
	t.left.SetTextContent("")
	if len(t.navigationIcon) > 0 {
		btn := dom.CreateElement("button").SetClassName("mdc-icon-button material-icons mdc-top-app-bar__action-item").SetText(string(t.navigationIcon))
		t.parent.addResource(btn.AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
			if t.navigationAction != nil {
				t.navigationAction(t.parent)
			}
			return nil
		}, false))
		t.left.AppendChild(btn)
	}

	if len(t.title) > 0 {
		t.left.AppendChild(dom.CreateElement("span").AddClass("mdc-top-app-bar__title").SetText(t.title))
	}

	for _, action := range t.actions {
		fAction := action
		btn := dom.CreateElement("button").SetClassName("mdc-icon-button material-icons mdc-top-app-bar__action-item").SetText(string(action.ico)).SetAriaLabel(action.label)
		t.parent.addResource(btn.AddEventListener("click", func(this js.Value, args []js.Value) interface{} {
			if fAction.action != nil {
				fAction.action(t.parent)
			}
			return nil
		}, false))
		t.right.AppendChild(btn)
	}

	t.fnd = js2.Attach(js2.TopAppBar, t.header)

	//new mdc.topAppBar.MDCTopAppBar.attachTo(topAppBarElement);
	/*js.Global().
	Get("mdc").
	Get("topAppBar").
	Get("MDCTopAppBar").
	Get("attachTo").
	New(t.header.Unwrap())*/

	t.parent.addResource(t.fnd)
}

func (t *topAppBarStyleStandard) setNavigationIcon(i icon.Icon, action func(v View)) {
	t.navigationIcon = i
	t.navigationAction = action
	t.rebuild()
}

func (t *topAppBarStyleStandard) setTitle(s string) {
	t.title = s
	t.rebuild()
}

func (t *topAppBarStyleStandard) addActions(item ...*IconItem) {
	t.actions = append(t.actions, item...)
	t.rebuild()
}

type IconItem struct {
	ico    icon.Icon
	label  string
	action func(v View)
}

func NewIconItem(ico icon.Icon, label string, action func(v View)) *IconItem {
	return &IconItem{
		ico:    ico,
		label:  label,
		action: action,
	}
}
