// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

// Constants
const lenght = 25

// Structs
// -State is equivalent a 'ncard' of compterphile
type state struct {
	c_wr byte	// new character to write in the tape
	shift byte	// left or right
	nstate int	// next state
}


// Create new busybeaver (turing machine) whith selection of one predefined
func newbb(s int) []state {
	switch s {
		case 0:
			return make([]state, 5)			
		default:	
	}
	return make([]state, 4)
}

// Print the definition os busybeaver, function only for evaluation
func printbb(bb []state) {
	fmt.Println("lenght: ", len(bb))
	fmt.Println("capacity: ", cap(bb))	
	fmt.Println("---")
}

func printt(t []byte) {
	fmt.Println("Tape lenght: ", len(t))
	for i := 0; i < len(t); i++ {
		fmt.Print(i,",",t[i], "-")
	}
	fmt.Println("...")
}

// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var tape []byte
	var bb []state

	tape = make([]byte, lenght)
	printt(tape)
	bb = newbb(0)
	printbb(bb)
	bb = newbb(-1)
	printbb(bb)
}
