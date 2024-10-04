package main
import (
	"fmt"
	"math"
)

func pow(z, x, cim float64) float64 {
	if m := math.Pow(z, x); m < cim {
		return m
	} else {
		fmt.Printf("%g >= %g\n", m, cim)
	}
	return cim

}

func main() {
	fmt.Println(
	pow(4, 6, 9),
	pow(9, 10, 3),
	)
}
