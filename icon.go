package wtk

import "github.com/worldiety/wtk/theme/material/icon"

type Icon struct {
	Value icon.Icon
	*absComponent
}

func NewIcon(str icon.Icon) *Icon {
	t := &Icon{}
	t.absComponent = newComponent(t, "i")
	t.node().AddClass("material-icons")
	t.Set(str)
	return t
}

func (t *Icon) Set(str icon.Icon) *Icon {
	t.Value = str
	t.absComponent.elem.SetInnerText(string(str))
	return t
}

func (t *Icon) Style(style ...Style) *Icon {
	t.absComponent.style(style...)
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Icon) Self(ref **Icon) *Icon {
	*ref = t
	return t
}

