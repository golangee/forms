package property

type String interface {
	Set(v string)
	Get() string
	Bind(dst *string)
	Observe(onDidSet func(old, new string)) Func
}

func NewString() String {
	return &stringProperty{absProperty: newAbsProperty()}
}

type stringProperty struct {
	*absProperty
}

func (s *stringProperty) Set(v string) {
	s.absProperty.Set(v)
}

func (s *stringProperty) Get() string {
	if s.absProperty.Get() == nil {
		return ""
	}

	return s.absProperty.Get().(string)
}

// TODO when and where to unbind?
func (s *stringProperty) Bind(dst *string) {
	s.absProperty.Observe(func(old, new interface{}) {
		*dst = new.(string)
	})
	// TODO unclear if this is a good idea, it is the only time, it will make this way
	s.Set(*dst)
}

func (s *stringProperty) Observe(onDidSet func(old, new string)) Func {
	return s.absProperty.Observe(func(old, new interface{}) {
		if old == nil {
			old = ""
		}

		if old != new {
			onDidSet(old.(string), new.(string))
		}
	})
}
