// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

// Structs
// -State is equivalent a 'ncard' of computerphile
type state struct {
	c_wr   byte // new character to write in the tape
	shift  byte // left or right
	nstate int  // next state
}

// Create new busybeaver (turing machine) whith selection of one predefined
func newbb(s int) [][2]string {
	var bb [][2]string
	switch s {
	case 0:
		// tm1a - loops off to the right
		bb = [][2]string{{"000", "000"}, {"111", "111"}}
	default:
		bb = make([][2]string, 0)
	}
	return bb
}

// Print the definition os busybeaver, function only for evaluation
func printbb(bb [][2]string) {
	fmt.Println("lenght: ", len(bb))
	fmt.Println("capacity: ", cap(bb))
	fmt.Println("bb: :", bb)
	fmt.Println("---")

}

func runbb(bb [][2]string, tp tape) {
	printt(tp)
}
