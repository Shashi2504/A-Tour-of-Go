package main
import "fmt"

type Make interface {
	Ice()
}

type fake struct {
	Sea string
}

type hein float64

func (f *fake) Ice() {
	fmt.Println(f.Sea)
}

func (h hein) Ice() {
	fmt.Println(h)
}


func main() {
	 var ming Make

	 ming = &fake("HUHUHHUHUUHHUUHHHHHHHhhhhh")
	 describe(ming)
	 ming.Ice()

	 ming = hein(math.Pi)
	 describe(ming)
	 ming.Ice()
}

func describe(ming Make) {
	fmt.Println("()")
}
