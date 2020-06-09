// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package forms

import (
	"github.com/golangee/forms/dom"
	"syscall/js"
)

type ThemeController struct {
	root dom.Element
}

// Color sets the current primary theme background color
func (t ThemeController) SetColor(color Color) {
	t.root.Style().SetProperty("--mdc-theme-primary", color.String())
	t.root.Style().SetProperty("--mdc-theme-secondary", color.String())
	t.root.Style().SetProperty("--wtk-primary-alpha", color.SetAlpha(222).String())
}

// Color returns the current primary theme background color
func (t ThemeController) Color() Color {
	val := js.Global().Call("getComputedStyle", t.root.Unwrap())
	val = val.Call("getPropertyValue", "--mdc-theme-primary")

	return ParseColor(val.String())
}

// SetForegroundColor sets the current primary theme foreground color
func (t ThemeController) SetForegroundColor(color Color) {
	t.root.Style().SetProperty("--mdc-theme-on-primary", color.String())
}

// ForegroundColor returns the current primary theme foreground color
func (t ThemeController) ForegroundColor() Color {
	val := js.Global().Call("getComputedStyle", t.root.Unwrap())
	val = val.Call("getPropertyValue", "--mdc-theme-on-primary")

	return ParseColor(val.String())
}

func Theme() ThemeController {
	return ThemeController{root: dom.GetWindow().Document().DocumentElement()}
}
