package txtutil_test

import (
	"testing"

	"github.com/HankB/txtutil"
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

// TestIsPrint tests IsPrint()
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

func ExampleTxtutildump() {

	txtutil.Dump("hello world")
	txtutil.Dump(DISCOVERY)
	txtutil.Dump(ON)
	txtutil.Dump(MAGIC)
	txtutil.Dump("hello world agai")
	txtutil.Dump("hello world")
	txtutil.Dump("hello worl")
	txtutil.Dump("hello wor")
	txtutil.Dump("hello wo")
	txtutil.Dump("hello w")
	txtutil.Dump("hello ")
	txtutil.Dump("hello world again")
	txtutil.Dump("hello world again, how long should we go on")
	// Output:
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
	// 0000000b
	// 00000000  00 06 71 61                                       |..qa|
	// 00000004
	// 00000000  01                                                |.|
	// 00000001
	// 00000000  68 64                                             |hd|
	// 00000002
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 20 61 67 61 69  |hello world agai|
	// 00000010
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64                 |hello world|
	// 0000000b
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c                    |hello worl|
	// 0000000a
	// 00000000  68 65 6c 6c 6f 20 77 6f  72                       |hello wor|
	// 00000009
	// 00000000  68 65 6c 6c 6f 20 77 6f                           |hello wo|
	// 00000008
	// 00000000  68 65 6c 6c 6f 20 77                              |hello w|
	// 00000007
	// 00000000  68 65 6c 6c 6f 20                                 |hello |
	// 00000006
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 20 61 67 61 69  |hello world agai|
	// 00000010  6e                                                |n|
	// 00000011
	// 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64 20 61 67 61 69  |hello world agai|
	// 00000010  6e 2c 20 68 6f 77 20 6c  6f 6e 67 20 73 68 6f 75  |n, how long shou|
	// 00000020  6c 64 20 77 65 20 67 6f  20 6f 6e                 |ld we go on|
	// 0000002b

}
