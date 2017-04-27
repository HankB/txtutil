// Package txtutil provides the ability to dump binary (or ASCII) data to
// the console.
package txtutil

import (
	"fmt"
)

var requiredDumpPri = 4

// PriDump Dumps if priority criteria is met
func PriDump(pri int, b string) {
	if pri >= requiredDumpPri {
		Dump(b)
	}
}

// SetDumpPriority updates the priority and returns the previous value
func SetDumpPriority(pri int) int {
	retVal := requiredDumpPri
	requiredDumpPri = pri
	return retVal
}

var requiredPrintlnPri = 4

// PriFmtPrintln outputs remaining arguments according to the value
// of the first argument.
func PriFmtPrintln(pri int, a ...interface{}) (n int, err error) {
	if pri >= requiredPrintlnPri {
		return fmt.Println(a...)
	}
	return 0, nil
}

// SetPriPrintlnPriority updates the priority and returns the previous value
func SetPriPrintlnPriority(pri int) int {
	retVal := requiredPrintlnPri
	requiredPrintlnPri = pri
	return retVal
}
