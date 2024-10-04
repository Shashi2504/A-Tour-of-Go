package main
import "fmt"

func main() {

	lucky_num := [] int {2, 7, 6, 7}
	fmt.Println("Your lucky numbers are:", lucky_num)
	

	lucky_num = append(lucky_num, 9, 10, 67)
	fmt.Println("Adding few numbers to the lottery:", lucky_num)
}
