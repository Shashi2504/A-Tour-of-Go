package main

import "fmt"

// A function that returns another function
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor 
	}
}

func main() {
	double := makeMultiplier(2) 
	triple := makeMultiplier(3) 

	fmt.Println(double(5)) 
	fmt.Println(triple(5)) 
}

