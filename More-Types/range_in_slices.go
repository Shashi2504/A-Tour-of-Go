package main
import "fmt"

func main() {

	slice := [] string {"a", "b", "c"}
	for index, value := range slice {
	fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
}
