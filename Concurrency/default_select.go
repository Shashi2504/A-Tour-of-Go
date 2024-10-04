package main
import "fmt"

func main() {
	ch := make(chan int)

	select {
		case msg1 := <-ch:
			fmt.Println("Received message:", msg1)
		default:
			fmt.Println("No message was received so it was empty")
	}
}
