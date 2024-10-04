package main
import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is my Monday")

	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today")

	case today + 1:
		fmt.Println("Tomorrow")

	case today + 2:
		fmt.Println("After two days")

	default:
		fmt.Println("Tooooo farrrrrhuuhhhhhhh")
	}
}
