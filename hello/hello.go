package main

import (
	"fmt"

	"rsc.io/quote"

	"strconv"
)

var i int = 43

func main() {

	j := 35.56

	i = 87

	s1 := "32"
	s2 := "65"

	si1, _ := strconv.Atoi(s1)

	si2, _ := strconv.Atoi(s2)

	fmt.Println("Hello World")
	fmt.Println(quote.Go())
	fmt.Printf("%d and %T \n", i, i)
	fmt.Printf("%v and %T \n", j, j)
	fmt.Printf("%v and %T \n", s1, s1)
	fmt.Println(si1)
	fmt.Println(si2)

	var a int = 61
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(&a, *b)
	fmt.Printf("%v %v %p\n", a, b, b)
}
