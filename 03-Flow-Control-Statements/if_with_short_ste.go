package main
import (
	"fmt"
	"math"
)

func pow(a, b, bin float64) float64 {
	if  v := math.Pow(a, b); v < bin {
	return v
	}
	return bin
}

func main() {
	fmt.Println(
	pow(3, 2, 10),
	pow(3, 3, 20),
	)
}
