# Go Fucntions

In Go, a **function** is a block of code that performs a specific task. Functions allow you to write reusable pieces of logic, making you code easier to maintain and understand.

## Defining a Function

To define a function in Go, you use the `func` keyword followed by the function name, a list of parameters (if any), and the return type (if any). The function body is enclosed by curly brraces `{}`.

Here's the basic structure:

```go
func functionName(paramerter) returnType {
      // Function body
}
```

### Example of a Simple Function:

```go
func greet(name string) string {
      return "Hello", + name
}
```
In this example:
  - `greet` is the function name
  - `name` parameter of type string `string`.
  - The function returns a `string` value.
  - The `return` keyword is used to send the result back to the caller.

## Calling a Function

To call a function, you just use it name followed by parenthesis with the required arguments (if any):

```go
greeting := great("Go mate")
fmt.Println(greeting)    // Output: Hello, Go mate
```

## Multiple Return Values

Go functions can return multiple values, which is useful for returning both a result and an error, or additional data. For example:

```go
func divide(a, b float64) (float64 error) {
      if b == 0 {
          return 0, fmt.Errof("cannot divide by zero")
      }
      return a/b, nil
}
```
In this function:
  - It takes two `float64` parameters.
  - It returns two values: a `float64` result and an `error`.

You can call this function like this:

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println("Result:", result)
}
```

## Variadic Functions

Go also supports **variadic functions**, which can take variable number of aruguments. You define these by `...` before the type of parameter. For example:

```go
func sum(numbers ...int) int {
      total := 0
      for _, num := range numbers {
          total += num
      }
      return total
}
```
Here, `sum` can accept any number of `int` arguments:
```go
fmt.Println(sum(1, 2, 3, 4, 5, 6))    // Output: 21
```

## Recap

  - Functions help you organize and reuse code.
  - You define a function with the `func` keyword, parameters, and a return type.
  - Functions can return mutliple values and handle variable number of arguments.
