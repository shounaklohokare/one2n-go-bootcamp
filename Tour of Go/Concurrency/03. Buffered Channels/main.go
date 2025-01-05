package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3 // this overfills the buffer and results in a error. if we receive from the channel before then we can free up the buffer slot
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
