// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

// Constants
const lenght_tape = 25

// Structs
// -State is equivalent a 'ncard' of computerphile
type state struct {
	c_wr   byte // new character to write in the tape
	shift  byte // left or right
	nstate int  // next state
}

type tape struct {
	t []byte
	h int // head position
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

// New tape
func newt(tp tape, lenght int) {
	tp.t = make([]byte, lenght)
	tp.h = lenght / 2
	printt(tp)
}

// Print tape
func printt(tp tape) {
	fmt.Println("Tape lenght: ", len(tp.t))
	if len(tp.t) != 0 {
		fmt.Println("Head position: ", tp.h)
		fmt.Println(tp.t)
		fmt.Print("[")
		for i := 0; i < tp.h; i++ {
			fmt.Print(". ")
		}
		fmt.Print("^ ")
		for i := tp.h + 1; i < len(tp.t)-1; i++ {
			fmt.Print(". ")
		}
		fmt.Println(".]")
	}
}

// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var t tape
	var bb [][2]string

	newt(t, lenght_tape)
	printt(t)
	bb = newbb(0)
	printbb(bb)
	bb = newbb(-1)
	printbb(bb)
}
