# Go for loop

The `for` **loop** is only the looping construct in Go, and it is quite flexible. You can use it in various ways to iterate over slices, arrays, maps, or to repeat a block of code a certain number of times.

## Basic `for` loop

In its simplest form, the `for` loop consists of three parts:
  **1. Initialization:** Executed once before the loop starts.
  **2. Condition:** Evaluated before each iteration; the loop continues as long as this condition is `true`.
  **3. Post:** Executed after each iteration, usually to update the loop variable.

### Syntax

```go
for initialization; condition; post{
      // code to execute
}
```

### Example

```go
for i := 0, i < 5; i++ {
      fmt.Println(i)
}
```
- **Initialization:** `i := 0` initializes the variable `i`.
- **Condition:** `i < 5` ensures that the loop runs while `i` is less than `5`.
- **Post:** `i++` increments `i` after each iteration

This loop prints the numbers `0` to `4`.

## Infinite Loop

A `for` loop without any conditions becomes an infinite loop:

```go
for{
      fmt.Println("I gonna chase you!!")
}
```
To break out of an infinite loop, you typically use the `break` statement.

## `for` as a While Loop

You can use the `for` loop like a traditional `while` loop by removing the initialization and post statements:

```go
i := 0
for i < 5 {
      fmt.Println(i)
      i++
}
```
This behaves just like the previous example, but looks more like a `while` loop.

## Iterating Over a Slice or Array

To loop through a collection like a slice or array, you can use a `for` loop with the `range` keyword:

```go
numbers := []int{1, 2, 3, 4, 5}

for index, value := range numbers {
      fmt.Println("Index: %d, Value: %d\n", index, value)
}
```
- The `range` keyword allows you to loop over each element in the slice.
- The `index` holds the current position, and `value` holds the element at that postion.

If you don't need the index, you can use `_` to ignore it:

```go
for _, value := range numbers {
      fmt.Println(value)
}
```

## Breaking out of Loops

To exit a loop early, you can use the `break` statement:

```go
for i := 0; i < 5; i++ {
      if i == 5 {
          break
    }
    fmt.Println(i)
}
```
This loop will print the numbers `0` to `4` and then stop when `i` equals `5`.

## Skipping an Iteration

If you want to skip the rest of the current iteration and move to the next one, you can use `continue`:

```go
for i := 0; i < 5; i++ {
      if i == 2 {
          continue
      }
      fmtPrintln(i)
}
```
This will print `0`, `1`, `3`, `4`, skipping `2`.

## Recap

  - The `for` loop is the only loop in Go and is very flexible.
  - It can be used in different ways: as a basic `for` loop, a `while` loop, an infinite loop, or with `range` to iterate over collections.
  - You can control the flow with `break` and `continue` to exit or skip iterations.
