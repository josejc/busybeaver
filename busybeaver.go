// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

// Constants
const lenght = 25

// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var tape [lenght]byte

	for i := 0; i < lenght; i++ {
		fmt.Println(i, tape[i])
	}
}
