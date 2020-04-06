package wtk

import (
	"github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/theme/material/icon"
	"github.com/worldiety/wtk/theme/material/js"
)

type TextField struct {
	*absComponent
	layoutCtr textFieldLayoutController
}

func NewTextField() *TextField {
	t := &TextField{}
	t.absComponent = newComponent(t, "div")
	t.layoutCtr = newTFFilled(t.node())
	t.absComponent.addResource(t.layoutCtr)
	return t
}

func (t *TextField) SetLeadingIcon(ico icon.Icon) *TextField {
	t.layoutCtr.setLeadingIcon(ico)
	return t
}

func (t *TextField) SetTrailingIcon(ico icon.Icon) *TextField {
	t.layoutCtr.setTrailingIcon(ico)
	return t
}

func (t *TextField) SetEnabled(b bool) *TextField {
	t.layoutCtr.setEnabled(b)
	return t
}

func (t *TextField) SetText(str string) *TextField {
	t.layoutCtr.setInput(str)
	return t
}

func (t *TextField) SetLabel(str string) *TextField {
	t.layoutCtr.setLabel(str)
	return t
}

func (t *TextField) Style(style ...Style) *TextField {
	t.absComponent.style(style...)
	return t
}

func (t *TextField) SetPassword(b bool) *TextField {
	t.layoutCtr.setPassword(b)
	return t
}

func (t *TextField) SetHelper(str string) *TextField {
	t.layoutCtr.setHelper(str)
	return t
}

// Self assigns the receiver to the given reference
func (t *TextField) Self(ref **TextField) *TextField {
	*ref = t
	return t
}

type textFieldLayoutController interface {
	setInput(val string)
	setLabel(val string)
	setTrailingIcon(i icon.Icon)
	setLeadingIcon(i icon.Icon)
	setEnabled(b bool)
	setPassword(b bool)
	setHelper(str string)
	Release()
}

//  <div class="mdc-text-field mdc-text-field--with-leading-icon mdc-text-field--with-trailing-icon">
//    <i class="material-icons mdc-text-field__icon">favorite</i>
//    <i class="material-icons mdc-text-field__icon">visibility</i>
//    <input class="mdc-text-field__input" id="text-field-hero-input">
//    <div class="mdc-line-ripple"></div>
//    <label for="text-field-hero-input" class="mdc-floating-label">lorem ipsum</label>
//  </div>
type tfFilled struct {
	div               dom.Element
	icon1             dom.Element
	icon2             dom.Element
	input             dom.Element
	lineRipple        dom.Element
	label             dom.Element
	helperDiv         dom.Element
	helperText        dom.Element
	valueLeadingIcon  icon.Icon
	valueTrailingIcon icon.Icon
	valueLabel        string
	valueInput        string
	valueHelper       string
	foundation        js.Foundation
}

func newTFFilled(parentDiv dom.Element) *tfFilled {
	t := &tfFilled{div: parentDiv}

	t.icon1 = dom.CreateElement("i").SetClassName("material-icons mdc-text-field__icon")
	t.icon2 = dom.CreateElement("i").SetClassName("material-icons mdc-text-field__icon")
	t.input = dom.CreateElement("input").AddClass("mdc-text-field__input")
	t.input.SetId(nextId())
	t.lineRipple = dom.CreateElement("div").AddClass("mdc-line-ripple")
	t.label = dom.CreateElement("label").SetFor(t.input.Id()).AddClass("mdc-floating-label")

	t.helperText = dom.CreateElement("div").AddClass("mdc-text-field-helper-text")
	t.helperDiv = dom.CreateElement("div").AddClass("mdc-text-field-helper-line")
	t.helperDiv.AppendChild(t.helperText)

	t.apply()
	return t
}

func (t *tfFilled) apply() {
	t.foundation.Release()

	t.div.SetTextContent("")
	t.div.SetClassName("mdc-text-field")

	if len(t.valueLeadingIcon) > 0 {
		t.div.AddClass("mdc-text-field--with-leading-icon")
		t.icon1.SetTextContent(string(t.valueLeadingIcon))
		t.div.AppendChild(t.icon1)
	}

	if len(t.valueTrailingIcon) > 0 {
		t.div.AddClass("mdc-text-field--with-trailing-icon")
		t.icon2.SetTextContent(string(t.valueTrailingIcon))
		t.div.AppendChild(t.icon2)
	}
	t.setInput(t.valueInput)
	t.div.AppendChild(t.input)

	t.div.AppendChild(t.lineRipple)

	if len(t.valueLabel) > 0 {
		t.label.SetTextContent(t.valueLabel)
		t.div.AppendChild(t.label)
	} else {
		t.div.AddClass("mdc-text-field--no-label")
	}

	if len(t.valueHelper) > 0 {
		t.helperText.SetTextContent(t.valueHelper)
		t.div.AppendChild(t.helperDiv)
	}

	t.foundation = js.Attach(js.TextField, t.div)
}

// setInput fixes FOUC label
func (t *tfFilled) setInput(val string) {
	t.label.SetClassName("mdc-floating-label")
	if len(val) > 0 {
		t.label.AddClass("mdc-floating-label--float-above")
	}
	t.input.Unwrap().Set("value", val)
	t.valueInput = val
}

func (t *tfFilled) setLabel(val string) {
	t.valueLabel = val
	t.apply()
}

func (t *tfFilled) setTrailingIcon(i icon.Icon) {
	t.valueTrailingIcon = i
	t.apply()
}

func (t *tfFilled) setLeadingIcon(i icon.Icon) {
	t.valueLeadingIcon = i
	t.apply()
}

func (t *tfFilled) Release() {
	t.foundation.Release()
}

func (t *tfFilled) setEnabled(b bool) {
	t.input.SetDisabled(!b)
	if b {
		t.div.RemoveClass("mdc-text-field--disabled")
	} else {
		t.div.AddClass("mdc-text-field--disabled")
	}
}

func (t *tfFilled) setPassword(b bool) {
	if b {
		t.input.Unwrap().Set("type", "password")
	} else {
		t.input.Unwrap().Set("type", "text")
	}
}

func (t *tfFilled) setHelper(str string) {
	t.valueHelper = str
	t.apply()
}
