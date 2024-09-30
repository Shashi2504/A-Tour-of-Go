package main
import "fmt"

func doThis(a interface{}) {
	
	switch v := a.(type) {
	
	case string:
		fmt.Println("Work Hard duddeeee:\n", v)

	case int:
		fmt.Println("The number you need to do daily work is\n", v/7)

	default:
		fmt.Println("The following type is\n:")
	}
}

func main() {
	doThis("Ashhhhh work")
	doThis(28)
	doThis(3.14)
}
