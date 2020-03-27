package wtk

type ModelChanged func(value interface{})

type Handle struct {
	id int
	cb ModelChanged
}

func (f Handle) Valid() bool {
	return f.id > 0
}

type Model struct {
	Value     interface{}
	callbacks []Handle
	ctr       int
}

func (p *Model) Notify() {
	for _, f := range p.callbacks {
		if f.id > 0 {
			f.cb(p.Value)
		}
	}
}

func (p *Model) AddListener(cb ModelChanged) Handle {
	p.ctr++
	f := Handle{id: p.ctr, cb: cb}
	p.callbacks = append(p.callbacks, f)
	return f
}

func (p *Model) RemoveListener(f Handle) {
	for i, x := range p.callbacks {
		if x.id == f.id {
			p.deleteListener(i)
			return
		}
	}
}

func (p *Model) deleteListener(i int) {
	a := p.callbacks
	if i < len(a)-1 {
		copy(a[i:], a[i+1:])
	}
	a[len(a)-1] = Handle{}
	a = a[:len(a)-1]
}
