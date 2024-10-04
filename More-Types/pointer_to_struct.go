package main
import "fmt"

type Book struct {
	Title string
	Author string
}

func main() {
	
	mybook := Book {
	Title: "Wings of Fire",
	Author: "A.P.J Abdul Kalam",
	}
	
	a := &mybook

	fmt.Println("Author is", a.Author)
	fmt.Println("Book name is", a.Title)
}
