# Go Constants

In Go, **constants** are immutable values that are known at compile time and do not change throughout the program's execution. They are typically used to represent fixed values like mathematical constants, configuration settings, or other data that should not be modified.

To declare a constant in Go, you can use the `const` keyword followed by the constant's name, an optional type, and its value. The value must be a compile-time constant, like a number or string.

### Example of Constant Declaration

```go
const pi float64 = 3.14159
const greeting = "Hello, Paaaaaaaaaaap!!!"
```
In this exmaple:
  - `pi` is a `float64` constant representing the value of Pi.
  - `greeting` is a string constant with the value `"Hello, Paaaaaaaaaaap!!!"`


## Numeric Constants

Numeric constants in Go are a bit more flexible than regular constants. They are special because they are **untyped** until they are directly used in a context that requires a specific type. This allows numeric constants to take on different types depending on the context in which they are used.

**Example of Numeric Constants*:*
```go
const z = 100     // z is an untyped constant
```
Here, `z` is a numeric constant with the value of `100`, but it is **untyped**. When you use `z` in a particular context, Go will indicate its type  based on how it's used.

### Example of Type Inference with Numeric Constants:

```go
var a int = z    // z is indicated as int
var b float64 = z    // z is indicated as float64
var c byte = z    // z is indicated as vyte
```
Even though `z` is declared as a constant without a type, Go allows it to be used as `int`, `float64`, or `byte`, depending on the context in which its used. This flexibility makes numeric constants very useful.

### Why Use Numeric Constants?
  - **Precision:** Go allows numeric constants to represent values with greater precision than any type can hold. For example, large or precise values of float64 or integers can be stored in numeric constants without loss of accuracy until they're assigned a specific type.

### Example with Operations

```go
const a = 50000
const b = 3.14

fmt.Println(a * 2)    // Go infers 'a' as int for this multiplication
fmt.Println(b * 2)    // Go infers 'b' as float64 for this multiplication
```

### Mutiple Constants

You can declare multiple constants at once using parentheses:

```go
const (
    maxConnections = 100
    timeout        = 60
    threshold      = 1.5
)
```

### Recap

  - Constants are immutable values declared using `const`.
  - Numeric constants are untyped until used, making them flexible for different types like `int`, `float64`, or `byte`.
  - Constants cannot be changed once assigned, ensuring that they remain fixed throughout your program.
