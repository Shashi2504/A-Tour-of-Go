package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	cal  T
}

func (M *List[T]) Insert(value T) {
	ippudu := M
	for ippudu.next != nil {
		ippudu = ippudu.next
	}
	ippudu = &List[T]{cal: value}
}

func (M *List[T]) Traverse() {
	ippudu := M
	for ippudu.next != nil {
		fmt.Println(ippudu.cal)
	}
	ippudu = ippudu.next
}

func (M *List[T]) Length() int {
	length := 0
	ippudu := M
	for ippudu != nil {
		length++
		ippudu = ippudu.next
	} 
	return length
}

func main() {
	list := &List[int]{cal: 1}
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	
	fmt.Println("List of contents:")
	list.Traverse()
	
	fmt.Println("Length of the list:", list.Length())
}

