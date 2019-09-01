package gowtk

import "testing"



func TestStrProperty_SetValue(t *testing.T) {
	type blub struct{
		title StringProperty
	}

	bla := &blub{}
	bla.title.SetString("hello")
	if bla.title.String()!="hello"{
		t.Fatal("fail")
	}
}