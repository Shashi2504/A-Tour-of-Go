package main
import "fmt"

func main() {
	sum := 4
	for sum < 1000 {
	sum += sum
	}
	fmt.Println(sum)
}
