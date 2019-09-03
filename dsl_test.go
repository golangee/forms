package gowtk

import (
	"fmt"
	"testing"
)

//snip
type MyView struct {
	counter IntProperty

	statusText TextView
	button     Button
	VBox
}

func (v *MyView) Attach() {
	v.AddView(
		NewTextView().
			SetContent("MY COMPONENT").
			SetFontSize(24),
		v.statusText.
			SetContent("initial title text").
			SetPadding(4, 4, 4, 4),
		v.button.
			SetCaption("click me").
			OnClick(func() {
				v.statusText.SetContent("clicked: " + v.counter.Inc().String())
				addedView := NewTextView().SetContent("added view")
				v.AddView(addedView)
			}),
	)
}

func NewMyView() *MyView {
	return &MyView{}
}

//snap

type MyView2 struct {
	MyView1 MyView
	MyView2 MyView
	VBox
}

func (v *MyView2) Attach() {
	v.AddView(&v.MyView1, &v.MyView2)
}

func TestMyView(t *testing.T) {
	var view View = NewMyView()
	view.Attach()
	fmt.Printf("%#v\n", view)

	var root View = &MyView2{}
	root.Attach()
	fmt.Printf("%#v\n", root)
}

/*
func Example() {


	var SettingsView ViewSpec = func(ctx View) {
		counter := IntProperty(0)
		title := StringProperty{}
		saveHistory := Bool(false)

		getText := func() string {
			return "clicked"
		}

		Text(ctx).
			SetContent(title).
			SetPadding(4, 4, 4, 4).
			OnClick(func() {
				title.SetString(getText() + " " + counter.Inc().StringProperty())
				saveHistory.SetBool(true)
			})
	}

	// create a new SettingsView and attach it to the parent
	SettingsView(nil)

}
*/
