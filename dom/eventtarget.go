package dom

// EventTarget is a DOM interface implemented by objects that can receive events and may have listeners for them.
//
//Element, Document, and Window are the most common event targets, but other objects can be event targets too, for example XMLHttpRequest, AudioNode, AudioContext, and others.
//
//Many event targets (including elements, documents, and windows) also support setting event handlers via onevent properties and attributes.
type EventTarget interface {

	// The EventTarget method addEventListener() sets up a function that will be called whenever the specified event is delivered to the target. Common targets are Element, Document, and Window, but the target may be any object that supports events (such as XMLHttpRequest).
	//
	//addEventListener() works by adding a function or an object that implements EventListener to the list of event listeners for the specified event type on the EventTarget on which it's called.
	AddEventListener(typ string, listener func())
}
