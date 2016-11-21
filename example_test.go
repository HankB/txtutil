package txtutil_test

import (
  "fmt"
  "github.com/HankB/util"
)


func ExampleTxtutil() {
  fmt.Println("hello world\nHello World")
  // Output:
  // hello world
  // Hello World
}

func ExampleTxtutildump() {
  dump.Dump("hello world")
  fmt.Println("\nEOL check")
  // Output:
  // 00000000  68 65 6c 6c 6f 20 77 6f  72 6c 64
  // EOL check
}
