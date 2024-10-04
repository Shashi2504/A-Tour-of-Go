package main

import "fmt"

func main() {
    a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
    
    slice1 := a[0:10]
    slice2 := a[:10]
    slice3 := a[0:]
    slice4 := a[:]
    
    fmt.Println("Slice 1:", slice1)
    fmt.Println("Slice 2:", slice2)
    fmt.Println("Slice 3:", slice3)
    fmt.Println("Slice 4:", slice4)
}

