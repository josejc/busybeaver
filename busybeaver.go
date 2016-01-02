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
func newbb(s string) [][]state {
	var bstr [][2]string
	var bb [][]state

	// bstr for reuse the definitions of <computerphile>
	switch s {
	case "tm1a": // loops off to the right
		bstr = [][2]string{{"0000", "1000"}, {"0111", "1111"}}
	case "tm1b": // loops off to the left
		bstr = [][2]string{{"0000", "1000"}, {"0101", "1101"}}
	case "tm1c": //prints a single 1 and moves to the left
		bstr = [][2]string{{"0000", "1000"}, {"0100", "1000"}}
	case "tm1d": //prints a single 1 and moves to the right
		bstr = [][2]string{{"0000", "1000"}, {"0110", "1000"}}
	case "tm2": // simple example of 2-state with optimal score of 4.
		bstr = [][2]string{{"0000", "1000"}, {"0112", "1102"}, {"0101", "1110"}}
	case "tm3a": //classic 3-state BB producing a score of 6
		bstr = [][2]string{{"0000", "1000"}, {"0102", "1113"}, {"0111", "1102"}, {"0112", "1100"}}
	case "tm3aa": //another 3-state BB for sigma ?
		bstr = [][2]string{{"0000", "1000"}, {"0112", "1110"}, {"0013", "1112"}, {"0103", "1101"}}
	case "tm3b": //permutation of tm3a which never terminates
		bstr = [][2]string{{"0000", "1000"}, {"0112", "1100"}, {"0111", "1102"}, {"0102", "1113"}}
	case "tm3c": //another permutation of tm3a which terminates but has a score of only 2
		bstr = [][2]string{{"0000", "1000"}, {"0102", "1113"}, {"0112", "1100"}, {"0111", "1102"}}
	case "tm4": //classic 4-state tm producing a score of 23 in 107 moves
		bstr = [][2]string{{"0000", "1000"}, {"0112", "1102"}, {"0101", "1003"}, {"0110", "1104"}, {"0114", "1011"}}
	}

	bb = make([][]state, len(bstr))
	//write values to array
	for row, a := range bstr {
		bb[row] = make([]state, len(a))
		for col, _ := range a {
			// fmt.Println("row:", row, "col:", col)
			bb[row][col] = bstr2bb(bstr[row][col])
			// fmt.Println("bstr:", bstr[row][col])
			// fmt.Println("bb:", bb[row][col])
		}
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
func printbb(bb [][]state) {
	fmt.Println("bb:", bb)

}

func runbb(bb [][]state, t *tape) error {
	var s, ns int // Index of state in TuringMachine (busybeaver), s=actual state, ns=next state
	var r int     // byte read in tape, casting byte to int ;)
	var e error

	fmt.Println("BusyBeaver:")
	printbb(bb)
	fmt.Println("Initial tape is shown below: Note that head position at each step is marked by ^")
	printt(*t)
	fmt.Println("Game start:")

	ns = 1
	e = nil
	for s = 1; s != 0; s = ns {
		r = int(readt(*t)) - 48 // 48 is ascii '0' code
		// fmt.Println("read:", r, "state:", s, "bb[s][r]:", bb[s][r])
		writet(t, bb[s][r].c_wr)
		e = check(shiftt(t, bb[s][r].shift))
		printt(*t)
		if e == nil {
			ns = int(bb[s][r].nstate) - 48
		} else {
			ns = 0
		}
	}
	return e
}
