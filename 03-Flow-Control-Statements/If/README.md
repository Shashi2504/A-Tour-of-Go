# Go If statements

The `if` **statement** in Go used to run a block of code conditionally, based on whether a specified condition evaluates to `true` or `false`. It is similar to `if` statements in other programming languages , but Go has a specifc features that make it unique.

## Basic Syntax

```go
    if condition {
      // code to execute when if condition is true
}
```
- The condition is evaluated
- If the condition is `true`, the block of the code inside the curly braces `{}` is executed.
- If the condition is `false`, the block is skipped.

**Example**

```go
age := 20

if age >= 18 {
      fmt.Println("You are an adult.")
}
```
In this example, since age is `20`, which is greater than or equal to `18`, the message "You are an adult." will be printed.

## `else` Clause

You can add an `else` clause to run a block of code if the condition is `false`:

```go
age := 16

if age >= 18 {
      fmt.Println("You are an adult.")
} else {
      fmt.Println("You are a minor kiddoo")
}
```
In this case, since `age` is `16`, the condtion `age >= 18` is `false`, so the else block is executed, printing "You are a minor kiddoo".

## `else if` Clause

To check multiple conditions, you can use `else if`:

```go
age := 13

if age >= 18 {
      fmt.Println("You are an adult.")
} else if >= 13 {
      fmt.Println("You are an teenager.")
} else {
      fmt.Println("You are a kiddooo")
}
```
Here, since age is 13, the first condition age >= 18 is false, but the second condition age >= 13 is true, so "You are a teenager." is printed.

## Short Statements in `if`

Go allows you to include a short statement before the condition, often used for variable initialization. This short statement is excuted before the condition is evaluated.

**Example**

```go
if num := 10; num > 5 {
      fmt.Println("Number is greter than 5")
}
```
- The variable `num` is declared and initialized to `10`.
- The condition `num > 5` is checked, and since it is `true`, the block is executed, printing "Number is greater than 5".

This `num` variable is only available only inside the `if` block and does not exist outside of it.

## Recap

- The `if` statement runs code based on whethere the condition is `true` or `false`.
- You can use `else` for alternate code when the condition is `false`.
- `else if` allows for multiple conditions to be checked.
- Go supports a short statement in `if` for initializing variables jsut before the checking the condition.
