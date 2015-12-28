// Implementation in Go of busybeaver
package main

import (
	"errors"
	"fmt"
)

// Constants in CAPITALS
const LENGHT = 25

// Structs
// -State is equivalent a 'ncard' of computerphile
type state struct {
	c_wr   byte // new character to write in the tape
	shift  byte // left or right
	nstate int  // next state
}

// -Tape is 'finit' tape of byte (char) whith the head position
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
func newt(tp *tape, lenght int) {
	tp.t = make([]byte, lenght)
	tp.h = lenght / 2
}

// Print tape
func printt(tp tape) {
	//fmt.Println("Tape lenght: ", len(tp.t))
	if len(tp.t) != 0 {
		//fmt.Println("Head position: ", tp.h)
		fmt.Print("[")
		for i := 0; i < len(tp.t); i++ {
			if tp.t[i] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", tp.t[i])
			}
		}
		fmt.Println("]")
		fmt.Print("[")
		for i := 0; i < tp.h; i++ {
			fmt.Print(".")
		}
		fmt.Print("^")
		for i := tp.h + 1; i < len(tp.t); i++ {
			fmt.Print(".")
		}
		fmt.Println("]")
	}
}

// Read character
func readt(tp tape) byte {
	return tp.t[tp.h]
}

// Write character
func writet(tp *tape, c byte) {
	tp.t[tp.h] = c
}

// Shift the head of the tape
func shiftt(tp *tape, shift byte) error {
	switch shift {
	case 'L', 'l':
		{
			if tp.h > 0 {
				tp.h--
			} else {
				return errors.New("ERROR: No space on the left side of the tape")
			}
		}
	case 'R', 'r':
		{
			if tp.h < LENGHT {
				tp.h++
			} else {
				return errors.New("ERROR: No space on the right side of the tape")
			}
		}
	default:
		return errors.New("ERROR: Don't move the head position")
	}
	return nil
}

// This helper will streamline our error checks below.
func check(e error) {
	if e != nil {
		//panic(e)
		fmt.Println(e)
	}
}

func main() {
	var t tape
	//var bb [][2]string

	newt(&t, LENGHT)
	printt(t)

	/* Code for check implementation function for use the tape:
	check(shiftt(&t, 'l'))
	writet(&t, '1')
	printt(t)
	check(shiftt(&t, 'r'))
	printt(t)
	check(shiftt(&t, 'R'))
	printt(t)
	check(shiftt(&t, 'L'))
	printt(t)
	check(shiftt(&t, 'x'))
	printt(t)
	fmt.Printf("Read character: %c.\n", readt(t))
	check(shiftt(&t, 'l'))
	fmt.Printf("Read character: %c.\n", readt(t))
	*/
	/* Code for implementation busybeavers and their execution
	bb = newbb(0)
	printbb(bb)
	bb = newbb(-1)
	printbb(bb)
	*/
}
