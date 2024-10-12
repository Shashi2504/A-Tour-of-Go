## Function Values

In Go, **function values** mean that functions themselves can be treated like a regular values. You can assign a function to a variable, pass it as an argument to another function, or return it from a function. This allows you to create more flexible and modular code, where functions can be manipulated and passed around just like any other data type.


### Assigning a Function to a Variable

You can store a function in a variable, and then call that function using the variable the variable name:

```go
func add(a, b int) int {
    return a + b
}

sum := add    // Assign the function to the variable 'sum'

result := sum(3, 4)    // Call the function via 'sum'
fmt.Println(result)    // Output: 7
```
In this example, the `add` function is assigned to the `sum` variable, and then we use `sum` to all the function and get the result.

### Passing a Function as an Argument

Functions can be passed as arguments to other functions, allowing you to create more flexible and dynamic behavior.

```go
func calculate(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func multiply(a, b int) int {
    return a * b
}

result := calculate(5, 6 multiply)    // Passing `multiply` function as an arugument
fmt.Println(result)    // Output: 30
```
Here, the `calculate` function takes two integers and a function as arguments. You can pass different functions (like `multiply`) to acheieve different operations.

### Return a Function from Another Function

You can also return a function from another function, which is usefuk for creating higher-order functions (functions that operate on other functions).

```go
func getMutliplier(factor int) func(int) int {
    return func(num int) int {
        return num * factor
    }
}

timesTwo := getMultiplier(2)
result := timesTwo(4)    // Calling the returned function
fmt.Println(result)    // Output: 8
```
In this example, `getMultiplier` returns a function that multiplies a number by the given factor. The `timesTwo` variable now holds a fucntion that multiplies its input by 2.

### Functions Values in Practice

Functions values are commonly used in Go for things like:
  - **Callbacks:** Passing functions to be executed later (e.g. in event-driven programs).
  - **Higher-order functions:** Functions that return other functions or take them as arguments (like the `calculate` and `getMultiplier` examples.)
  - **Custom logic:** Paasing different functions to be executed based on specifi logic or conditions.

**Example with Anonymous Functions**

You can also assign **anonymous functions** (functions without a name) to variables:

```go
double := func(x int) int {
    return x * 2
}

fmt.Println(double(5))    // Output: 10
```
This creates a function that doubles a number and stores it in the `double` variable, which you can call later.

### Key Points
- Functions in Go can be assigned to variables, passed as arguments, and returned from other functions.
- Function values add flexibility, allowing for modular and reusable code.
- Anonymous functions can be stored in variables and used just like named functions.
