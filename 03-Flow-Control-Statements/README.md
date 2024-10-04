# Flow Control Statements

Welcome to the **Flow Control Statements** section of the repository! This section covers how to control the flow of execution in Go programs using loops, conditionals, and switches. Mastering these concepts will help you write logic that can make decisions and iterate over code based on conditions.

## What You'll Learn

- How to use `if` and `else` statements to control program flow based on conditions.
- Writing different types of loops in Go, including `for` as the only loop construct.
- How to use the `switch` statement for multi-way branching.
- Managing deferred function calls with `defer`.

## Code Files Overview

| File Name               | Description                                                                 |
|-------------------------|-----------------------------------------------------------------------------|
| `for_as_while.go`        | Demonstrates how the `for` loop in Go can act as a `while` loop.              |
| `if_with_short_ste.go`   | Shows `if` statements with short variable declarations.                      |
| `loop_and_func.go`       | Explains how to combine loops with functions.                                |
| `switch_cases.go`        | Demonstrates the usage of `switch` cases in Go.                              |
| `defer_func.go`          | Shows how to defer function calls until the surrounding function returns.    |
| `for_conti.go`           | Demonstrates using the `continue` statement within loops.                    |
| `ifandelse.go`           | Introduces basic `if` and `else` constructs.                                |
| `runtime.go`             | Shows runtime control flow.                                                  |
| `switch_eval.go`         | Evaluates cases within a `switch` statement dynamically.                     |
| `stacking_defer.go`      | Demonstrates how `defer` can be stacked and executed in LIFO order.          |
| `infinite_loop.go`       | Shows how to create an infinite loop using `for`.                            |
| `switch_no_cond.go`      | Demonstrates a `switch` without an explicit condition.                       |

## Key Concepts

1. **If Statements**:
   - `if` statements are used to execute code conditionally. You can even declare variables inside the condition!
   - Learn more in [`if_with_short_ste.go`](./if_with_short_ste.go) and [`ifandelse.go`](./ifandelse.go).

2. **For Loops**:
   - The `for` loop is the only looping construct in Go, but it's incredibly versatile. It can behave like a traditional `for` loop, a `while` loop, or even create infinite loops.
   - See examples in [`for_as_while.go`](./for_as_while.go) and [`for.go`](./for.go).

3. **Switch Statements**:
   - Go’s `switch` statement is simpler and more powerful than many other languages' `switch` constructs. It allows you to compare values, conditions, and even types.
   - Learn more in [`switch_cases.go`](./switch_cases.go) and [`switch_no_cond.go`](./switch_no_cond.go).

4. **Defer**:
   - The `defer` statement is used to postpone the execution of a function until the surrounding function returns, which is useful for resource management like closing files or cleaning up.
   - Check out how to stack deferred calls in [`defer_func.go`](./defer_func.go) and [`stacking_defer.go`](./stacking_defer.go).

5. **Error Handling with Control Statements**:
   - Go provides clean error handling using control flow statements, which you can see in various examples, especially when combined with `if` statements.

## How to Run the Code

1. **Navigate to the folder**:
   ```bash
   cd 03-Flow-Control-Statements
