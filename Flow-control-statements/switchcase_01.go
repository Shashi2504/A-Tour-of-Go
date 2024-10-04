package main
import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("The current operating sysetm is")
	switch os := runtime.GOOS; os {
	case "Linux":
		fmt.Println("Penguine")
	case "windows":
		fmt.Println("Kitikeeelu")
	}
}
