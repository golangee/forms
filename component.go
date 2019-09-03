package gowtk

// The Component provides a default satisfaction of the View interface to avoid writing boilerplate code.
var _ View = (*Component)(nil)
var _ Touchable = (*Component)(nil)

// The Component should be used to compose the View functionality into custom views
type Component struct {
	children  []View
	listeners map[uintptr]func()
}

func (v *Component) SetPadding(l, t, r, b int) *Component {
	return v
}

// AddEventListener
func (v *Component) AddEventListener(eventType EventType, f func()) {
	//TODO event type is wrong
	if v.listeners == nil {
		v.listeners = make(map[uintptr]func())
	}
	v.listeners[Identity(f)] = f
}

func (v *Component) RemoveEventListener(f func()) {
	if v.listeners != nil {
		delete(v.listeners, Identity(f))
	}
}

// Children returns the backing slice with all appended views so far. Do not modify.
func (v *Component) Children() []View {
	return v.children
}

// Attach calls View#Attach() on all children
func (v *Component) Attach() {
	for _, v := range v.children {
		v.Attach()
	}
}

// Detach calls View#Detach on all children
func (v *Component) Detach() {
	for _, v := range v.children {
		v.Detach()
	}
}

// Add appends the given views and calls View#Attach() on them.
func (v *Component) AddView(views ...View) *Component {
	v.children = append(v.children, views...)
	for _, v := range views {
		v.Attach()
	}
	return v
}

// RemoveView removes the given views and calls View#Detach() on each.
func (v *Component) RemoveView(views ...View) *Component {
	res := make([]View, 0)
	for _, child := range v.children {
		isRemoved := false
		for _, removedChild := range views {
			if removedChild == child {
				isRemoved = true
				break
			}
		}
		if !isRemoved {
			res = append(res, child)
		}
	}
	return v
}
