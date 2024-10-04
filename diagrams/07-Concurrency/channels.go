package main
import "fmt"

func sum(s []int, c chan int) {
	total := 0
	for _, v := range s {
		total += v
	}
	c <- total
}

func main() {
	s := []int{1, 3, 4, 5, 7, 8, 9, 10, 12, 14}
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[:len(s)/2], c)

	x, y := <-c, <-c

	fmt.Println("Total will be:", x+y)
}
