package main
import "fmt"

func split(sum int) (x, y, z int) {
	x = sum * 5/6
	y = sum - 10
	z = sum / 25
	return
}

func main() {
	fmt.Println(split(44))
}
