package wtk

import (
	"github.com/worldiety/wtk/dom"
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
