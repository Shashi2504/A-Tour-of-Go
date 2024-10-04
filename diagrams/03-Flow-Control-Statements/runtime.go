package main
import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	fmt.Println("Your operating  system is:", os)

	info := runtime.GOARCH
	fmt.Println("Detail of the operating system:", info)

}
