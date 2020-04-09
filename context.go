package wtk

import (
	"net/url"
	"strconv"
	"strings"
)

type NamedParameter struct {
	Key   string
	Value string
}

type Context interface {
	router() *Router
	Navigate(path string, params ...NamedParameter)
	Routes() []Route
}

func IntParam(key string, val int) NamedParameter {
	return NamedParameter{key, strconv.Itoa(val)}
}

func StrParam(key string, val string) NamedParameter {
	return NamedParameter{key, val}
}

func NewContext(router *Router) Context {
	return &myContext{r: router}
}

type myContext struct {
	r *Router
}

func (c *myContext) Navigate(path string, params ...NamedParameter) {
	sb := &strings.Builder{}
	sb.WriteString("#")
	sb.WriteString(path)
	if len(params) > 0 {
		sb.WriteString("?")
		for i, p := range params {
			sb.WriteString(p.Key)
			sb.WriteString("=")
			sb.WriteString(url.QueryEscape(p.Value))
			if i < len(params)-1 {
				sb.WriteString("&")
			}

		}
	}
	u, err := url.Parse(sb.String())
	if err != nil {
		panic(err)
	}
	c.router().Navigate(u)
}

func (c *myContext) Routes() []Route {
	return c.r.Routes()
}

func (c *myContext) router() *Router {
	return c.r
}
