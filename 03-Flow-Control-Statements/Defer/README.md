# Go Defer

The `defer` **statement** in Go used to schedule a function or a statement to be executed after the surrounding function completes, regardless of whether the function exits normally or due to an error. Deferred functions are commonly used for cleanup tasks like closing files, network connections, or releasing resources.

## Basic Syntax

```go
defer functionCall()
```
- The deferred function is not executed immediately. Instead, it is added to a list of deferred calls.
- When the surrounding funciton finishes (either by returning or due to an error), the deferred function is executed.

### Example

```go
package main'

import "fmt"

func main() {
    fmt.Println("Start")
    defer fmt.Println("Deferred message")
    fmt.Println("End")
}
```
**Output**

```go
Start
End
Deferred message
```
**Explanation**

- "Start" is printed first.
- Then, the `defer` schedules the message " Deferred message" to be printed after the main() function completes.
- "End" is printed next, and then the deferred message is printed at last, right before the function exists.

## Multiple `defer` Statements

If you have multiple deferred calls, they are executed in **last-in, first-out (LIFO)** order.

**Example**

```go
func main() {
    defer fmt.Println("First")
    defer fmt.Println("Second")
    defer fmt.Println("Third")
    fmt.Println("End of the function")
}
```
**Output**
```go
End of the function
Third
Second
First
```

**Explanation**

The deferred functions are stacked, and when the function ends, they are executed in reverse order: "Third", "Second", "First".

## Common Use Cases: Closing Resources

A common use for `defer` is to to ensure resources like files or database connnections are properly closed when you are done using them.

**Example**

```go
package main

import (
      "fmt"
      "os"
)

func main() {
      file, err := os.Open("example.txt")
      if err != nil {
          fmt.Println(err)
          return
    }
    defer file.Close()    // Ensure the file is closed at the end of the function

    // Use the file (e.g read from it)
    fmt.Println("Reading the file.......")
}
```
In this example:
- The file is opened.
- The `defer file.Close()` ensures that the file is closed once the `main` function is done, whether it finishes normally or due to an error.
- This ensures proper resource management.

## Recap

- **Defer** schedules functions to run after the surrounding function finishes.
- Multiple deferred fucntions are executed in reverse order (LIFO).
- `defer` is typically used for cleanup tasks like closing files, releasing locks, or freeing resources.
