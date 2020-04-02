package wtk

import (
	"github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/event"
	"github.com/worldiety/wtk/theme/material/icon"
	"github.com/worldiety/wtk/theme/material/js"
)

type StyleKind string

const Default StyleKind = ""
const Raised StyleKind = "--raised"
const Unelevated StyleKind = "--unelevated"
const Outlined StyleKind = "--outlined"

type Alignment string

const Leading Alignment = "leading"
const Trailing Alignment = "trailing"

type Button struct {
	Text  string
	label dom.Element
	*absComponent
}

func NewButton(text string) *Button {
	t := &Button{}
	t.absComponent = newComponent(t, "button")
	t.node().AddClass("mdc-button")
	t.node().AppendChild(dom.CreateElement("div").AddClass("mdc-button__ripple"))
	t.label = dom.CreateElement("span").AddClass("mdc-button__label")

	t.node().AppendChild(t.label)

	t.SetText(text)
	t.absComponent.addResource(js.Attach(js.Ripple, t.node()))
	return t
}

func (t *Button) AddIcon(icon icon.Icon, alignment Alignment) *Button {
	switch alignment {
	case Leading:
		t.node().InsertBefore(dom.CreateElement("i").SetClassName("material-icons mdc-button__icon").SetInnerText(string(icon)), t.label)
	case Trailing:
		t.node().AppendChild(dom.CreateElement("i").SetClassName("material-icons mdc-button__icon").SetInnerText(string(icon)))
	default:
		panic("unsupported alignment: " + alignment)
	}
	return t
}

func (t *Button) SetStyleKind(s StyleKind) *Button {
	t.node().SetClassName("mdc-button mdc-button" + string(s))
	return t
}

func (t *Button) SetEnabled(b bool) *Button {
	t.node().SetDisabled(!b)
	return t
}

func (t *Button) SetText(str string) *Button {
	t.Text = str
	t.label.SetInnerText(str)
	return t
}

func (t *Button) Style(style ...Style) *Button {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given reference
func (t *Button) Self(ref **Button) *Button {
	*ref = t
	return t
}

func (t *Button) AddClickListener(f func(v View)) *Button {
	t.addEventListener(event.Click, f)
	return t
}
