package main

import "fmt"

type I interface {
	M()
}

// Since there is no type inside the interface tuple to indicate which concrete method to call, it will result in a run time error
func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
