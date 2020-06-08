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
	"strconv"
	"strings"
	"syscall/js"
)

var htmlId = 1

// assertNotAttached bails out if parent is not nil
func assertNotAttached(v View) {
	if v.parent() != nil {
		panic("invalid state: view is already attached")
	}
}

// assertAttached bails out if parent is nil
func assertAttached(v View) {
	if v.parent() == nil {
		panic("invalid state: view is not attached")
	}
}

func floatToPx(v float64) string {
	return strconv.Itoa(int(v)) + "px"
}

func intToPx(v int) string {
	return strconv.Itoa(v) + "px"
}

func nextId() string {
	htmlId++
	return "id-" + strconv.Itoa(htmlId)
}


func debugStr(value js.Value) string {
	sb := &strings.Builder{}
	sb.WriteString(value.Type().String())
	sb.WriteString(":")
	keys := js.Global().Get("Object").Call("keys", value)
	for i := 0; i < keys.Length(); i++ {
		sb.WriteString(keys.Index(i).String())
		sb.WriteString(",")
	}
	return sb.String()
}
