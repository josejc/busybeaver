// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

func main() {
	var bb [][]state
	var t tape
	var e error

	x := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(x)

	newt(&t, LENGHT)
	//printt(t)

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
	// Code for implementation busybeavers and their execution
	bb = newbb("tm1d")
	//printbb(bb)
	//bb = newbb("x")
	//printbb(bb)
	e = runbb(bb, &t)
	if e != nil {
		fmt.Println("ERROR: Tape too short")
	}
	fmt.Println("Final tape config is:")
	printt(t)
}
