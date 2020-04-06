package wtk

import (
	"github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/theme/material/icon"
	"github.com/worldiety/wtk/theme/material/js"
)

type InputType string

const Password InputType = "password"
const Txt InputType = "text"
const EMail InputType = "email"
const URL InputType = "url"
const Range InputType = "range"
const Number InputType = "number"
const Date InputType = "date"
const DateTime InputType = "datetime"
const DateTimeLocal InputType = "datetime-local"
const Month InputType = "month"
const Search InputType = "search"
const Tel InputType = "tel"
const Time InputType = "time"
const Week InputType = "week"

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

// Styles changes the container
func (t *TextField) Style(style ...Style) *TextField {
	t.absComponent.style(style...)
	return t
}

// Styles changes the
func (t *TextField) InputStyle(styles ...Style) *TextField {
	for _, s := range styles {
		s.applyCSS(t.layoutCtr.mdcTextField())
	}
	return t
}

func (t *TextField) SetInputType(in InputType) *TextField {
	t.layoutCtr.inputField().Unwrap().Set("type", string(in))
	return t
}

func (t *TextField) SetRange(min, max int) *TextField {
	t.SetInputType(Range)
	t.layoutCtr.inputField().Unwrap().Set("min", min)
	t.layoutCtr.inputField().Unwrap().Set("max", max)
	return t
}

func (t *TextField) SetHelper(str string) *TextField {
	t.layoutCtr.setHelper(str)
	return t
}

func (t *TextField) SetMaxLength(chars int) *TextField {
	t.layoutCtr.setMaxLength(chars)
	return t
}

func (t *TextField) SetInvalid(b bool) *TextField {
	if b {
		t.layoutCtr.mdcTextField().AddClass("mdc-text-field--invalid")
	} else {
		t.layoutCtr.mdcTextField().RemoveClass("mdc-text-field--invalid")
	}
	return t
}

// Self assigns the receiver to the given reference
func (t *TextField) Self(ref **TextField) *TextField {
	*ref = t
	return t
}

func (t *TextField) SetRequired(b bool) *TextField {
	if b {
		t.layoutCtr.inputField().Unwrap().Set("required", "required")
	} else {
		t.layoutCtr.mdcTextField().Unwrap().Delete("required")
	}
	return t
}

type textFieldLayoutController interface {
	setInput(val string)
	setLabel(val string)
	setTrailingIcon(i icon.Icon)
	setLeadingIcon(i icon.Icon)
	setEnabled(b bool)
	setHelper(str string)
	setMaxLength(c int)
	mdcTextField() dom.Element
	inputField() dom.Element
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
	container         dom.Element
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
	maxLen            int
	foundation        js.Foundation
}

func newTFFilled(parentDiv dom.Element) *tfFilled {
	t := &tfFilled{container: parentDiv}

	t.div = dom.CreateElement("div")
	t.icon1 = dom.CreateElement("i").SetClassName("material-icons mdc-text-field__icon")
	t.icon2 = dom.CreateElement("i").SetClassName("material-icons mdc-text-field__icon")
	t.input = dom.CreateElement("input").AddClass("mdc-text-field__input")
	t.input.SetId(nextId())
	t.lineRipple = dom.CreateElement("div").AddClass("mdc-line-ripple")
	t.label = dom.CreateElement("label").SetFor(t.input.Id()).AddClass("mdc-floating-label")

	t.helperText = dom.CreateElement("div").SetClassName("mdc-text-field-helper-text mdc-text-field-helper-text--persistent mdc-text-field-helper-text--validation-msg")
	t.helperDiv = dom.CreateElement("div").AddClass("mdc-text-field-helper-line")
	t.helperDiv.AppendChild(t.helperText)

	t.apply()
	return t
}

func (t *tfFilled) apply() {
	t.foundation.Release()

	t.container.SetClassName("text-field-container")
	t.container.SetTextContent("")
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

	t.container.AppendChild(t.div)

	if len(t.valueHelper) > 0 {
		t.helperText.SetTextContent(t.valueHelper)
		t.container.AppendChild(t.helperDiv)
	}

	if t.maxLen != 0 {
		t.input.Unwrap().Set("maxLength", t.maxLen)
		t.helperDiv.AppendChild(dom.CreateElement("div").SetClassName("mdc-text-field-character-counter"))
	}

	t.foundation = js.Attach(js.TextField, t.container)
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

func (t *tfFilled) setHelper(str string) {
	t.valueHelper = str
	t.apply()
}

func (t *tfFilled) mdcTextField() dom.Element {
	return t.div
}

func (t *tfFilled) inputField() dom.Element {
	return t.input
}

func (t *tfFilled) setMaxLength(c int) {
	t.maxLen = c
	t.apply()
}
