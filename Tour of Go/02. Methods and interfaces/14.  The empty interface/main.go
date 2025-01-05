package main

import "fmt"

func main() {
	var i interface{} // An interface type that specifies zero methods is known as the empty interface:
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
