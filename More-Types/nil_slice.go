package main

import "fmt"

func main() {
    var nilSlice []int

    fmt.Println("Nil Slice:", nilSlice)            
    fmt.Println("Length of Nil Slice:", len(nilSlice))  
    fmt.Println("Capacity of Nil Slice:", cap(nilSlice)) 

    if nilSlice == nil {
        fmt.Println("The slice is nil.")
    } else {
        fmt.Println("The slice is not nil.")
    }
}

