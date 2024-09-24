package main
import (
	"fmt"
	"math/cmplx"
)

var (
	MustWork bool = true
	NeedTo uint64 = 1<<64 - 2
	Me complex128 = cmplx.Sqrt(-24 + 56i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", MustWork, MustWork)
	fmt.Printf("Type: %T Vaule: %v\n", NeedTo, NeedTo)
	fmt.Printf("Type: %T Value: %v\n", Me, Me)
}
