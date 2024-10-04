# Packages, Variables, and Functions

Welcome to the **Packages, Variables, and Functions** section of the repository! This section will introduce you to the building blocks of Go programs. You'll learn how Go organizes code into packages, how to work with variables, and how to define and use functions effectively.

## What You'll Learn

- How Go uses packages for organizing code.
- Declaring and initializing variables in Go.
- The importance of functions and how to write them.
- Key concepts like named returns, constants, and type inference.

## Code Files Overview

| File Name            | Description                                                                 |
|----------------------|-----------------------------------------------------------------------------|
| `packages.go`       | Explains how packages work in Go.                                              |
| `imports.go`       | Shows how to import packages in Go.                                             |
| `export_names.go`         | Explains exported names and how to use them across packages.             |
| `multi_resu.go`         | Shows how to handle multiple return values from functions.                 |
| `named_return.go`     | Shows how to use named return values in functions.                           |
| `variables.go`        | Explains variable declarations and their scope.                              |
| `var_init.go`          |  Demonstrates variable declaration and initialization.                      |
| `short_var_dec.go`    | Demonstrates short variable declaration using `:=`.                          |
| `basic_types.go`      | Introduces Go's basic data types (int, float, string, etc.).                 |
| `zero_values.go`      | Shows how variables are automatically initialized to zero values in Go.      |
| `conversion.go`     | Demonstrates type conversion in Go.                                            |
| `type_inference.go`   | Explains how Go performs type inference during variable initialization.      |
| `contants.go`         | Introduces the concept of constants in Go.                                   |
| `numeric_constants.go`| Demonstrates numeric constants and their usage in Go.                        |

## Key Concepts

1. **Packages:**
   - Go programs are organized into **packages**, which are collections of related Go files.
   - Use the `import` keyword to bring in functionality from other packages.
   - Learn more in [`packages.go`](./packages.go) and [`imports.go`](./imports.go).

2. **Variables:**
   - Variables in Go are explicitly declared, but Go provides several convenient ways to do this.
   - Learn more about variable declarations in [`variables.go`](./variables.go), [`var_init.go`](./var_init.go), and [`short_var_dec.go`](./short_var_dec.go).

3. **Functions:**
   - Functions are essential to structuring Go programs. You’ll explore:
     - Basic function definitions.
     - Multiple return values ([`multi_resu.go`](./multi_resu.go)).
     - Named return values ([`named_return.go`](./named_return.go)).
     - Closures (introduced later).
   
4. **Type Inference:**
   - Go can automatically infer types based on the right-hand side of the variable declaration.
   - See examples in [`type_inference.go`](./type_inference.go).

5. **Constants and Zero Values:**
   - Constants are variables that cannot be modified after their declaration.
   - Go also provides zero values for uninitialized variables, covered in [`contants.go`](./contants.go) and [`zero_values.go`](./zero_values.go).

## How to Run the Code

To run the Go programs in this folder, follow these steps:

1. **Navigate to the folder:**
   ```bash
   cd 02-Packages-Variables-Functions
