package property

type DidSet func(old, new interface{})

type Func struct {
	id int
	f  DidSet
}

type absProperty struct {
	value     interface{}
	callbacks map[int]Func
	ctr       int
}

func newAbsProperty() *absProperty {
	return &absProperty{
		callbacks: map[int]Func{},
	}
}

func (a *absProperty) Set(v interface{}) {
	old := a.value
	a.value = v
	a.Notify(old, v)
}

func (a *absProperty) Get() interface{} {
	return a.value
}

func (a *absProperty) Observe(f DidSet) Func {
	a.ctr++
	fun := Func{
		id: a.ctr,
		f:  f,
	}
	a.callbacks[a.ctr] = fun

	return fun
}

func (a *absProperty) Notify(old, new interface{}) {
	for _, f := range a.callbacks {
		f.f(old, new)
	}
}
