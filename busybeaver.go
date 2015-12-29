// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

// -State is equivalent a 'ncard' of computerphile
type state struct {
	c_wr   byte // new character to write in the tape
	shift  byte // left or right
	nstate byte // next state
}

// Create new busybeaver (turing machine) whith selection of one predefined
func newbb(s string) [][2]state {
	var bstr [][2]string
	var bb [][2]state

	// bstr for reuse the definitions of <computerphile>
	switch s {
	case "tm1a": // loops off to the right
		bstr = [][2]string{{"0000", "1000"}, {"0111", "1111"}}
	case "tm1b": // loops off to the left
		bstr = [][2]string{{"0000", "1000"}, {"0101", "1101"}}
	case "tm1c": //prints a single 1 and moves to the left
		bstr = [][2]string{{"0000", "1000"}, {"0100", "1000"}}
		/*
			//prints a single 1 and moves to the right
			char * tm1d [2] [2] = {"0000", "1000", "0110", "1000"};
			// simple example of 2-state with optimal score of 4.
			char * tm2 [3] [2] = {"0000", "1000", "0112", "1102", "1101", "1110"};
			//classic 3-state BB producing a score of 6
			char * tm3a [4] [2] = {"0000", "1000", "0102", "1113", "0111", "1102", "0112", "1100"};
			//another 3-state BB for sigma ?
			char * tm3aa [4] [2] = {"0000", "1000", "0112", "1110", "0013", "1112", "0103", "1101"};
			//permutation of tm3a which never terminates
			char * tm3b [4] [2] = {"0000", "1000", "0112", "1100", "0111", "1102", "0102", "1113"};
			//another permutation of tm3a which terminates but has a score of only 2
			char * tm3c [4] [2] = {"0000", "1000", "0102", "1113", "0112", "1100", "0111", "1102"};
			//classic 4-state tm producing a score of 23 in 107 moves
			char * tm4 [5] [2] = {"0000", "1000", "0112", "1102", "0101", "1003", "0110", "1104", "0114", "1011" };
		*/
	}
	if len(bstr) > 0 {
		bb = make([][2]state, len(bstr))
		for i := 0; i < len(bstr); i++ {
			bb[i][0] = bstr2bb(bstr[i][0])
			bb[i][1] = bstr2bb(bstr[i][1])
		}
	} else {
		bb = make([][2]state, 0)
	}
	return bb
}

func bstr2bb(bstr string) state {
	var s state

	// The definition of the states in string of <computerphile> the first digit [0] is the character read in the tape
	s.c_wr = bstr[1]
	s.shift = bstr[2]
	s.nstate = bstr[3]
	return s
}

// Print the definition os busybeaver, function only for evaluation
func printbb(bb [][2]state) {
	fmt.Println("lenght: ", len(bb))
	fmt.Println("capacity: ", cap(bb))
	fmt.Println("bb: :", bb)
	fmt.Println("---")

}

func runbb(bb [][2]state, tp tape) {
	printt(tp)
}
