package main
import "fmt"

type Multiply struct {
	X, Y float64
}

func (k *Multiply) Kolatha(d float64) {
	k.X = k.X * d
	k.Y = k.Y * d
}

func kolathaFunc(k *Multiply, d float64){
	k.X = k.X * d
	k.Y = k.Y * d
}

func main() {
	k := Multiply{5, 7}
	k.Kolatha(6)
	kolathaFunc(&k, 12)

	s := &Multiply{9, 2}
	s.Kolatha(7)
	kolathaFunc(s, 9)

	fmt.Println(k, s)
}
