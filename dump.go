// Package txtutil provides the ability to dump binary (or ASCII) data to
// the console.
package txtutil

import (
	"fmt"
)

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
	add := count - len(str)
	if add < 0 { // test for pathalogical case
		add = 0
	}
	// add spaces so ASCII rep lines up with other lines
	for i := 0; i < add; i++ {
		fmt.Printf("   ")
	}
	// add an extra 'midline' space when needed
	if add >= count/2 {
		fmt.Printf(" ")
	}
	fmt.Printf(" |%s|\n", str)
}

// Dump binary data using format similar to 'hexdump -C'
func Dump(b string) {
	chars := 0 // chars printed
	asciiRep := []byte{}
	for i := 0; i < len(b); i++ { // iterate over entire array
		if (chars % 16) == 0 { // starting a full row?
			if chars > 0 { // need to print end of previous line?
				endLine(asciiRep)   //print end of stuff
				asciiRep = []byte{} // clear end of line buffer
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
	if len(asciiRep) > 0 { // starting a full row?
		endLine(asciiRep)
	}
	fmt.Printf("%8.8x\n", chars) // print final character count
}

// IsPrint() is exported only for testing. (Can't get the export_test.g thing to work)
func IsPrint(b byte) bool {
	if b >= ' ' && b <= '~' {
		return true
	}
	return false
}
