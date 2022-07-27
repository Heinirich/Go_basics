package main

import (
	"fmt"
)

const a = iota

func main()  {
	const x int= 12
	
	fmt.Printf("%v ,%T\n",x,x)
}