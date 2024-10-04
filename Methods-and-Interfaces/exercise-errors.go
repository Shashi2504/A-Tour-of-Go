package main
import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e)) 
}


func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	result, err := Sqrt(2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of 2 is:", result)
	}

	result, err = Sqrt(-2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Square root of -2 is", result)
	}
}
