package main
import "fmt"

func main() {
		
	// Here we have implemeted the "_" symbol which is skipping or ignoring right.
		meme := []int{1, 2, 4, 5, 7, 8, 9}

		for i, _ := range meme {
			fmt.Println("INdex:", i)
		}
	
	// Here check the actual difference where I didnt ignore anything so we can know exactly what was getting ignored before.
		for i, values := range meme {
		fmt.Println("Index:, Values:", i, values)
		}
}
