package forms

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golangee/forms/theme/material/icon"
	"net/url"
)

type ErrorView struct {
	*VStack
	body *Group
	err  error
	hide bool
	btn *Button
}

func NewErrorView(err error) *ErrorView {
	view := &ErrorView{}
	view.err = err
	view.VStack = NewVStack().AddViews(
		NewText(view.guestimateHelp(err)).Style(Font(Headline5)),
		NewGroup().Self(&view.body),
		NewHStack().AddViews(
			NewButton("Details").Self(&view.btn).SetLeadingIcon(icon.ArrowDropDown).
				AddClickListener(func(v View) {
					view.body.RemoveAll()
					if view.err != nil && !view.hide {
						view.hide = true
						view.body.AddViews(newDetailErrorView(view.err))
						view.btn.SetLeadingIcon(icon.ArrowDropUp)
					} else {
						view.hide = false
						view.btn.SetLeadingIcon(icon.ArrowDropDown)
					}
				}),
			NewButton("reload").AddClickListener(func(v View) {
				v.Context().router().Reload(false)
			}),
		),
	)
	return view
}

func (t *ErrorView) guestimateHelp(err error) string {
	var urlError *url.Error
	if errors.As(err, &urlError) {
		return "Please check your internet connection and try again."
	}

	return "Something went wrong. If this problem persists, please contact your support."
}

type detailErrorView struct {
	*List
}

func newDetailErrorView(err error) *detailErrorView {
	view := &detailErrorView{}
	view.List = NewList()
	vg := view.List
	vg.AddItems(NewListHeader("error trace"))
	vg.AddItems(NewListSeparator())
	root := err
	for root != nil {
		msg := root.Error()
		localizedMsg := ""
		if obj, ok := root.(interface{ LocalizedError() string }); ok {
			localizedMsg = obj.LocalizedError()
		}

		var causedBy error
		if obj, ok := root.(interface{ Unwrap() error }); ok {
			causedBy = obj.Unwrap()
		}

		id := ""
		if obj, ok := root.(interface{ ID() string }); ok {
			id = obj.ID()
		}

		class := ""
		if obj, ok := root.(interface{ Class() string }); ok {
			class = obj.Class()
		}

		var payload interface{}
		if obj, ok := root.(interface{ Payload() interface{} }); ok {
			payload = obj.Payload()
		}

		if localizedMsg == "" {
			localizedMsg = msg
		}

		if id == "" {
			id = class
		}

		typ := id
		if id != class {
			typ += " (" + class + ")"
		}

		vg.AddItems(NewListTwoLineItem(msg, typ).SetLeadingView(NewIcon(icon.Error)))
		if payload != nil {
			vg.AddItems(NewListItem(stuffToString(payload)).SetLeadingView(NewIcon(icon.Details)))
		}

		if causedBy != nil {
			vg.AddItems(NewListSeparator())
			vg.AddItems(NewListHeader("caused by"))
		}
		root = causedBy
	}

	return view
}

func stuffToString(i interface{}) string {
	buf, err := json.MarshalIndent(i, " ", " ")
	if err != nil {
		return fmt.Sprintf("%+v", i)
	}
	return string(buf)
}
