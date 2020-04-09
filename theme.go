package wtk

import "github.com/worldiety/wtk/dom"

type ThemeController struct {
	root dom.Element
}

func (t ThemeController) SetPrimaryColor(color Color) {
	t.root.Style().SetProperty("--mdc-theme-primary", color.String())
	t.root.Style().SetProperty("--wtk-primary-alpha", color.SetAlpha(222).String())
}

func Theme() ThemeController {
	return ThemeController{root: dom.GetWindow().Document().DocumentElement()}
}
