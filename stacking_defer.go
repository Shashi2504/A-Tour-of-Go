package main
import (
	"fmt"
	"runtime"
)

func main() {

	switch os := "Linux"; os {
	case "Linux":
		fmt.Println("Penguin")

	case "Windows":
		fmt.Println("Ktitititkeeelu")

	fmt.Println("Checking the os animal name funny right lets see:")
	}


	for i := 0; i < 10; i++ {
	defer fmt.Println(i)
	}

	fmt.Println("done")

	defer fmt.Println(runtime.GOOS)
}
