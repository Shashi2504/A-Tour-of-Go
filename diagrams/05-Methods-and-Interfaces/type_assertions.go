package main
import "fmt"


func main() {
	var k interface{} = "Hiuhhhh chaari gaaru"

	d := k.(string)
	fmt.Println(d)

	d, ok := k.(string)
	fmt.Println(d, ok)

	sr, ok := k.(float64)
	fmt.Println(sr, ok)

	sr = k.(float64)
	fmt.Println(sr)

}
