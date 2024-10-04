package main
import "fmt"

func Index[T comparable] (s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	ints := []int {213, 123, 134, 656, 3168}
	fmt.Println(Index(ints, 134))

	strings := []string {"tum tum", "Moose", "chahahhhh"}
	fmt.Println(Index(strings, "Moose"))

	flaotas := []float64 {23.123, 24.31, 657.342}
	fmt.Println(Index(flaotas, 24.31))


	fmt.Println(Index(ints, 100))
}
