package property

// Int defines the contract for an int property.
type Int interface {
	// Set updates the property value and fires an event, if the new value is different than the old value.
	Set(v int)

	// Get returns the current property value
	Get() int

	// Bind connects the given pointer to a string with the value. The first time, the value from dst is read and
	// populates the property. Afterwards the direction is always the opposite, and updates to the property
	// will update the dst.
	Bind(dst *int)

	// Observe registers a callback which is fired, if the value has been set. It is not fire, if the value has not
	// been changed, e.g. if setting the same string.
	Observe(onDidSet func(old, new int)) Func
}

// NewInt creates a new self-contained property.
func NewInt() Int {
	return &intProperty{absProperty: newAbsProperty()}
}

type intProperty struct {
	*absProperty
}

func (s *intProperty) Set(v int) {
	s.absProperty.Set(v)
}

func (s *intProperty) Get() int {
	if s.absProperty.Get() == nil {
		return 0
	}

	return s.absProperty.Get().(int)
}

// TODO when and where to unbind?
func (s *intProperty) Bind(dst *int) {
	s.absProperty.Observe(func(old, new interface{}) {
		*dst = new.(int)
	})
	// TODO unclear if this is a good idea, it is the only time, it will make this way
	s.Set(*dst)
}

func (s *intProperty) Observe(onDidSet func(old, new int)) Func {
	return s.absProperty.Observe(func(old, new interface{}) {
		if old == nil {
			old = 0
		}

		if old != new {
			onDidSet(old.(int), new.(int))
		}
	})
}
