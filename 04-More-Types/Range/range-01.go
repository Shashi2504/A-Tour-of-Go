package main

import "fmt"

func main() {
    
    pow := make([]int, 10)

    for i := range pow {
        pow[i] = 1 << uint(i)  // This 1 shifted left by 'i' bits, which equals 2^i.
    }

    
    for _, value := range pow {
        fmt.Printf("%d\n", value)
    }
}

