package gowtk

type View interface {
}


type Component interface {
	SetPadding(l, t, r, b int) Component
	OnClick(func())Component
}

type Context interface{}

type ViewSpec func(ctx View)

type TextView interface {
	Content() StringProperty
	SetContent(p StringProperty) Component

Component

}

func Text(ctx Context) TextView {
	return nil
}
