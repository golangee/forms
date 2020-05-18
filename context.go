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
