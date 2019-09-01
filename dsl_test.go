package gowtk

import "testing"

//snip
type MyView struct {
	counter IntProperty

	statusText TextView
	button     Button
	body       VBox
}

func (m *MyView) Children() []View {
	return []View{&m.body}
}

func NewMyView() *MyView {
	view := &MyView{}
	view.body.Layout(
		NewTextView().
			SetContent("MY COMPONENT").
			SetFontSize(24),
		view.statusText.
			SetContent("initial title text").
			SetPadding(4, 4, 4, 4),
		view.button.
			SetCaption("click me").
			OnClick(func() {
				view.statusText.SetContent("clicked: " + view.counter.Inc().String())
			}),
	)
	return view
}

//snap

func TestMyView(t *testing.T) {
	NewMyView()
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
