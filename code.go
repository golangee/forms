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
	"strings"
)

const codeBlue = "wtk-code-blue"
const codeYellow = "wtk-code-yellow"
const codeGreen = "wtk-code-green"
const codeRed = "wtk-code-red"

type Highlighter struct {
	keywords map[string]string
}

var GoSyntax = Highlighter{keywords: map[string]string{
	"func":    codeYellow,
	"type":    codeYellow,
	"import":  codeYellow,
	"package": codeYellow,
	"const":   codeYellow,
	"struct":  codeYellow,
	"string":  codeBlue,
	"*":       codeGreen,
	"&":       codeRed,
	"New":     codeRed,
}}

type Code struct {
	*absComponent
	hl Highlighter
}

func NewCode(hl Highlighter, str string) *Code {
	t := &Code{hl: hl}
	t.absComponent = newComponent(t, "div")
	t.node().AddClass("wtk-code")
	t.Set(str)
	return t
}

func (t *Code) Set(str string) *Code {
	t.absComponent.elem.SetInnerText("")
	for _, line := range strings.Split(str, "\n") {
		pre := dom.CreateElement("pre")
		for _, token := range strings.Split(line, " ") {
			found := false
			for k, color := range t.hl.keywords {
				if strings.HasPrefix(token, k) {
					pre.AppendChild(dom.CreateElement("span").AddClass(color).SetInnerText(token + " "))
					found = true
					break
				}
			}
			if !found {
				pre.AppendChild(dom.CreateElement("span").SetInnerText(token + " "))
			}
		}
		t.absComponent.elem.AppendChild(pre)
	}
	return t
}

func (t *Code) Style(style ...Style) *Code {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Code) Self(ref **Code) *Code {
	*ref = t
	return t
}
