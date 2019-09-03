package gowtk

type EventType int

const (
	Click EventType = iota
	DoubleClick
)

// A Touchable provides the possibility to add and remove event listeners
type Touchable interface {
	AddEventListener(eventType EventType, f func())
	RemoveEventListener(f func())
}
