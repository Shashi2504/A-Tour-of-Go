package main

import "fmt"

func main() {
    a := make([]int, 5) 
    fmt.Println("Slice a:", a)                
    fmt.Println("Length of a:", len(a))       
    fmt.Println("Capacity of a:", cap(a))     

    b := make([]int, 0, 5) 
    fmt.Println("Slice b:", b)               
    fmt.Println("Length of b:", len(b))       
    fmt.Println("Capacity of b:", cap(b))     

    b = b[:cap(b)] 
    fmt.Println("Extended Slice b:", b)      
    fmt.Println("Length of b:", len(b))       
    fmt.Println("Capacity of b:", cap(b))     

    b = b[1:] 
    fmt.Println("After removing first element:", b) 
    fmt.Println("Length of b:", len(b))       
    fmt.Println("Capacity of b:", cap(b))     
}

