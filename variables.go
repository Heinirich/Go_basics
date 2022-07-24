package main

import (
	"fmt"
	"strconv"
)

//at package level we cannot use package colon syntax
var m int =45

var (
	actorName string = "Smith Heinrich"
	companionName string = "Hello"
)

func main()  {
	var i int
	i = 42
	fmt.Println(i)

	var j int = 43
	fmt.Println(j)

	k := 44
	fmt.Println(k)
	//%v is value while %T is type
	fmt.Printf("%v,%T",k,k)
	fmt.Println(m)

	var theHTTP string
	theHTTP = "facebook.com"
	fmt.Println(theHTTP)
	//Type conversion

	var l float32 = 32.5
	var m int
	m = int(l)
	fmt.Printf("%v, %T\n",m,m)

	var o string
	o = strconv.Itoa(i)
	fmt.Printf("%v , %T\n",o,o)
}