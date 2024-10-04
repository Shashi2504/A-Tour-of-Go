package main
import "fmt"

func main() {
	var a[5]int

	a[0] = 10
	a[1] = 20
	a[2] = 30
	a[3] = 40
	a[4] = 50

	fmt.Println("a:", a)


	// Another way of creating an array
	array1 := [5] int {100, 200, 465, 242, 3234}
	fmt.Println("array1:", array1)

	// Creating another way

	HIMYM := [5] string {"Ted", "Marshal", "Lilly", "Barney", "Robin"}

	fmt.Println("How I Met Your Mother sitcom is all about:", HIMYM)
	fmt.Println("Here comes the guy who always hookup with girls and think one of them is his dream gaaarl:", HIMYM[0])
	fmt.Println("The Aw wait for it some guy in the whole sitcom is:", HIMYM[3])
}
