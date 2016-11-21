package txtutil_test

import (
	"fmt"
	"github.com/HankB/txtutil"
	"testing"
)

func TestIsPrint(t *testing.T) {
	if !txtutil.IsPrint(' ') {
		t.Fail()
	}
	if !txtutil.IsPrint('~') {
		t.Fail()
	}
	if !txtutil.IsPrint('\x10') {
		t.Fail()
	}
	if txtutil.IsPrint('\xF0') {
		t.Fail()
	}
}

func ExampleTxtutil() {
	fmt.Println("hello world")
	fmt.Println("Hello World")
	// Output:
	// hello world
	// Hello World
}

func ExampleTxtutildump() {
	// don't understand why this test fails
	txtutil.Dump("hello world")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64
}
