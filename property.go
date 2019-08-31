package gowtk

type Property interface {
	Value() interface{}
	SetValue(v interface{})
}

type StringProperty interface {
	String() string
	SetString(str string)
	OnChanged(func(old, new string))
	Bind(other StringProperty) StringProperty
	Property
}

type IntProperty interface {
	Int() int
	SetInt(i int) IntProperty
	Inc() IntProperty
	String()string
	Property
}

type BoolProperty interface {
	Bool() bool
	SetBool(i bool)
	Property
}

func Int(i int) IntProperty {
	return nil
}
func Bool(b bool) BoolProperty {
	return nil
}

func String(str string) StringProperty {
	return nil
}

type implProperty struct {
	value interface{}
}

func (p *implProperty) Value() interface{} {
	return p.value
}

func (p *implProperty) SetValue(v interface{}) {
	p.value = v
}

// TODO this works around polymorphic overloading but casting let us insert invalid types, e.g. ints into string
// properties
type implStringProperty struct {
	*implProperty
}

func (p *implStringProperty) String() string {
	return p.Value().(string)
}

func (p *implStringProperty) SetString(str string) {
	p.SetValue(str)
}
