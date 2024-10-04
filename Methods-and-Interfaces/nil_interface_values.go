package main
import "fmt"

type wah interface {
	awee()
}

func main() {
	var a wah
	describe(a)
	a.awee()
}

func describe(a wah) {
	fmt.Printf("(%v, %T)\n", a, a)
}
