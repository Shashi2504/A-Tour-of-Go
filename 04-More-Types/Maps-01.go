package main
import "fmt"

func main() {
	batman := make(map[string]int)

	batman["age"] = 26
	batman["year"] = 1972

	fmt.Println("Batman age is:", batman["age"])
	fmt.Println("Batman was born in the year:", batman["year"])


	 _, exists := batman["cat woman"]
	if !exists  {
		fmt.Println("cat woman doesn't exist")
	} else {
		fmt.Println("Cat woman exists")
	}
}
