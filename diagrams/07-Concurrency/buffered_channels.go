package main
import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <-1
	ch <-2

	fmt.Println("Buffer is full now!")

	ch <-3

	fmt.Println("This line ownt be reached until there is some space in the buffer")
}
