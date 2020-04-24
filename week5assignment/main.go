// Linked List in Golang
package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	data interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Add(data interface{}) {
	list := &Node{
		next: L.head,
		data: data,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.data)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.data)
		list = list.next
	}
	fmt.Println()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.data)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}

func main() {
	link := List{}
	link.Add(5)
	link.Add(9)
	link.Add(13)
	link.Add(22)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", link.head.data)
	fmt.Printf("Tail: %v\n", link.tail.data)
	link.Display()
	fmt.Println("\n==============================\n")
	fmt.Printf("head: %v\n", link.head.data)
	fmt.Printf("tail: %v\n", link.tail.data)
	link.Reverse()
	fmt.Println("\n==============================\n")
}
