package main
import (
	"fmt"
	"math"
)

func sqrt(z float64) string {
	if z < 0 {
	return sqrt(-z) + "i"
	}
	return fmt.Sprint(math.Sqrt(z))
}

func main() {
	fmt.Println(sqrt(-12), sqrt(4))
}
