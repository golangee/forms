package wtk

type Listener struct {
	wrapped   interface{}
	onRelease func()
}

func NewListener(f interface{}, onRelease func()) Resource {
	return Listener{
		wrapped:   f,
		onRelease: onRelease,
	}
}

func (l Listener) Release() {
	l.onRelease()
}
