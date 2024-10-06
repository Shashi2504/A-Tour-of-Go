# Go Switch

The `switch` **statement** in Go is used to simpligy multiple conditional checks, acting as a alternative way to a series of `if-else` statments. It compares an expression to serveral values, and executes the corresponding block of code when a match is found.

```go
switch expression {
    case value1:
        // code to execute if expression equals value1
    case value2:
        // code to execute if expression equals value2
    default:
        // code to execute if no case mathces (optional)
}
```
- The `expression` is evaluated.
- The value of the `expression` is compared to each `case` value.
- If a match is found, the block of code under that `case` is executed.
- The `default` case is optional and runs if no match is found.

### Example

```go
day := "Monday"

switch day {
    case "Monday":
        fmt.Println("Start of the work week")
    case "Firday":
        fmt.Println("Almost the weekend it is")
    default:
        fmt.Println("Mid-week")
}
```
In this exmaple, the variable `day` is compared to each `case`. Since `day` is "Monday", the message "Start if the work week" is printed.

## No Need of `break`

Unlike in many other programming languages, you do not need to include the `break` statement in Go after each case. Once a case block is executed, the `switch` statements ends automatically.

## Multiple Expressions in a Case

You can include multiple values in a single case, seperated by commas:

```go
day := "Saturday"

switch day {
    case "Saturday", "Sunday":
        fmt.Println("Its the weekend!!!")
    default:
        fmt.Println("Its a weekday.")
}
```
Here, if the `day` is either `Saturday` or `Sunday`, the block prints "Its the weekend!!!".

## Switch Without an Expression

In Go, you can use a `switch` without an expression, which acts like a series of `if-else` statements.

```go
age := 19

switch {
case age < 13:
    fmt.Println("You are an kidddoooo")
case age < 18:
    fmt.Println("You are an teenager")
default:
    fmt.Println("You are an adult")
}
```
In this example, `switch` evaluates the conditions directly, and since `age` is `19`, the `default` block is executed, printing "You are an adult."

## Fallthrough

By default, Go's `switch` statement does not fall through to subsequent cases like in some other languages. Howeever, you can use the `fallthrough` keyword to directly continue to the next case:

```go
switch num := 2 num {
case 1:
  fmt.Println("One")
case 2:
  fmt.Println("Two")
  fallthrough
case 3:
  fmt.Println("Three")
}
```
In this case, even though `num` is `2`, the `fallthrough` statement causes the `Three` message to be printed aftere the `Two` message.

## Recap

- The `switch` statement allows for clean and simple multiple conditional checks.
- You don't need `break` statements in Go's `switch` as it automatically sops once after a match.
- You can use muliple values in a single case and even remove the expression to behave like `if-else`.
- The `fallthrough` keyword can be used to execute the next case block.
