# Go Variables

In Go, **variables** are used to store values that can be used and manipulated in your program. A variable has a name, a type, and can hold a value of that type. You need to decalre a variable before you can use it.

## Declaring Variables

You declare a variable in Go using the `var` keyword followed by the variable name and its type. Here's the basic syntax:

```go
var variableName variableType
```
Example:
```go
var age int
```
This declares a variable named  `age` of type `int` (integer). At this point, age is assiged a default value of `0` because it hasn't been given any other value yet.

## Intialzing Variables

You can also decalre and intialize a variable at the same time:

```go
var name string = "Paaaapppp!!!"
```
In this case, the variable `name` is of type `string`, and it's intialized with the value `"Paaaapppp!!!"`.

## Short Declaration

Go allows a shorthand for variable declaration and initialization using the `:=` operator. With this syntax, Go conclude the variable's type based on the value you assign:

```go
count := 10
```
Here, `count` is automatically assigned the type `int` because `10` is an integer.

## Multiple Variable Declaration

You can declare multiple variables of the same type in one line:

```go
var x, y, z int
```
Or, you can declare and initalize multiple variables at once:

```go
var a, b, c = 1, 2, "three"
```

## Variable Types

Go is statically typed, meaning the type of variable must be specified, and it cannot change later. Some common types are:

  - `int` (for integers)
  - `float64` (for decimal numbers)
  - `string` (for text)
  - `bool` (for true/false values)

## Zero Values

If you declare a variable without initializing it, Go assigns a **zero value** based on the variable type:
  - `0` for numeric types (int, float64, etc.)
  - `""` (empty string) for strings
  - `false` for booleans
  - `nil` for pointers, slices, maps, etc.

### Example of Zero Values

```go
var isAvailable bool
fmt.Println(isAvailable)    // Output: false
```

## Constants

You can also declare **constants** using the `const` keyword. Constants are variables whose values cannot be changed:

```go
const pi = 3.14
```

## Recap

  - Variables are declared using `var` and you can use `:=` for shorthand declarations.
  - Go variables are statically typed, and types are either directly stated.
  - Variables have zero values if not initialized.
  - Constants are immutable and declared with `const`.
