package main
import "fmt"

func main() {
	a, b := 23, 4506  // Here I have declared variables using shorthand. So "a" value will be 23 and "b" value will be 45.

	p := &a  // Here we are creating the pointer using the "&" operator. So now p will hold the vaule of "a" which is 23.
	
	fmt.Println(*p)  // Here we are accessing the created pointer using "*" opertor.


	*p = 56  // Before we have seen accessing the pointer. Now we will see how to change the value. So when we say "*p := 56" then this "*p" to "a". So it gonna change the value of "a" to 56.

	fmt.Println(a)  // Checking the value of "a"


	p = &b // Doing the same provess but here we are updating with "j" thats it.

	*p = *p / 76
	
	fmt.Println(b)
}
