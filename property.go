package gowtk

import "strconv"

type Property interface {
	Value() interface{}
	SetValue(v interface{})
}

type StringProperty struct {
	val string
}

func (p *StringProperty) String() string {
	return p.val
}

func (p *StringProperty) SetString(v string) StringProperty {
	p.val = v
	return *p
}

func (p *StringProperty) OnChanged(func(old, new string)) *StringProperty {
	return p
}

type IntProperty struct {
	val int
}

func (p *IntProperty) Int() int {
	return p.val
}

func (p *IntProperty) SetInt(v int) *IntProperty {
	p.val = v
	return p
}

func (p *IntProperty) OnChanged(func(old, new string)) *IntProperty {
	return p
}

func (p *IntProperty) Inc() *IntProperty {
	return p
}

func (p *IntProperty) String() string {
	return strconv.Itoa(p.val)
}


type BoolProperty interface {
	Bool() bool
	SetBool(i bool)
	Property
}

func Bool(b bool) BoolProperty {
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
