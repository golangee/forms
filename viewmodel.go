// Copyright 2020 Torben Schinke
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package forms

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
