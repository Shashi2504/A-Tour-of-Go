package main
import "fmt"

type meh interface {
	ace()
}

type spike struct {
	nishi string
}

func (s *spike) ace() {
	if s == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(s.nishi)
}

func main() {
	var a meh
	var s *spike
	
	a = s
	describe(a)
	a.ace()

	a = &spike{"Happyuhhhhhhhh"}
	describe(a)
	a.ace()
}

func describe(a meh) {
	fmt.Printf("(%v, %T)\n", a, a)
}
