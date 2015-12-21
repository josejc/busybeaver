// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

// Constants
const lenght = 25

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

// Print tape
func printt(t []byte) {
	fmt.Println("Tape lenght: ", len(t))
	fmt.Println(t)
	//	for i := 0; i < len(t); i++ {
	//		fmt.Print(i, ",", t[i], "-")
	//	}
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
	var bb [][2]string

	tape = make([]byte, lenght)
	printt(tape)
	bb = newbb(0)
	printbb(bb)
	bb = newbb(-1)
	printbb(bb)
}
