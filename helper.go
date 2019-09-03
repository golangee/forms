package gowtk

import "reflect"

// Determines the identity of a given reference (or pointer), independently from which variable (and therefore
// address) the interface value originates. It panics if v denotes a value type.
//
// If v's Kind is Func, the returned pointer is an underlying code pointer,
// but not necessarily enough to identify a single function uniquely. However, as long as it is not a method
// expression created by reflection, the returned pointer is robust and unique. For method expressions created with
// reflection, makeMethodValue always returns the same code pointer (assuming it is a wrapper, this behavior is
// still consistent). It is questionable what happens, if the Go runtime starts using a moving GC...
func Identity(v interface{}) uintptr {
	return reflect.ValueOf(v).Pointer()
}
