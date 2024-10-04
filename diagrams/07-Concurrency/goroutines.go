package main
import (
	"fmt"
	"time"
)

func wait() {
	fmt.Println("Tough things takes time to solve")
}

func main() {
	go wait()
	fmt.Println("Easy peasy")

	time.Sleep(1 * time.Second)
}
