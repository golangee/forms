package wtk

import (
	h "github.com/worldiety/wtk/dom"
	"github.com/worldiety/wtk/theme/material/icon"
	js2 "github.com/worldiety/wtk/theme/material/js"
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
	label          h.Element
	labelNotch     h.Element
	helperLine     h.Element
	helperText     h.Element
	leadingIco     h.Element
	trailingIco    h.Element
	characterCount h.Element
	mdcTextField   h.Element
	input          h.Element
	fndTF          js2.Foundation
	dirty          bool
}

func NewTextField() *TextField {
	t := &TextField{}
	t.absComponent = newComponent(t, "div")
	labelId := nextId()
	h.Wrap(t.node(), h.Class("text-field-container"),
		h.Div(h.Class("mdc-text-field", "mdc-text-field--outlined"), h.Class("mdc-text-field--no-label"),
			h.I(h.Class("material-icons", "mdc-text-field__icon")).Self(&t.leadingIco),
			h.I(h.Class("material-icons", "mdc-text-field__icon")).Self(&t.trailingIco),
			h.Input(h.Class("mdc-text-field__input"), h.Type("text"), h.AriaLabelledby(labelId)).Self(&t.input),
			h.Div(h.Class("mdc-notched-outline"),
				h.Div(h.Class("mdc-notched-outline__leading")),
				h.Div(h.Class("mdc-notched-outline__notch"),
					h.Span(h.Class("mdc-floating-label"), h.Id(labelId)).Self(&t.label),
				).Self(&t.labelNotch),
				h.Div(h.Class("mdc-notched-outline__trailing")),
			),

		).Self(&t.mdcTextField),
		h.Div(h.Class("mdc-text-field-helper-line"),
			h.Div(h.Class("mdc-text-field-helper-text", "mdc-text-field-helper-text--persistent", "mdc-text-field-helper-text--validation-msg")).Self(&t.helperText),

			h.Div(h.Class()).Self(&t.characterCount),
		).Self(&t.helperLine),
	)
	t.labelNotch.Style().Set("display", "none")
	t.invalidate(false)
	return t
}

// invalidate postpones foundation construction until view is attached
func (t *TextField) invalidate(force bool) {
	if t.isAttached() || force {
		t.dirty = false
		t.initFoundation()
	} else {
		t.dirty = true
	}
}

func (t *TextField) initFoundation() {
	t.fndTF.Release()
	t.fndTF = js2.Attach(js2.TextField, t.mdcTextField)
}

func (t *TextField) SetLeadingIcon(ico icon.Icon) *TextField {
	t.mdcTextField.RemoveClass("mdc-text-field--with-leading-icon")
	if len(ico) > 0 {
		t.mdcTextField.AddClass("mdc-text-field--with-leading-icon")
		t.leadingIco.Style().Set("display", "block")
	} else {
		t.leadingIco.Style().Set("display", "none")
	}
	t.leadingIco.SetText(string(ico))
	return t
}

func (t *TextField) SetTrailingIcon(ico icon.Icon) *TextField {
	t.mdcTextField.RemoveClass("mdc-text-field--with-trailing-icon")
	if len(ico) > 0 {
		t.mdcTextField.AddClass("mdc-text-field--with-trailing-icon")
		t.trailingIco.Style().Set("display", "block")
	} else {
		t.trailingIco.Style().Set("display", "none")
	}
	t.trailingIco.SetText(string(ico))
	return t
}

func (t *TextField) SetEnabled(b bool) *TextField {
	t.input.SetDisabled(!b)
	t.mdcTextField.RemoveClass("field--disabled")
	if !b {
		t.mdcTextField.AddClass("field--disabled")
	}
	t.invalidate(false)
	return t
}

func (t *TextField) SetText(str string) *TextField {
	t.input.SetAttr("value", str)
	t.invalidate(false)
	return t
}

func (t *TextField) SetLabel(str string) *TextField {
	t.mdcTextField.RemoveClass("mdc-text-field--no-label")
	if len(str) == 0 {
		t.mdcTextField.AddClass("mdc-text-field--no-label")
		t.labelNotch.Style().Set("display", "none")
	} else {
		t.labelNotch.Style().Set("display", "block")
	}
	t.label.SetText(str)
	t.invalidate(false)
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
		s.applyCSS(t.mdcTextField)
		_ = s
	}
	return t
}

func (t *TextField) SetInputType(in InputType) *TextField {
	t.input.SetAttr("type", string(in))
	return t
}

func (t *TextField) SetRange(min, max int) *TextField {
	t.SetInputType(Range)
	t.input.SetAttr("min", min)
	t.input.SetAttr("max", max)
	return t
}

func (t *TextField) SetHelper(str string) *TextField {
	t.helperText.SetTextContent(str)
	return t
}

func (t *TextField) SetHelperPersistent(b bool) *TextField {
	t.helperText.RemoveClass("mdc-text-field-helper-text--persistent")
	if b {
		t.helperText.AddClass("mdc-text-field-helper-text--persistent")
	}
	return t
}

func (t *TextField) SetMaxLength(chars int) *TextField {
	t.characterCount.SetClassName("mdc-text-field-character-counter")
	t.input.SetAttr("maxLength", chars)
	t.invalidate(false)
	return t
}

func (t *TextField) SetInvalid(b bool) *TextField {
	t.SetHelperPersistent(b)
	t.mdcTextField.AddClass("mdc-text-field--invalid")
	if !t.fndTF.IsValid() {
		t.invalidate(true)
	}
	t.fndTF.Unwrap().Set("useNativeValidation", false)
	t.fndTF.Unwrap().Set("valid", !b)
	return t
}

// Self assigns the receiver to the given reference
func (t *TextField) Self(ref **TextField) *TextField {
	*ref = t
	return t
}

func (t *TextField) SetRequired(b bool) *TextField {
	if b {
		t.input.Unwrap().Set("required", "required")
	} else {
		t.input.Unwrap().Delete("required")
	}
	return t
}

func (t *TextField) attach(v View) {
	t.absComponent.attach(v)
	if t.dirty {
		t.dirty = false
		t.invalidate(true)
	}
}

func (t *TextField) Release() {
	t.fndTF.Release()
	t.absComponent.Release()
}
