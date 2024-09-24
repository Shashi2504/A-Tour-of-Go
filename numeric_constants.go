package main 
import "fmt"


const (
	Big = 1 << 100
	Small = Big >> 99
)

func meInt(x int) int { 
	return x*24 + 56 
}

func meFloat(x float64) float64 { 
	return x*34.6 
}

func main() {
	fmt.Println(meInt(Small))
	fmt.Println(meFloat(Small))
	fmt.Println(meFloat(Big))
}
