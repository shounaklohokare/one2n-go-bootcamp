package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// it will print a mix of world and hello instead due to the concurrent & non blocking execution of the go routine
func main() {
	go say("world")
	say("hello")
}
