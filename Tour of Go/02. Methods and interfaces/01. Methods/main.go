package main

import (
	"fmt"
)

type Rectangle struct {
	Length, Breadth float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breadth
}

func main() {
	r := Rectangle{3, 4}
	fmt.Println(r.Area())
}
