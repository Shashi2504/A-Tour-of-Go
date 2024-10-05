# Go Exported names

In Go, 𝗲𝘅𝗽𝗼𝗿𝘁𝗲𝗱 𝗻𝗮𝗺𝗲𝘀 are identifiers (like variables, constants, functions, types, etc.) that can be accessed from other packages. To make something "exported" and available to other packages, its name must start with an 𝘂𝗽𝗽𝗲𝗿𝗰𝗮𝘀𝗲 𝗹𝗲𝘁𝘁𝗲𝗿.

## Key Concepts of Exported Names

### Exported vs Unexported
  - 𝗘𝘅𝗽𝗼𝗿𝘁𝗲𝗱: Any identifier whose name starts with an uppercase letter (e.g., Print, MyFunction) is considered exported. Other packages can access it.

  - 𝗨𝗻𝗲𝘅𝗽𝗼𝗿𝘁𝗲𝗱: If an identifier starts with a lowercase letter (eg., internalVar, helperFunc), it is 𝘂𝗻𝗲𝘅𝗽𝗼𝗿𝘁𝗲𝗱, and only the pacakge where it is defined can access it.

Example of exported and unexported identifiers:
```go
// In the "mathutils" package
package main

var ExportVar = 10  // This can be accessed from other packages
var unexportedVar = 5  // This cannto be accessed from other packages

func Add(a int, b int) int {  // Exported function
      return a + b
}

func subtract(a int, b int) int {  // Unexported function
      return a - b
}
```

### Using Exported Names

When you import a package, you can only access its 𝗲𝘅𝗽𝗼𝗿𝘁𝗲𝗱 functions, variables, and types. For example, if we import the mathutils package, we can use `Add`, but not `subtract`, because `subtract` is unexported.

```go
package main

import (
      "fmt"
      "yourmodule/mathutils"
)

func main() {
      sum := mathutils.Add(3, 5)      // Allowed: Add is exported
      // diff := mathutils.subtract(5, 6)    // Error: subtract is unexported
      fmt.Println("Sum:", sum)
}
```

### Why Use Exported/Unexported Names?

  - 𝗦𝘂𝗺𝗺𝗮𝗿𝗶𝘇𝗮𝘁𝗶𝗼𝗻: Using unexported names allows you to hide implementation details and expose only what other packages need to use.
  - 𝗖𝗼𝗻𝘁𝗿𝗼𝗹 𝗔𝗰𝗰𝗲𝘀𝘀: By exporting only necessary functions or variables, you can prevent other packages from accidentally modifying internal logic.

## Recap

  - Start names with an uppercase letter to make them 𝗲𝘅𝗽𝗼𝗿𝘁𝗲𝗱 (accessible from other packages).
  - Start names with an lowercase letter to keep them 𝘂𝗻𝗲𝘅𝗽𝗼𝗿𝘁𝗲𝗱 (restricted to the same package).
