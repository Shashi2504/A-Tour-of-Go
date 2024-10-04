package main
import "fmt"

func applyFunction(fn func(int) int, value int) int {
	return fn(value)
}

func double(x int) int {
	return x * 2
}

func main() {
	result := applyFunction(double, 5)
	fmt.Println(result)
}
