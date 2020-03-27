package wtk

type ListenerCallback func(old interface{}, new interface{})

type Func struct {
	id int
	cb ListenerCallback
}

func (f Func) Valid() bool {
	return f.id > 0
}




type Property struct {
	value     interface{}
	callbacks []Func
	ctr       int
}


func (p *Property) Get() interface{} {
	return p.value
}

func (p *Property) Set(value interface{}) {
	old := p.value
	p.value = value
	for _, f := range p.callbacks {
		if f.id > 0 {
			f.cb(old, value)
		}
	}
}

func (p *Property) AddListener(cb ListenerCallback) Func {
	p.ctr++
	f := Func{id: p.ctr, cb: cb}
	p.callbacks = append(p.callbacks, f)
	return f
}

func (p *Property) RemoveListener(f Func) {
	for i, x := range p.callbacks {
		if x.id == f.id {
			p.deleteListener(i)
			return
		}
	}
}

func (p *Property) deleteListener(i int) {
	a := p.callbacks
	if i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	a[len(a)-1] = Func{}
	a = a[:len(a)-1]
}




type String struct {
	*Property
}

func (p *String) Get() string {
	if p.Property == nil {
		return ""
	}
	return p.Property.Get().(string)
}

func (p *String) Set(s string) {
	if p.Property == nil {
		p.Property = &Property{}
	}
	p.Property.Set(s)
}
