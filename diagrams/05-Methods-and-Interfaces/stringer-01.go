package main
import "fmt"

type Stringer interface {
	String() string
}

type Book struct {
	Title string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s\n", b.Title, b.Author)
}


type Car struct {
	Name string
	Model string
}

func (c Car) String() string {
	return fmt.Sprintf("%s %s", c.Name, c.Model)
}

func main(){
	b := Book{Title: "Surely You're Joking Mr.Feynman", Author: "Richard"}
	c := Car{Name: "Hyundai", Model: "Creta"}

	fmt.Println(b, c)
}
