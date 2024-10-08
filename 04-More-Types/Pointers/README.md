# Go Pointers

In Go, **pointers** allow you to reference the memory address of a value rather than the value itself. This is useful when you want to modify a variable's value inside a function or method or when you want to share large data strucutres efficienetly copying them.

## What is a Pointer?

A **pointer** is a variable that stores the memory address of another variable. In Go, you declare a pointer using the `*` symbol, and the `&` symbol is used to get the memory address of a variable.

### Example of Pointers

Here's a simple example of how pointers work:'

```go
var x int = 10    // Declare an integer variable
var p *int = &x   // Declarea pointer that holds the memory address of `x`

fmt.Println(x)    // Output will be: 10 (value of x)
fmt.Println(p)    // Output will be: (memory address of x)
fmt.Println(*p)   // Output will be: 10 (the value at the memory address p points to)
```
- `p` is a pointer to x, meaning `p` stores the address of `x`.
- `*p` is called dereferencing the pointer, which allows you to access the value stored at the address `p` points to.

## Why Use Pointers?

**1. Modifying Values:** If you pass a variable in a fucntion, a copy of the value is passed, and changes indside the function won't affect the original variable. However, with pointers, you can modify the original variable.

```go
func modify(val *int) {
    *val = 20    // Dereference the pointer to modify the value
}

x := 10
modify(&x)    // Pass the address if x
fmt.Println(x)    // Output will be: 20 (x was modified)
```
In this example, passing `&x` (the address of `x`) allows the `modify` function to change the value of `x` directly.

**2. Efficient Memory Usage** Pointers allow you to pass large structures or arrays without copying them. Instead of creating a duplicate in memory, you simply pass the address of the structure, which is more efficient.

## Zero Value of Pointer

The zero value of a pointer is `nil`, which means it doesn't point to any memory address.

```go
var p *int
fmt.Println(p)    // Output will be: <nil> (it points out to nothing)
```

## Poiinters to Struct

Pointers are often used with structs to modify the fields of the orignal struct:

```go
type Person struct {
      Name string
      Age int
}

func increaseAge(p *Person) {
      p.Age++    // Modify original struct using pointer
}

p := Person {
      Name: Noharaa,
      Age: 30
}

increaseAge(&p)    // Pass the address of the struct
fmt.Println(p.Age)    // Output: 31 (Age was incremented)
```
Here, passing the pointer to `p` allows the `increaseAge` function to modify the original `Person` struct.

## No Pointer Arithemcatic in Go

Unlike languages like C or C++, Go does not support pointer arithematic (e.g. increamenting a pointer to move the next memory allocation). This design decision helps prevent many common programming errors.

## Recap
- A pointer holds the memory address of another variable.
- Use `&` to get the memory address of a variable.
- Use `*` to dereference the pointer and access pr modify the value.
- Pointers help modify variables inside fucntions, and they are memory-efficient for passing large data structures.
