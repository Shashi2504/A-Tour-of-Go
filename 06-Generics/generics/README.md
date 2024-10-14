# Go Generics

**Generics** in Go allows you to wirte functions, types, and data structures that can work with any data type, making your code more flexible and reusable. Introduced in Go 1.18, generics let you define a function or type that can handle values of different types without writing seperate code for each type.

## Why Generics?

Without generics if you want to write a function that works with different types (like `int`, `string`, `float64`), then you have to write a version of the function for each type. With generics, you have to write the function once, and it can work with any type.

## Syntax of Generics

To use generics in Go, you define a type parameter inside square brackets `[]` after the function name.
