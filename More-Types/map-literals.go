package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// This one type
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// Another type

var n = map[string]Vertex{
        "Shagaayi town Labs": {40.68433, -74.39967},
        "aoongaaan": {37.42202, -122.08408},
}


func main() {
	fmt.Println(m)
	fmt.Println(n)
}

