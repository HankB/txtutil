package txtutil

import (
	"fmt"
)

// analog of libc strlen on a []byte
func strlen(s []byte) int {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == 0 {
			break
		}
	}
	return i
}

const count = 16 // Number of bytes to print each line

// endLine has two jobs. First pad remainder of line if <16
// bytes have been printed and then pad and add the ascii
// representation (and '\n') This is part of the code that will print
// binary strings using a format like 'hexdump -C'
/*
hbarta@swanky:~/Documents/go-wks/src/github.com/HankB/playground$ hexdump -C ~/go/bin/go|tail -3
0097e1f0  5f 73 74 61 72 74 00 63  72 6f 73 73 63 61 6c 6c  |_start.crosscall|
0097e200  5f 61 6d 64 36 34 00                              |_amd64.|
0097e207
hbarta@swanky:~/Documents/go-wks/src/github.com/HankB/playground$
*/
func endLine(str []byte) {
  add := count - strlen(str)
  if add < 0 {  // test for pathalogical case
    add = 0
  }
  // add spaces so ASCII rep lines up with pervious line
  for i:=0; i<add; i++ {
    fmt.Printf("   ")
  }
  if add <= count/2 {
    fmt.Printf(" ")
  }
	fmt.Printf("|%s|\n", str)
}

// Format output similarly to 'hexdump -D'
func Dump(b string) {
	chars := 0 // chars printed
	asciiRep := []byte{}
	for i := 0; i < len(b); i++ { // iterate over entire array
		if (chars % 16) == 0 { // starting a full row?
			if chars > 0 { // need to print end of previous line?
				for i := 0; (i+chars)%16 != 0; i++ { // fill out line < 16 bytes
					fmt.Printf("   ")
					asciiRep = append(asciiRep, ' ')
					if chars < 8 {
						asciiRep = append(asciiRep, ' ')
					}
					fmt.Printf(" |%s|\n", asciiRep)
					asciiRep = []byte{}
				}
			}
			fmt.Printf("%8.8x  ", chars)
		} else if (chars % 8) == 0 { // mid row?
			fmt.Printf(" ") // mid row space
		}

		// add the hex representation to the output
		fmt.Printf("%2.2x ", b[i])
		// add the ASCII rep (if printable)
		if IsPrint(b[i]) {
			asciiRep = append(asciiRep, b[i])
		} else {
			asciiRep = append(asciiRep, '.')
		}

		chars++
	}
  endLine(asciiRep)
}

//}

func IsPrint(b byte) bool {
	if b >= ' ' && b <= '~' {
		return true
	}
	return false
}
