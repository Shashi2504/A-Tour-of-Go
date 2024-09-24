package main
import "fmt"

type fictional struct {
	Marvel	string
	DC	string
	StarWars string
}

func main() {
	
	characters := fictional{
	Marvel: "Avengers",
	DC: "Batman",
	StarWars: "The Force",
	}

	fmt.Println("Marvel Team:", characters.Marvel)
	fmt.Println("DC Team:", characters.DC)
	fmt.Println("StarWars Team:", characters.StarWars)
}
