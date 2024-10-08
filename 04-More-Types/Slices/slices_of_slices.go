package main
import "fmt"

func main() {
	
	table := [][] int {
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("The second row in the table which is counting with index:", table[2])
	fmt.Println("Then doing the slices of slices:", table[1][2])
}
