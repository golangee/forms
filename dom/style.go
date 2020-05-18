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

package dom

import "syscall/js"

type Style struct {
	val js.Value
}

func (s Style) SetColor(v string) {
	s.val.Set("color", v)
}

func (s Style) SetPaddingLeft(v string) {
	s.val.Set("padding-left", v)
}

func (s Style) SetPaddingTop(v string) {
	s.val.Set("padding-top", v)
}

func (s Style) SetPaddingBottom(v string) {
	s.val.Set("padding-bottom", v)
}

func (s Style) SetPadding(v string) {
	s.val.Set("padding", v)
}

func (s Style) SetMargin(v string) {
	s.val.Set("margin", v)
}

func (s Style) SetMarginBottom(v string) {
	s.val.Set("margin-bottom", v)
}

func (s Style) SetMarginTop(v string) {
	s.val.Set("margin-top", v)
}

func (s Style) Unwrap() js.Value {
	return s.val
}

func (s Style) Set(k, v string) {
	s.val.Set(k, v)
}

func (s Style) SetProperty(k, v string) {
	s.val.Call("setProperty", k, v)
}
