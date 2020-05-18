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

import (
	"github.com/golangee/forms/dom"
	"time"
)

type AbsApplication interface {
	SetView(view View)
}

type Application struct {
	rootView           *Window
	ctx                *myContext
	this               AbsApplication
	versionWatch       *Watch
	hasVersionMismatch bool
}

func NewApplication(this AbsApplication, expectedVersion string) *Application {
	a := &Application{}
	a.this = this
	a.versionWatch = NewWatch(expectedVersion)
	a.versionWatch.AddListener(func(found string, expected string) {
		if a.hasVersionMismatch {
			return
		}
		a.hasVersionMismatch = true
		a.versionWatch.Stop()
		a.showVersionMismatch()
	})
	a.versionWatch.SetInterval(5 * time.Minute)
	a.ctx = &myContext{r: NewRouter()}
	a.rootView = &Window{window: dom.GetWindow(), ctx: a.ctx}
	return a
}

func (a *Application) showVersionMismatch() {
	if !a.hasVersionMismatch {
		return
	}
	NewSnackbar("New App version found", "Reload").SetTimeout(-1).SetAction(func(v View) {
		a.ctx.router().Reload(true)
	}).Show(a.rootView)
}

func (a *Application) Window() *Window {
	return a.rootView
}

func (a *Application) Context() Context {
	return a.ctx
}

func (a *Application) Route(path string, f func(Query) View) *Application {
	a.Context().router().AddRoute(path, func(query Query) {
		a.this.SetView(f(query))
	})
	return a
}

func (a *Application) SetView(view View) {
	if !a.hasVersionMismatch {
		a.versionWatch.Check()
	}

	a.rootView.RemoveAll()
	a.rootView.AddView(view)
	a.showVersionMismatch()
}

func (a *Application) UnmatchedRoute(f func(Query) View) *Application {
	a.Context().router().SetUnhandledRouteAction(func(query Query) {
		a.this.SetView(f(query))
	})
	return a
}

func (a *Application) Start() {
	a.Context().router().Start()
	a.versionWatch.Start()
	select {}
}
