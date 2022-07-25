package main

import (
	
	"fmt"
)

func main()  {
	s := "Hello Mate"
	fmt.Println(s)
	//Conversion to bytes
	b := []byte(s)
	fmt.Printf("%v, %T\n",b,b)
}