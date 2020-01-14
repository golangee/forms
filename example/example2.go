package example

import . "github.com/worldiety/gowtk"

type ContentView2 struct {
	body *VBox
}

func NewContentView2() *ContentView2 {
	return &ContentView2{
		body: NewVBox(
			NewText("hello world").
				SetFont(Title).
				SetForegroundColor(Green),
			NewHBox(
				NewText("National Park").
					SetFont(SubHeadline),
				NewSpacer(),
				NewText("Lower Saxony").
					SetFont(SubHeadline),
			),
		).SetAlignment(Leading).
			SetPadding(DefaultPadding),
	}
}
