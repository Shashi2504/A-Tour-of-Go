package main
import "fmt"


// // The function adder creates and returns a closure (another function)

func adder() func(int) int {
	sum := 0
	return func(x int) int {   // So this is the closure
		sum += x
		return sum
	}
}


func main() {

	addONe := adder()
	addTWo := adder()

	fmt.Println(addONe(2))
	fmt.Println(addONe(9))


	fmt.Println(addTWo(2))
	fmt.Println(addTWo(3))
}
