package wtk

type LinkTarget string

const TargetBlank LinkTarget = "_blank"
const TargetSelf LinkTarget = "_self"
const TargetParent LinkTarget = "_parent"
const TargetTop LinkTarget = "_top"

type Link struct {
	*absComponent
}

func NewLink(caption string, href string) *Link {
	t := &Link{}
	t.absComponent = newComponent(t, "a")
	t.SetCaption(caption)
	t.SetRef(href)
	return t
}

func (t *Link) SetRef(r string) *Link {
	t.elem.SetHref(r)
	return t
}

func (t *Link) SetCaption(str string) *Link {
	t.absComponent.elem.SetText(str)
	return t
}

func (t *Link) Style(style ...Style) *Link {
	t.absComponent.style(style...)
	return t
}

func (t *Link) SetTarget(target LinkTarget) *Link {
	t.node().Unwrap().Set("target", string(target))
	return t
}

// Self assigns the receiver to the given pointer to reference
func (t *Link) Self(ref **Link) *Link {
	*ref = t
	return t
}
