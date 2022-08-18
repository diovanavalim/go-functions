package main

import (
	"ddvalim/fc/function"
	"fmt"
)

func main() {
	var n1 int8 = 30
	var n2 int8 = 20

	s1 := function.Sum(n1, n2)
	s2 := function.Sub(n1, n2)

	fmt.Println(s1, s2)

	var n3 int8 = 10
	var n4 int8 = 15

	s3, s4 := function.MultipleOp(n3, n4)
	fmt.Println(s3, s4)

	s5, _ := function.MultipleOp(n3, n4)
	fmt.Println(s5)

	var f = func(txt string, name string) string {
		return txt + ", says " + name
	}

	fmt.Println(f("Hello, World!", "Diovana"))
}
