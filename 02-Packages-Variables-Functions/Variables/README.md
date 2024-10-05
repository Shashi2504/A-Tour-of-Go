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

## Type Conversions

In Go, **type conversions** allows you to change the type of a value from one type to another. Go is a statically typed language, meaning variables have a specific type, and you cannot automatically assign a value of one type to a variable of another type without directly converting it.

### Basic Syntax of Type Conversion

To convert a value from one type to another, you use the followinng syntax:

```go
newType(value)
```
Here, `newType` is the type you want to convert the value to, and `value` is the data you're converting.

### Example

```go
var a int = 42
var b float64 = float64(a)
```
In this example:
  - `a` is the integer (`int`), and we're converting it into a floating-point number (float64).
  - After conversionm the value of `b` will be `42.0` as a float.

## Main Points

  **1. Explicit Conversion Required:** Go does not perform automatic type conversion between different types, even if it seems like they might be compatible (e.g. between `int` and `float64`). You must directly convert between types.

  **Example**

  ```go
var x int = 10
var y float64 = 5.6
// z := x + y    // This will cause an error
z := float64(x) + y    // Correct: directly convert 'x' to float64
```
  **2. Conversion Between Numeric Types:** You can convert between numeric types like `int`, `float64`, `int32`, etc. but you have to be careful about loss of clarity.

  **Example**

  ```go
var f float64 = 9.8
var i int = int (f)    // Converting float64 into int shorten the decimal part
fmt.Println(i)    // Output: 9
```
In this case, the decimal part is dropped because an integer cannot hold fractional values.

  **3. String Conversion:** Converting a integer or a floating-point value to a string requires special handling using the  `strconv` package. Howecer, coverting a byte to a string is simple:

  **Example**

  ```go
var b []byte = []byte {72, 101, 108, 108, 111}
var s string = string(b)
fmt.Println(s)    // Output: Hello
```
  **4. Boolean Conversion:** There's no direct conversion between `int` and `bool`. Instead, you can manually check conditions to assign boolen values.

  **5. Type Conversion Between Custom Types:** If you define your own types, you can convert values between those types if they share the same underlying type.

### Example with `strconv` for String Conversion

  If you want to convert a number to a string, you can use the `strconv` package:
  ```go
import "strconv"

var num int = 149
var text string = str.conv.Itoa(num)    // Itoa converts int to string
fmt.Println(text)
```

## Type Inference

In Go, **type inference** is a feature that allows a Go compiler to automatically determine the type of variable based on the value assigned to it, rather than requiring you to directly specify the type. This makes variable declarations more concise and easier to read.

### How Type Inference Works

When you use the short variable declaration syntax `:=`, the Go compiler indicate the type from the right-hand side expression. Here's how it works:

### Example of Type Inference

```go
x := 42    // x is indicated as int
y := 3.14    // y is indicated as float64
name :=  "Paaaaaaaapp!!"    // name is indicated as string
isAvailable := true    // isAvailable is indicated as boolean
```
In this example:
  - `x` is indicated to be of type `int` because it's assigned an integer value
  - `y` is indicated to be `float64` because it's assigned a floating-point value
  - `name` is indicated as `string` and `isAvailable` as `bool`

### Points to Note

  **1. No Direct Type Required:** You don't need to directly specify the type when using `:=`. The compiler takes care of it, which simplifies code writing.
  
  **2. Type Inference Only on Declaration:** Type Inference happens only when you are declaring a new variable. If you use the `var` keyword for an existing variable, you must specify its type:

  ```go
var a = 10    // Here, a is indicated as int
var b int    // Here, you must explicitly specify the type
```
  **3. Complex Expressions:** Type inference works well with more complex expressions too. For example:

  ```go
c := 10 + 5.5    // c is indicated as float64
```
In this case, the result of `10 + 5.5` is a floating-point number, so `c` becomes `float64`.

  **4. Composite Types:** Type inference also works with composite types like slices, maps, and structs.

  ```go
mySlice := []int{1, 2, 3}    //mySlice is indicated as []int
myMap := map[string]int{"Level": 1}    // myMap is indicated as map[string]int
```

## Recap

  - Variables are declared using `var` and you can use `:=` for shorthand declarations.
  - Go variables are statically typed, and types are either directly stated.
  - Variables have zero values if not initialized.
  - Type conversion in Go is explicit; you must use `newType(value)` to convert between types.
  - For string and numeric conversions, the `strconv` package is often used.
  - Type inference in Go simplifies variable declarations by allowing the compiler to automatically determine the type based on the assigned value.
