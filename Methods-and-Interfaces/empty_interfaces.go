package main
import "fmt"

func main() {
	var a interface{}
	
	a = 424
	fmt.Println(a)

	a = "Hello"
	fmt.Println(a)

	a = true
	fmt.Println(a)
}
