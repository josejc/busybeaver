// Implementation a tape for Turing Machine (busybeaver)
package main

import (
	"errors"
	"fmt"
)

// Constants in CAPITALS
const LENGHT = 25

// -Tape is 'finit' tape of byte (char) whith the head position
type tape struct {
	t []byte
	h int // head position
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
