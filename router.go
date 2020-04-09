package wtk

import (
	"github.com/worldiety/wtk/dom"
	"log"
	"net/url"
	"sort"
	"strconv"
	"syscall/js"
)

type Query struct {
	path   string
	values url.Values
}

func (p Query) Path() string {
	return p.path
}

func (p Query) Str(key string) string {
	return p.values.Get(key)
}

func (p Query) Int(key string) int {
	i, _ := strconv.ParseInt(p.Str(key), 10, 64)
	return int(i)
}

type Route struct {
	Path        string
	Constructor func(q Query)
}

type Router struct {
	routes2Actions map[string]func(q Query)
	funcs          []js.Func
	lastLocation   string
	lastFragment   string
	unmatchedRoute func(Query)
}

func NewRouter() *Router {
	r := &Router{
		routes2Actions: make(map[string]func(Query)),
	}

	f := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		r.checkLocation()
		return nil
	})
	r.funcs = append(r.funcs, f)
	dom.GetWindow().Unwrap().Set("onhashchange", f)
	dom.GetWindow().Unwrap().Set("hashchange", f)
	dom.GetWindow().Unwrap().Set("onpopstate", f)

	r.lastLocation = "$%&/"
	r.lastFragment = r.lastLocation
	return r
}

func (r *Router) Routes() []Route {
	var res []Route
	for k, v := range r.routes2Actions {
		res = append(res, Route{
			Path:        k,
			Constructor: v,
		})
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i].Path < res[j].Path
	})
	return res
}

func (r *Router) AddRoute(path string, f func(Query)) *Router {
	log.Println("registered route " + path)
	r.routes2Actions[path] = f
	return r
}

func (r *Router) SetUnhandledRouteAction(f func(Query)) *Router {
	r.unmatchedRoute = f
	return r
}

func (r *Router) Start() {
	r.checkLocation()
}

func (r *Router) Release() {
	for _, f := range r.funcs {
		f.Release()
	}
}

func (r *Router) checkLocation() {
	location := dom.GetWindow().Unwrap().Get("window").Get("location").Get("href").String()
	if r.lastLocation != location {
		u, err := url.Parse(location)
		if err != nil {
			log.Printf("Failed to parse location %s: %v", location, err)
			return
		}
		r.onLocationChanged(u)
		r.lastLocation = location

		if u.Fragment != r.lastFragment {
			f, err := url.Parse(u.Fragment)
			if err != nil {
				log.Printf("Failed to parse fragment as url %s: %v", u, err)
				return
			}
			r.onFragmentChanged(f.Path, f.Query())
			r.lastFragment = u.Fragment
		}

	}
}

func (r *Router) onLocationChanged(location *url.URL) {

}

func (r *Router) onFragmentChanged(path string, query url.Values) {
	if path == "" {
		path = "/"
	}
	q := Query{values: query, path: path}
	f := r.routes2Actions[path]
	if f != nil {
		f(q)
	} else {
		if r.unmatchedRoute != nil {
			r.unmatchedRoute(q)
		} else {
			log.Printf("unmatched route %s->%v\n", path, query)
		}
	}
}

func (r *Router) Navigate(u *url.URL) {
	dom.GetWindow().Unwrap().Set("location", u.String())
}
