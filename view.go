package wtk

import "github.com/worldiety/wtk/dom"

type View interface {
	wasmView
}

type wasmView interface {
	attach(parent dom.Element)
	detach(parent dom.Element)
}

type ViewGroup interface {
	AddView(view View)
}
