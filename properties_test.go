package wtk

import (
	"fmt"
	"testing"
)

func TestString_Get(t *testing.T) {
	str := &String{}
	str.Set("hello")
	if str.Get() != "hello"{
		fmt.Println(str.value)
	}
}