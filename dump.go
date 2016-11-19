package dump

import (
  "fmt"
)

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
}

//}

func IsPrint(b byte) bool {
	if b >= ' ' && b <= '~' {
		return true
	}
	return false
}
