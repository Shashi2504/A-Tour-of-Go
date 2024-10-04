package main
import "fmt"

type List[T any] struct {
	next *List[T]
	val T
}

func (l *List[T]) Insert(value T) {
	current := l
	for current.next != nil {
		current = current.next
	}
	current.next = &List[T]{val: value}
}

func (l *List[T]) Traverse() {
	current := l
	for current != nil {
		fmt.Println(current.val)
		current = current.next
	}
}

func (l *List[T]) Length() int {
	length := 0
	current := l
	for current != nil {
	length++
	current = current.next
	}
	return length
}


func main() {
	list := &List[int]{val: 1}

	list.Insert(2)
	list.Insert(3)
	list.Insert(4)

	fmt.Println("List contents:")
	list.Traverse()

	fmt.Println("Length of the list:", list.Length())
}
