package example

import . "github.com/worldiety/gowtk"

type ContentView struct {
	body *TextView
}

func NewContentView() *ContentView {
	return &ContentView{
		body: NewText("hello world").
			SetFont(Title).
			SetForegroundColor(Green),
	}
}
