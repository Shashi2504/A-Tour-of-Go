# Go Interfaces

In Go, **interfaces** are a powerful feature that allow you to define a set of method signatures (withour implemetation) that types can implement. This enables Go to support polymorphism, meaning that different types can be treated the same way if they implement the same interface.

## What is an Interface?

An **interfce** in Go is a type that specifies a set of methods that other types must implement to be considered as implementing the interface. Unlike some other languages, Go's interface are implicit, meaning that a type automatically implements an interface if it has the required methods.

Here's a basic example:

 ```go
type Speaker interface {
    Speak() string
}
```
In this example, the `Speaker` interface requires a `Speak` method that returns a string. Any type that has a method with this signature is considered to implement the `Speaker` interface.

## Implementing an Interface

To implement an interface in Go, a type needs to provide all the methods defined by that interface. Let's create two different types that implement the `Speaker` interface:

```go
type Person struct {
    name string
}

func (p string) Speak() string {
    return "Hehheeeee", + p.name
}

tyep Dog struct {
    breed string
}

func (d Dog) Speak() string {
    return "Woof!! Iam a" + d.breed
}
```
Here, both the `Person` and `Dog` types implement the `Speak` method, so they both implicitly implements the `Speaker` interface.

## Using Interfaces

Once a type implements an interface, you can use that type whenever the interface is expected. This allows for flexible code, as different types that implement te same interface can be used interchangeably.

Here's how you might use the `Speaker` interface:

```go
func introduce(s Speaker) {
    fmt.Println(s.Speaker())
}

func main() {
    person := Person{name: "Ameeer"}
    dog := Dog{breed: " Lab"}

    introduce(person)    // Output: Hehheeeee, Ameeer
    introduce(dog)    // Output: Woof!! Iam a Lab
}
```
In this example, both `Person` and `Dog` are passed to `introduce` fuction because they both implement the `Speaker` interface. The function doesn't need to know the exact type of the argument, it only cares that the type implements the `Speaker` interface.

## Empty Interfaces

 In Go, the empty interfaces `interface{}` is special because it can be implemented by any type. This makes it useful for handling the values of unknown type, similar to `Object` in other langugaes:

 ```go
func printVal(v interface{}) {
    fmt.Println(v)
}
```
Since every type satifies the empty interface, you can pass any value to the `printVal` function.

## Interface Valus

When you assign a value to an interafce, Go stores both the **value** and the **type** in the interface variable. This enables Go to use the correct method implementation for the specific type at runtime.

Example:

```go
var S Speaker

s = Person{name: "Beeru"}
fmt.Println(s.Speak())    // Output: Hehheeeee, Beeru

s = Dog{breed: "Beagle"}
fmt.Println(s.Speak())    // Output: Woof!! Iam a Beagle
```
Here's the `Speaker` interface variable `s` can hold different types that can implement the `Speaker` interface, the correct method is called for each type.

## Type Assertions

If you have an interface value and wnat to access the underlying value or type, you can use **type assertions**. This allows you to retrieve the concrete type from an interface.

```go
func describe(s Speaker) {
    if p, ok := s.(Person); ok {
        fmt.Println("This is a person named", p.name)
    } else {
        fmt.Println("Unknown type")
    }
}

describe(Person{name: "Reveee"})    // Output: This is a person named Reveee
```
In this example, we check if the `Speaker` interface contains a `Person` type using `s.(Person)`. If it does, we can access the specific `Person` fields.

## Type Switches

A **type switch** in Go is a special form of the `switch` statement that allows you to compare the dynamic type of an interface value against several potential types. It's particularly useful when you have an interface and want to perform different actions based on the concrete type stored in that interface.

## How a Type Switch Works

A type switch uses the `.(type)` syntax to inspect the type of a value held by an interface. You can then match the dynamic type to different cases in thhe switch statement.

Here's the general form of a type switch:

```go
switch v := i.(type) {
case int:
    fmt.Println("i is an integer:", v)
case string:
    fmt.Println("i is a string:", v)
case bool:
    fmt.Println("i is a boolean:", v)
default:
    fmt.Println("i is of an unknwwon type")
}
```
- i.(type) is the special syntax used in type switches.
- `v := i.(type)` assigns the value `i` to `v` with the concrete type found in each case.
- Each `case` checks if `i` is a particular type.
-  If none of the cases match, the `default` block runs.

**Example of Type Switch in Action**

Let's take a look at a real example:

 ```go
func printType(i interafce{}) {
    switch v := i.(type) {
    case int:
      fmt.Printf("Integer: %d\n", v)
    case string:
      fmt.Printf("String: %s\n", v)
    case bool:
      fmt.Printf("Boolean: %t\n", v)
    default:
      fmt.Println("Unknown type")
   }
}

func main() {
    printType(24)   // Output: Integer: 24
    printType("Hello")  // Output: String: Hellp
    printType(true)   // Output: Boolean: true
    printType(3.14)   // Output: Unknown type
}
```
Here, the `printType` function accepts an `interface{}`, which can hold any type. The type switch then mathes the type of the argument to different cases : `int`, `string`, and `bool`. If none of these types match (like when passing 3.14, which is a float), the `default`case is executed.

