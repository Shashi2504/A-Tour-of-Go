package main
import "fmt"

func main() {
	array := [7] int {3, 5, 6, 8, 10, 23, 34}

	slice1 := array[1:4]

	fmt.Println("Slice1 is:", slice1)
	fmt.Println("Slice1 is:", len(slice1))
	fmt.Println("Slice1 is:", cap(slice1))

	extendSlice := slice1[0:3]
	fmt.Println("Extended Slice1 is:", extendSlice)
}
