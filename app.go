package wtk

import "github.com/worldiety/wtk/dom"



type AbsApplication interface {
	SetView(view View)
}

type Application struct {
	rootView Window
	ctx      *myContext
	this AbsApplication
}

func NewApplication(this AbsApplication) *Application {
	a := &Application{}
	a.this = this
	a.ctx = &myContext{r: NewRouter()}
	a.rootView = Window{window: dom.GetWindow(), ctx: a.ctx}
	return a
}

func (a *Application) Window() Window {
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
	a.rootView.RemoveAll()
	a.rootView.AddView(view)
}

func (a *Application) UnmatchedRoute(f func(Query) View) *Application {
	a.Context().router().SetUnhandledRouteAction(func(query Query) {
		a.this.SetView(f(query))
	})
	return a
}

func (a *Application) Start() {
	a.Context().router().Start()
	select {}
}
