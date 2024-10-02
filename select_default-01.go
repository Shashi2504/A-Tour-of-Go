package main
import (
	"fmt"
	"time"
)

func main() {
	tic_tic := time.Tick(100 * time.Millisecond)
	boom := time.After(500 *  time.Millisecond)
	for {
	select {
		case <- tic_tic:
			fmt.Println("tic-tic-ticuihhhhh")
		case <-boom:
			fmt.Println("BooooooooooooooM")
		default:
			fmt.Println("        .")
			time.Sleep(50 * time.Millisecond)
	}
    }
}
