package main

import (
	"fmt"
)

const _ = iota
//iota is a constant
const a = iota

const (
	_ =  iota
	b = iota
	c
)

func main()  {
	//typed Constants
	const x int= 12
	//untyped constants
	const y = 12.0
	
	fmt.Printf("%v ,%T\n",x,x)

	fmt.Printf("%v ,%T\n",y,y)

	fmt.Println(c)
}