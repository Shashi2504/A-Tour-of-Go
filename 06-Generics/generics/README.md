# Go Generics

**Generics** in Go allows you to wirte functions, types, and data structures that can work with any data type, making your code more flexible and reusable. Introduced in Go 1.18, generics let you define a function or type that can handle values of different types without writing seperate code for each type.

## Why Generics?

Without generics if you want to write a function that works with different types (like `int`, `string`, `float64`), then you have to write a version of the function for each type. With generics, you have to write the function once, and it can work with any type.

## Syntax of Generics

To use generics in Go, you define a type parameter inside square brackets `[]` after the function name. The type parameters can be used within the function.

**Example of a Generic Function**

Here's an simple example of a generic function that works with any type of numbers:

```go
package main

import "fmt"

// A generic function that can take two values of any type
func add[T int | float64](a, b T) T {
    return a + b
}

func main() {
    fmt.Println(add(3, 4))         // Works with int
    fmt.Println(add(3.5, 4.7))     // Works with float64
}
```
Explanation:
T is the type parameter. It's a placeholder for the actual type.
int | float64 means that T can be either int or float64.
The add function can now handle both int and float64 types without needing two different functions.

## Constraints in Generics

You can restrict what types can be passed to the generic function uso
