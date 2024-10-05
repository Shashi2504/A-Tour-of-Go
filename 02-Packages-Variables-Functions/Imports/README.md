# Go Imports

In Go, imports allows you to bring external code from other packages into your program so you can use it. This helps you organize your code into reusable components and leverage Go's standard library or third-party libraries.

## How Imports Work

### Importing Standard Library Packages

Go has a power standard library with many packages for tasks like input/output, networking, and string manipulation. You can import them directly into your program. For example, the `fmt` package used for formatted I/O:

```go
package main

import "fmt"

func main() {
      fmt.Println("Hey, Go mate")
}
```
Here, we import `fmt` to use its `Println function`.

### Importing Multiple Packages

If you need more than one package, you can import multiple packages by listing them inside paranthesis. This is called a 𝗳𝗮𝗰𝘁𝗼𝗿𝗲𝗱 𝗶𝗺𝗽𝗼𝗿𝘁:

```go
import (
      "fmt"
      "math"
)
```
Now, you can use functions from both `fmt` and `math`.

### Importing Custom or Third-Party Packages

You can also import your own custom packages or third-party packages (like from GitHub) using their paths. For example, if you have a package in your project or want to use as an external library, you would import it like this:

```go
import "github.com/user/package"
```
After importing, you can access it functions and types.

### Alias Imports

You can assign a shorter or more meaningful alias to a package name using the `import` statement. This can make your code cleaner:

```go
import fm "fmt"

func main() {
      fmt.Println("Hello with alias")
}
```
Here, `fm` is an alias for `fmt` package.

### Blank Imports

Sometimes, you need to import a package for its side effects (like running initialization code) but don't need to use any of its functions directly. In such cases, you can use a 𝗯𝗹𝗮𝗻𝗸 𝗶𝗺𝗽𝗼𝗿𝘁 by using an underscore (`_`):

```go
import _ "some/package"
```
This tells Go to inlude the package but not to assign it a name in your code.

### Import Path Rules

Go requires you to use the full import path of a package. For example, if you're importing a package from GitHub, you must specify the full URL path like this:

```go
import "github.com/gorilla/mux"
```
