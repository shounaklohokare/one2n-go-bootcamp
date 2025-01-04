package main

import "fmt"

type List[T comparable] struct {
	head *Node[T]
}

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

func (list *List[T]) isEmpty() bool {

	if list.head == nil {
		return true
	}

	return false

}

func (list *List[T]) Add(value T) {

	newNode := &Node[T]{value: value}
	if list.isEmpty() {
		list.head = newNode
		return
	}

	currentNode := list.head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode

}

func (list *List[T]) Remove(value T) bool {

	if list.isEmpty() {
		return false
	}

	if list.head.value == value {
		list.head = list.head.next
		return true
	}

	current := list.head
	for current.next != nil && current.next.value != value {
		current = current.next
	}

	if current == nil {
		return false
	}

	current.next = current.next.next
	return true

}

func (list *List[T]) Print() {

	if list.isEmpty() {
		fmt.Println("Couldn't Print, Linked List is empty!")
		return
	}

	current := list.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}

	fmt.Println("nil")

}

func main() {

	list := List[int]{}

	fmt.Println("Is the list empty?", list.isEmpty())
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()

	list.Remove(2)
	list.Print()

	list.Remove(1)
	list.Print()

	fmt.Println("isEmpty?", list.isEmpty())
	list.Remove(3)
	list.Print()

	fmt.Println("isEmpty?", list.isEmpty())

}
