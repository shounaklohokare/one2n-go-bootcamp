package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// in order traversal of tree that sends all values to the channel
func walk(t *tree.Tree, ch chan int) {

	if t == nil {
		return
	}
	walk(t.Left, ch)
	ch <- t.Value
	walk(t.Right, ch)
}

func same(t1, t2 *tree.Tree) bool {

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() { // traverses first tree and closes channel afterwards
		walk(t1, ch1)
		close(ch1)
	}()

	go func() { // traverses second tree and closes channel afterwards
		walk(t2, ch2)
		close(ch2)
	}()

	for v1 := range ch1 {
		v2, ok := <-ch2 // checks if ch2 channel is still open

		if !ok || v1 != v2 {
			return false
		}
	}

	_, ok := <-ch2
	return !ok // checks if all values are consumed from the channel, if the ok is true then value is there and number of nodes are not equal hence return false, if it is false then the channel is consumed so number of nodes in both trees are same hence return true

}

func main() {

	t1 := tree.New(7)
	t2 := tree.New(7)

	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println("Are they equivalent :- ", same(t1, t2))

	t3 := tree.New(3)
	t4 := tree.New(4)

	fmt.Println(t3)
	fmt.Println(t4)
	fmt.Println("Are they equivalent :- ", same(t3, t4))

}
