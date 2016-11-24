package txtutil_test

import (
	 "fmt"
	"github.com/HankB/txtutil"
	"testing"
)

// some strings copied from S20.go for testing
const MAGIC = "\x68\x64"
const DISCOVERY = "\x00\x06\x71\x61"
const DISCOVERY_RESP = "\x00\x2a\x71\x61"
const SUBSCRIBE = "\x00\x1e\x63\x6c"
const SUBSCRIBE_RESP = "\x00\x18\x63\x6c"
const CONTROL = "\x00\x17\x64\x63"
const CONTROL_RESP = "\x00\x17\x73\x66"
const PADDING_1 = "\x20\x20\x20\x20\x20\x20"
const PADDING_2 = "\x00\x00\x00\x00"
const ON = "\x01"
const OFF = "\x00"


func TestIsPrint(t *testing.T) {
	if !txtutil.IsPrint(' ') {
		t.Fail()
	}
	if !txtutil.IsPrint('~') {
		t.Fail()
	}
	if txtutil.IsPrint('\x10') {
		t.Fail()
	}
	if txtutil.IsPrint('\xF0') {
		t.Fail()
	}
}

/*
func ExampleTxtutil() {
	fmt.Println("hello world")
	fmt.Println("Hello World")
	// Output:
	// hello world
	// Hello World
}
*/

func ExampleTxtutildump() {
  fmt.Printf("00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world| << canonical")
	// don't understand why this test fails
	txtutil.Dump("hello world")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
  txtutil.Dump(DISCOVERY)
  // Output:
  // 00000000  00 06 71 61                                       |..qa|
  txtutil.Dump(ON)
  // Output:
  // 00000000  00 06 71 61                                       |..qa|
  txtutil.Dump(MAGIC)
  // Output:
  // 00000000  00 06 71 61                                       |..qa|
  txtutil.Dump("hello world")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
  txtutil.Dump("hello worl")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
  txtutil.Dump("hello wor")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
  txtutil.Dump("hello wo")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
  txtutil.Dump("hello w")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
  txtutil.Dump("hello ")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|

}
