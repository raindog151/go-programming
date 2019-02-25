package main

import "testing"

func TestFoo(t *testing.T) {

	var v int

	v = Foo(6, 10)

	if v != 16 {
		t.Error("Expected 16, got", v)
	}
}

