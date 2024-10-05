package main
import (
	"fmt"
	"math"
)

func main() {
	var i, j int = 3, 5
	var f float64 = math.Sqrt(float64(i*i + j*j))
	var y uint = uint(f)
	fmt.Println(i, j, y)
}
