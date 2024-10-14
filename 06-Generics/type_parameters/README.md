# Go Type Parameters

In Go, **type parameters** are a key feature of **generics** that allow functions,, types, and methods to bbe written in a more general way. Instead of specifying concrete data types, you can use type parameters to create flexible, reusable code that works with multiple data types.

## What are Type Parameters?

Type parameters are like placeholders for actual types. When you define a function or a type using generics, you use type parameters to tell Go that it can accept any type that meets certain conditions (constraints). The actual type is specified later when the function or type is used.

**Syntax of Type Parameters**

Type parameters are defined inside square brackets `[]` after the function name or type name. For example:

```go
func Print[T any](value T) {
    fmt.Println(value)
}
```
In this example:
- `T` is a **type parameter**. It's like a variable that represents a type.
- `any` is a constraint, meaning `T` can be any type.
- The function `Print` takes an argument `value` of type `T`.

When you call `Print`, you specify the actual type for `T`:

```go
Print(42)    // Here T is int
Print("Hello Go!! ")    // Here T is string
```

## Type Parameters in Functions

Here's an example of a function that uses type parameters to work with different numeric types:

```go
package main
import "fmt"

// generic function that accepts any numeric type
func Add[T int | float64] (a, b T) T {
    return a + b
}

func main() {
    fmt.Println(Add(20, 40))    // T is int, and the result be: 60
    fmt.Println(Add(3.7, 7.8))    // T is float64, and the result be: 11.5
}
```
In this example:
- `T` is a type parameter that can be either `int` or `float64`
- The `Add` function works with both integers and floating-point numbers.

## Type Parameters in Types (Structs)

You can also use type parameters with Go structs. This allows you to define generic data structures. For example, yu could create a generic **Stack** that works with any data type:

```go
package main
import "fmt"

// Defining a generic Stack type
type Stack[T any] struct {
    items[]T
}

// Push adds an item to the stack
func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

// Pop removes and return the last item
func (s *Stack[T]) Pop() T {
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func main() {
    // Stack of Integers
    intStack := Stack[int]{}
    intStack.Push(10)
    intStack.Push(20)
    fmt.Println(intStack.Pop())    // Output: 20

    // Stack of strings
    stringStack := Stack[string]{}
    stringStack.Push("Hellp")
    stringStakc.Push("Go!!")
    fmt.Println(stringStack.Pop())    // Output: Go!!
}
```
In this example:
- `T` is a type parameter for the `Stack` struct.
- The `Struct` can work with any type, such as `int` or `string`.

## Constraints on Type Parameters

You can limit what types can be used with a type parameter by adding **constraints**. Constraints specify the set of types that are allowed.

For example, you can create a function that only works with types that support the `+` operator (like `int` or `float64`):

```go
func AddNumber[T int | float64] (a, b T) T {
    return a + b
}
```
In this case, `T` can only be `int` or `float64`, so you can't pass a `string` or other non-numeric type.

## Type Inference with Type Parameters

Go can often infer the type parameter from the arguments you pass. For example:

```go
func Print[T any] (value T) {
    fmt.Println(value)
}
```
You don't need to explicitly specify the type parameters in the function call, Go automatically figures it out based on the provided argument.

## Benefits of Type Parameters

- **Code Reusability:** You can write a function or data structure once, and it can work with any type.
- **Type Safety:** The compiler ensures that the types used with your function or type are correct, reducing the runtime error.
- **Flexibility:** Type parameters allow your code to handle a wide variety of types without having to write seperate versions of each one.

## Key Points
- **Type Parameters** allow you to define generic functions and types that can work with different data types.
- They are defined using square brackets `[]` after the function or type name.
- You can use **constraints** to limit which types can be passed as type parameters.
- Go can often infer the type parameters based on the arguments based on the arguments provided.
