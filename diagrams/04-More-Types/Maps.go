package main
import "fmt"

func main() {

	wwe := map[string]int{
	
	"Roman Reigns": 92,
	"Cody Rhodes": 90,
	"John Cena": 90,
	"Kane": 100,
	}

	// Forgot to add one more player in the function no proble we can add like this

	wwe["Braun Stroman"] = 93

	fmt.Println("Here comes the big dog Roman Reigns with power of", wwe["Roman Reigns"])

	// Now checking with the names like they are exist in our list or not

	sup, exists := wwe["Kane"]
	if !exists {
		fmt.Println("Kane doesn't exist")
	} else {
		fmt.Println("Kane exists", sup)
	}

	for wwe, sup := range wwe {
		fmt.Printf("%s is %d power\n", wwe, sup)
	}
}
