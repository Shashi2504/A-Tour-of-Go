# Go Named Returned Values

In Go, **named return values** allow you to name the return variables in the functions signature. This feature can make your code more readable, especially in functions that have multiple return values. When using named return values, you don't have to directly declare return variables inside the function-they are automatically declared and can be directly assigned within the function body.

## Syntax of Named Return Values

Here's the basic structure:

```go
func functionName(parameters) (returnType1, returnType2) {
      // You can directly assign values to the named return variables
      return
}
```

### Example

```go
func divide(a, b float64) (result float64, err error) {
      if b == 0 {
          err = fmt.Errorf("cannot divide by zero")
          return
      }
      result = a / b
      return
}
```
In this example:
  - The function `divide` returns two values: `result` (a `float64`) and `err` (an `error`).
  - These return values are named in the function signature: `result` and `err`
  - Inside the function, you can assign values to the `result` and `err` directly.
  - The `return` statement without any argumetns returns the named variables (`result` and `err`).

## Benefits of Named Return Values

  1. **Improves Readability:** The names of the return values provide context about what each value represents, making the function easier to understand.
  2. **Shorter Return Statements:** When using named return values, you don't need to specify the return values directly in the `return` statement-jsut `return` is enough, as Go automatically returns the named values.

### Example Without Named Return Values

For comparison, here's the same divide function without named return values:

```go
func divide(a, b float64) (float64, error) {
      if b == 0 {
          return 0, fmt.Println("cannot divide by zero")
    }
    return a/b, nil
}
```
Both versions work the same way, but named return values can make the above code clearer in more complex functions.

## Recap

  - Named return values allow you to name the return variables in the function signature.
  - You can assign values directly to the named return variables within the funcion.
  - A simple `return` statement will automatically return the named variables.
