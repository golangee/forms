package wtk

import (
	h "github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/theme/material/icon"
	js2 "github.com/worldiety/wtk/theme/material/js"
	"syscall/js"
)

type TabView struct {
	*absComponent
	tabBarRoot  h.Element
	tabBar      h.Element
	tabScroller h.Element
	contentPane h.Element
	tabs        []*Tab
	fnd         js2.Foundation
	fndScr      js2.Foundation
	activeIdx   int
}

func NewTabView() *TabView {
	t := &TabView{}
	t.absComponent = newComponent(t, "div")

	h.Wrap(t.node(),
		h.Div(h.Class("mdc-tab-bar"), h.Role("tablist"),
			h.Div(h.Class("mdc-tab-scroller"),
				h.Div(h.Class("mdc-tab-scroller__scroll-area"),
					h.Div(h.Class("mdc-tab-scroller__scroll-content")).Self(&t.tabBar),
				),
			).Self(&t.tabScroller),
		).Self(&t.tabBarRoot),
		h.Div().Self(&t.contentPane),
	)
	t.fnd = js2.Attach(js2.TabBar, t.tabBarRoot)

	t.addResource(t.tabBarRoot.AddEventListener("MDCTabBar:activated", func(this js.Value, args []js.Value) interface{} {
		t.activeIdx = args[0].Get("detail").Get("index").Int()
		t.SetActive(t.activeIdx)
		return nil
	}, false))
	return t
}

func (t *TabView) SetScrollable(b bool) *TabView {
	t.fndScr.Release()
	if b {
		t.fndScr = js2.Attach(js2.TabScroller, t.tabScroller)
	}
	return t
}

func (t *TabView) SetTabs(tabs ...*Tab) *TabView {
	t.fnd.Release()

	t.tabBar.Clear()
	t.contentPane.Clear()
	t.tabs = tabs
	for _, tab := range tabs {
		t.tabBar.AppendChild(tab.caption.node())
		tab.caption.attach(t)

		tab.body.node().Style().Set("display", "none")
		t.contentPane.AppendChild(tab.body.node())
		tab.body.attach(t)
	}

	t.fnd = js2.Attach(js2.TabBar, t.tabBarRoot)
	if len(tabs) > 0 {
		t.SetActive(0)
	}
	return t
}

func (t *TabView) SetActive(idx int) *TabView {
	t.activeIdx = idx
	for i, tab := range t.tabs {
		tab.caption.SetActive(i == idx)
		if i == idx {
			tab.body.node().Style().Set("display", "block")
		} else {
			tab.body.node().Style().Set("display", "none")
		}
	}
	return t
}

func (t *TabView) Style(style ...Style) *TabView {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *TabView) Self(ref **TabView) *TabView {
	*ref = t
	return t
}

func (t *TabView) Release() {
	t.fnd.Release()
	t.fndScr.Release()
	t.absComponent.Release()
}

type Tab struct {
	caption *tabHeaderView
	body    View
}

func NewTab(label string, body View) *Tab {
	return &Tab{
		caption: newTabHeaderView().SetLabel(label),
		body:    body,
	}
}

func NewTabWithIcon(ico icon.Icon, label string, body View) *Tab {
	return &Tab{
		caption: newTabHeaderView().SetLabel(label).SetIcon(ico),
		body:    body,
	}
}

func NewTabWithStackedIcon(ico icon.Icon, label string, body View) *Tab {
	return &Tab{
		caption: newTabHeaderView().SetLabel(label).SetIcon(ico).SetStacked(true),
		body:    body,
	}
}

type tabHeaderView struct {
	*absComponent
	label        h.Element
	tabIndicator h.Element
	ico          h.Element
}

func newTabHeaderView() *tabHeaderView {
	t := &tabHeaderView{}
	t.absComponent = newComponent(t, "button")
	h.Wrap(t.node(), h.Class("mdc-tab"), h.Role("tab"),
		h.Span(h.Class("mdc-tab__content"),
			h.Span(h.Class("mdc-tab__icon", "material-icons")).Self(&t.ico),
			h.Span(h.Class("mdc-tab__text-label")).Self(&t.label),
		),
		h.Span(h.Class("mdc-tab-indicator"),
			h.Span(h.Class("mdc-tab-indicator__content", "mdc-tab-indicator__content--underline")),
		).Self(&t.tabIndicator),
		h.Span(h.Class("mdc-tab__ripple")),
	)
	t.ico.Style().Set("display", "none")
	return t
}

func (t *tabHeaderView) SetIcon(ico icon.Icon) *tabHeaderView {
	if len(ico) > 0 {
		t.ico.Style().Set("display", "block")
	} else {
		t.ico.Style().Set("display", "none")
	}
	t.ico.SetTextContent(string(ico))
	return t
}

func (t *tabHeaderView) SetStacked(b bool) *tabHeaderView {
	t.node().RemoveClass("mdc-tab--stacked")
	if b {
		t.node().AddClass("mdc-tab--stacked")
	}
	return t
}

func (t *tabHeaderView) SetLabel(s string) *tabHeaderView {
	t.label.SetTextContent(s)
	return t
}

func (t *tabHeaderView) SetActive(b bool) {
	t.node().RemoveClass("mdc-tab--active")
	t.tabIndicator.RemoveClass("mdc-tab-indicator--active")
	if b {
		t.node().AddClass("mdc-tab--active")
		t.tabIndicator.AddClass("mdc-tab-indicator--active")
	}
}
