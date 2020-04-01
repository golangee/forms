package wtk

import "github.com/worldiety/wtk/dom"

// A Resource release method should be called for any resource, as soon as it is not required anymore to avoid
// memory leaks. Afterwards the Resource must not be used anymore.
// Even though we have a GC, we cannot rely on it, because the Resource may have registrations
// beyond our process, which requires holding global callback references, so that the outer system
// can call us. An example for this are go functions wrapped as callbacks in the wasm tier made available
// for the javascript DOM, like event handlers. Also cgo or rpc mechanism are possible.
type Resource interface {
	Release() // Release clean up references and the resource must not be used afterwards anymore.
}

// A View is a absComponent on screen. It may not be visible, if it has not been attached yet.
// It is usually created through a constructor method (New<View>) and is a pointer to
// a struct, which itself may contain more components. It is not safe to be used concurrently
// and must only be modified by the UI thread.
type View interface {
	internalView
	Resource
}

// internalView contains implementation specific details, which we don't want developers to depend on.
// The downside is, that a developer cannot implement low level views which interact with the DOM API.
// However, we don't guarantee the kind of DOM API, hell even the wasm API has changed in Go 1.12, 1.13 and 1.14
// every time in an incompatible way which drives all popular type safe bindings entirely useless.
// This API must be considered very unstable and it may also map directly to native ui widgets one day, e.g. for
// iOS or Android widgets.
type internalView interface {
	attach(parent View) // attach adds all nodes and modifications to the given parent node
	detach()            // detach is the reverse of attach
	parent() View       // parent returns nil or if attached the parent
	node() dom.Element  // node returns the underlying DOM element
}
