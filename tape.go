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
	for i := range tp.t {
		tp.t[i] = '0'
	}
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

// Print tape with the step number
func printtn(tp tape, s int) {
	//fmt.Println("Tape lenght: ", len(tp.t))
	if len(tp.t) != 0 {
		//fmt.Println("Head position: ", tp.h)
		fmt.Print(s, ":[")
		for i := 0; i < len(tp.t); i++ {
			if tp.t[i] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Printf("%c", tp.t[i])
			}
		}
		fmt.Println("]")
		fmt.Print(s, ":[")
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

	l := len(tp.t) - 1
	switch shift {
	case 'L', 'l', '0': // 0 definition in <computerphile>
		{
			if tp.h > 0 {
				tp.h--
			} else {
				return errors.New("ERROR: No space on the left side of the tape")
			}
		}
	case 'R', 'r', '1': // 1 definition in <computerphile>
		{
			if tp.h < l {
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

func number1s(tp tape) int {

	n := 0
	for i := 0; i < len(tp.t); i++ {
		if tp.t[i] == 49 {
			n++
		}
	}
	return n
}

// This helper will streamline our error checks below.
func check(e error) error {
	if e != nil {
		//panic(e)
		fmt.Println(e)
	}
	return e
}
