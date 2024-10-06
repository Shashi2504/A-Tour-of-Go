package main
import (
	"fmt"
	"time"
	"runtime"
)

func main() {
	defer fmt.Println(time.Now())

	fmt.Println(runtime.GOOS)

	defer fmt.Println("bsdiiikeeeee")
}
