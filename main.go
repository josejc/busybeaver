// Implementation in Go of busybeaver
package main

import (
	"fmt"
)

func main() {
	var bb [][2]state
	var t tape

	newt(&t, LENGHT)
	printt(t)

	// Code for check implementation function for use the tape:
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

	// Code for implementation busybeavers and their execution
	bb = newbb("tm1a")
	printbb(bb)
	bb = newbb("x")
	printbb(bb)
	//bb = newbb(1)
	//printbb(bb)
}
