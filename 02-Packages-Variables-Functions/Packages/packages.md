# Go Packages

In Go, a package is a way to organize and summarize code. It allows you to group related functions, types, and variables together, making your code easier to manage. Every Go program is made up of packages, and even a simple Go program uses the built-in main package.

## Key Points about Packages in Go

### Creating a Package

To create a package, you need to create a directory for it. Inside that directory, you create a Go file that begins with the `package` keyword, followed by the name of the package. For example:

```go
// File name: mathutils.go
package mathutils

func Add(a int, b int) int {
    return a + b
}
```
This creates a package named mathutils with a function `Add`.

### Using a Package

To use package in another Go file, you need to import it using the `import` statement. Here's how you can use the `mathutils` package:

```go
package main
import (
    "yourmodule/mathutils"    //Replace 'yourmodule' with your module name or use mathutils itself
    "fmt"
)

func main() {
    sum := mathutils.Add(3, 5)    //If you have created yourmodule then replace it with mathutils
    fmt.Println("Sum:", sum)    //Output: Sum: 8
}
```

### Standard Library Packages

Go comes with a rich standard library that provides many useful packages. For example, `fmt` for formatted I/O, `math` for mathematical functions, and `net/http` for HTTP server and client functionalities.

### Organizing Code

By organizing the code into packages, you improve readability and maintainability. Packages can also be reused across different projects.

### Module System

Go uses modules to manage dependencies, which allows you to work with packages from another projects or libraries easily. You can create a module using `go mod init` and then use `go get` to add dependencies.
