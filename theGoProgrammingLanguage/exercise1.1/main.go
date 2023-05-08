// Modify the echo program to also print os.Args[0], the name of the command that invoked it.

package main

import (
	"fmt"
	"os"
)

func main() {
	var s string
	for i := 0; i < len(os.Args); i++ {
		s += os.Args[i]
		s += ","
	}
	fmt.Println(s)
}
