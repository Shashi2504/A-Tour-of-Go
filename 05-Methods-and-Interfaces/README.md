# Methods and Interfaces

Welcome to the **Methods and Interfaces** section! In this part, you'll learn how Go implements object-oriented programming concepts through **methods** and **interfaces**. Unlike other languages that use classes, Go uses these features to achieve similar functionality, offering a lightweight yet powerful way to organize your code.

## What You'll Learn

- How to define and use **methods** in Go.
- Understanding **interfaces** and how Go uses them for polymorphism.
- The concept of **interface values** and how they work.
- How to implement **error handling** using custom error types.
- Practical exercises to deepen your understanding, such as building a **stringer** and an **image generator**.

## Code Files Overview

| File Name                  | Description                                                              |
|----------------------------|--------------------------------------------------------------------------|
| `method.go`                 | Introduction to defining and using methods in Go.                        |
| `diff_func_meth.go`         | The difference between functions and methods.                            |
| `interface.go`              | Learn how to define and implement interfaces.                            |
| `interface_values.go`       | Explore how interface values work and how they store types and values.   |
| `stringer.go`               | Implement the `Stringer` interface, one of the most common interfaces.   |
| `exercise-stringer.go`      | A practice exercise to implement your own `Stringer`.                    |
| `exercise-image.go`         | Build a simple image generator using interfaces and methods.             |
| `pointers_and_func.go`      | Explore the relationship between pointers and methods.                   |
| `type_assertions.go`        | Learn how to perform type assertions with interfaces.                    |
| `type_switch.go`            | Implementing type switches to handle different types.                    |
| `errors.go`                 | Introduction to error handling and custom error types in Go.             |
| `exercise-errors.go`        | Exercise to build custom error types and handle errors in your programs. |

## Key Concepts

1. **Methods**:
   - Methods in Go are functions with a special receiver argument. They can be attached to any type (not just structs) to provide behavior to your types.
   - Learn how methods work in [`method.go`](./method.go) and [`diff_func_meth.go`](./diff_func_meth.go).

2. **Interfaces**:
   - Interfaces define a set of method signatures and are implemented implicitly in Go. This allows for flexible and decoupled code.
   - Check out how to define and use interfaces in [`interface.go`](./interface.go).

3. **Interface Values**:
   - Interface values in Go contain both the value of the specific type that implements the interface and its type information. This combination enables dynamic dispatch.
   - Learn about interface values in [`interface_values.go`](./interface_values.go).

4. **Type Assertions and Type Switches**:
   - Type assertions are used to extract the underlying concrete value from an interface, while type switches allow you to handle multiple types in a clean way.
   - Examples of type assertions and switches are available in [`type_assertions.go`](./type_assertions.go) and [`type_switch.go`](./type_switch.go).

5. **Custom Error Handling**:
   - Go’s error handling model revolves around explicit error checking. In this section, you’ll learn how to implement custom error types and handle them effectively.
   - Dive into error handling in [`errors.go`](./errors.go) and the exercise in [`exercise-errors.go`](./exercise-errors.go).

6. **Stringer Interface**:
   - The `Stringer` interface in Go is one of the most commonly implemented interfaces. It's used to define how a type should be printed as a string.
   - Practice implementing the `Stringer` interface in [`stringer.go`](./stringer.go) and [`exercise-stringer.go`](./exercise-stringer.go).

7. **Practical Exercises**:
   - Build a simple image generator using interfaces and methods in [`exercise-image.go`](./exercise-image.go).
   - Practice error handling by building custom error types in [`exercise-errors.go`](./exercise-errors.go).

## How to Run the Code

1. **Navigate to the folder**:
   ```bash
   cd 05-Methods-and-Interfaces

2. **Run any of the Go files**:
   ```bash
   go run <filename>.go

3. **Try the practical exercises**:
   - Run the stringer implementation with:
     ```bash
     go run exercise-stringer.go
   - Build and test the image generator:
     ```bash
     go run exercise-image.go
