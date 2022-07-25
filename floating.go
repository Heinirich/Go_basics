package main

import (
	"fmt"
)

func main() {
	var m float64 = 3.14
	n := 3.14
	n = 13.7e72
	n = 2.1E14

	fmt.Printf("%v %T\n",m,m)
	fmt.Printf("%v %T\n",n,n)
}