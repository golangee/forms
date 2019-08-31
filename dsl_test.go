package gowtk

func Example() {
	var SettingsView ViewSpec = func(ctx View) {
		counter := Int(0)
		title := String("hello world")
		saveHistory := Bool(false)

		getText := func() string {
			return "clicked"
		}

		Text(ctx).
			SetContent(title).
			SetPadding(4, 4, 4, 4).
			OnClick(func() {
				title.SetString(getText() + " " + counter.Inc().String())
				saveHistory.SetBool(true)
			})
	}

	// create a new SettingsView and attach it to the parent
	SettingsView(nil)

}
