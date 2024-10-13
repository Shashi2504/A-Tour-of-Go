# Go Methods

In Go, **methods** are functions with a special receiver argument. This receiver is a value or pointer to a value that allows you to associate the method with a specific type, giving you the ability to define behaviors (methods) on types.

## Defining Methods

To define method, you can use the same syntax as a function but add a **receiver** before the function name. The receiver specifies which type the method belongs to. Here's an example:

```go
type rectangle struct {
    width, height float64
}

?? Method to calculate the area of a rectangle

func (r rectangle) Area() float64 {
    return r.width * r.height
}
```
**In this example:**
- The method `Area` is defined with a receiver `(r rectangle)`. This means that `Area` is a method that operates on the `rectangle` type.
- The method can be called on any `rectangle` instance and calculates the area based on the width and height of that rectangle.

## Calling a Method

Once you define a method, you can it on instances of the type, like this:

```go
rect := rectangle(width: 10, height: 4)
fmt.Println(rect.Area())
```
Here, `rect.Area()` calls the `Area` method, and it returns the area of the rectangle, which is `40`.

## Pointer Receivers vs Value Receivers

Methods can have either **values receivers or pointer receivers**. The difference affects how the method interacts with the original value.
- **Value receivers:** The method operates on a copy of the value, so it cannot modify the original value.
- **Pointer receivers:** The method operates on a pointer to the value, so it can modify the original value.

**Example of a Pointer Receiver**

```go
func (r *rectangle) Scale(factor float64) {
    r.width *= factor
    r.height *= factor
}
```
In this method, `r *rectangle` is a pointer receiver, meaning the method can modify the `width` and `height` of the `rectangle` it is called on.

Here's how to call this method:

```go
rect := rectangle(width: 10, height: 4)
rect.Scale(2)
fmt.Println(rect.width, rect.height)    // Output: 20, 8
```
Since `Scale` is defined with a pointer receiver, it modifies the original `rect` object, scaling its width and height by factor of 2.

## When to Use Pointer vs Value Receivers
- Use a **value receiver** when the method doesn't need to modify the value and works with a copy of the value.
- Use a **pointer receiver** when the method needs to modify the original value or to avoid copying the large structures when calling methods.

- ## Methods on Any Type

- Methods can be defined on any type, not just `struct` types. For example, you can define methods on custom types like this:

    ```go
    type MyInt int

    func (m MyInt) Double() int {
    return int(m * 2)
    }

    num := MyInt(5)
    fmt.Println(num.Double())    // Output: 10
    ```
    In this case, `Double` is a method on the custom `MyInt` type, and it returns double the value of the integer.

## Key Points
- Methods allow you to define function that belong to specific types.
- Methods can be defined with either value or pointer receivers.
- Value receivers operates on copies, while pointer receivers allow you to modify the original value.
- You can define methods on any custom type, not just `struct` types.