## Key Points

- **Type switches** allow you to perform actions based on the type of an interface value.
- You can handle multiple values in a clean and readable way.
- The special `.(type)` syntax is used to extarct the dynamic type of an interface.
- **Default case** is used when none of the specified types match.

## Why Use Type Switches?

Type switches are particularly useful when working with **interfaces** that can hold different types. Instead of susing multiple type asserions (`i.(type)`) with `if` statements, you can handle everything neatly with one `switch`.

## Comparison with Type Assertions

- **Type assertions:** Used when you expect the interface to hold one specific type.
  ```go
  s, ok := i.(string)
  if ok {
    fmt.Println("i is a string:", s)
  }
  ```
- **Type switch:** Used when you want to handle multiple types:
  ```go
  switch i.(type) {
  case string:
    fmt.Println("i is a string")
  case int:
    fmt.Println("i is a integer")
  }
  ```
Both approches are useful, but type switches provide a more elegant solution when you have multiple possible types.

# Go Stringers

In Go, a **Stringer** is an interface from the `fmt` package that defines how a type should be represented as a string when printed using functions like `fmt.Prinln()`, `fmt.Prinf()` etc. It allows you to control the output when your custom types are printed.

## Stringer Interface 

The `Stringer` interface is very simple, consisting of just one method:

```go
tyep Stringer interface {
   String() string
}
```
Any type that implements this `String()` method automatically satisfies the `fmt.Stringer` interface. When you print values of this type using `fmt` functions, Go will call the `String()` method to determine the string representation.

## Example of Impelementing `Stringer`

Here's an example where we  define a custom type and implement the `String()` method:

```go
package main

imprt "fmt"

// Define a custom type
type Person struct {
    Name string
    Age int
}

// Implement the Stringer interafce for the Perso type
func (p Person) String() string {
   return fmt.Printf("%s is %d years old", p.Name, p.Age)
}

func main() {
    p := Person{Name: "Heeru", Age: 36}

    // When printing p, Go will automatically use the String() method
    fmt.Println(p)    // Output: Heeru is 36 years old
}
```
In this example:
- We creatd a custom type `Person`.
- We implemented the `String()` method to specify how a `Person` should be printed as: "Name is Age years old".
- When `fmt.Println(p)` is called, it prints the result of `p.String()` automatically.

## Why Use the `Stringer` interface?

1. **Custom Formatting:** It allows you to define how your types should be printed, giving you full control over their string representation.

2. **Cleaner Code:** Instead of calling custom formatting functions every time you print a type, implementing `Stringer` ensures that formatting is built-in, making the code cleaner and easier to maintain.

3. **Integration with fmt Package:** The `fmt` package recognizes the `Stringer` interface, meaning it will always use the `String()` method when printing types that implement it.

### Another Example

Let's consider another custom type, `Coordinate`, and implement `Stringer` to customize its string output:

```go
package main

import (
    "fmt"
)

type Coordinate struct {
    X, Y int
}

func (c Coordinate) String() string {
    return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

func main() {
    point := Coordinate{X: 10, Y: 20}
    fmt.Println(point) // Output: (10, 20)
}
```
Here:

- The `Coordinate` type has `X` and `Y` fields.
- The `String()` method formats the coordinates as `"(X, Y)"`.
- When printing point, Go automatically uses the `String()` method to output `(10, 20)`.

### Final things
- The Stringer interface is part of the fmt package and contains just one method: `String() string`.
- Implementing `Stringer` allows you to define how your custom types are converted to strings when printed.
- Go's `fmt` package automatically uses the `String()` method for any type that implements `Stringer`.

## Practical Use Cases

Interfaces are particularly useful for:
- Writing **polymorphic functions** that work with multiple types.
- Abstracting behaviors, such as defining contracts for services or operations.
- Mocking in tests by passing in types that implement interfaces.

## Common Stanadard Interfaces

Go's standard library perovides many useful interfacesm such as:
- `fmt.Stringer`: Requires the `String()` method for custom string formatting.
- `io.Reader` and `io.Writer`: For reading and writing data streams.

## Key Points

1. **Interfaces define method sets** that types can implement.
2. Types **implicitly implement** interfaces if they have required methods.
3. Interfaces allow **polymorphism**, enabling functions to accept different types.
4. The **empty interfaces** (`interface{}`) can hold any value of type.
5. **Type assertions** are used to retrieve concrete values from an interface.
