package main
import (
	"fmt"
	"math"
)

type coli interface {
	Abs() float64
}

type keer float64

type sreen struct {
	X, Y float64
}


func (k keer) Abs() float64 {
	if k < 0 {
		return float64(-k)
	}
	return float64(k)
}

func (s sreen) Abs() float64 {
	return math.Sqrt(s.X*s.X+s.Y*s.Y)
}

func main() {
	var a coli
	k := keer(-math.Sqrt2)
	s := sreen{3, 4}

	a = k
	a = &s

	a = s

	fmt.Println(a.Abs())
}
