# Go Structs

In Go, `struct` (short for structure) is a user-defined type that allows you to group together multiple fields, each with its own data type. Structs are often used to represent complex data types and model real-world entities by bundling related data into a single unit.

## Definig a Struct

To define a `struct`, you use the `type` keyword, followed by the name of the struct and the fields inside it:

```go
type Person struct {
      Name string
      Age int
      City string
      Grade float64
}
```
Here, `Person` is a struct with three fields: `Name` (a string), `Age` (a int), `City` (a string), and `Grade` (a float64).

## Creating and Initializing a Struct

Once you have defined a struct, you can create an instance of it (which is called a struct value) and assign values to its fields:

```go
// Here Ininitializing all fields
p := Person {
      Name: "Shincha",
      Age: 10,
      City: "Japan",
      Grade: 6.31
}
fmt.Println(p)    // Here the output will be: {Shincha 10 Japan 6.31}
```
In this example, we create a new `Person` named `p` with specific values for each field.

## Accessing Strcut Fields

You can access and modify the indivisual fields of `strut` using the dot (`.`) operator.

```go
fmt.Println(p.Name)  Output will be: Shincha

p.Age = 22    // Here updating the Age field 
fmt.Println(p.Age)
```

## Zero Values for Struct Fields

If you create a `struct` without initializing all fields, then Go will assgin **zero values** for any uninitialized fields (which means for empty string value will be `string`, `0` for `int`, `false` for `bool`, `0.0` for `float64` etc):

```go
p2 := Person{
      Name: "Heema"
      fmt.Println(p2)    // Output wil be: {Heema O } (Age City and Grade have zero values)
}
```

## Anonymous Structs

Go also allows you to create structs without formally defining a named type. These are called **anonymous structs** and are useful for quick, one-time struct creation:

```go
p := struct {
      Name string
      Age int
}{Name: "Khazam", Age: 14}

fmt.Println(p)    // Output will be: {Khazam 14}
```

## Struct with Methods

In Go, you can define methods (functions) on structs, allowing structs to have behaviors in addition to just storing data. A method is just function with a **receiver** (the struct it operates on):

```go
func (p Person) Greet() string {
      return "Hello fellas, My name is " + p.Name
}

p := Person {
      Name: "Boooe",
      Age: 16,
      City: "Japan",
      Grade: 8.01

fmt.Println(p.Greet())    // Output : Hello fellas, My name is Boooe
}
```

## Structs and Pointers

You can also use **pointers** to structs to modify the original struct values. Without pointers, struct values are copied whenn passed to functions or methods, and changes are made on the copy, not the original.

```go
func (p *Person) HaveBirthday() {
      p.Age++
}

p := Person {
      Name: "Mazaaav",
      Age: 25,
      City: "Japan"
      Grade: 7.2
}

p.HaveBirthday()
fmt.Println(p.Age)      // Output: 26 (original value is modified)
```
By using a pointer receiver (`*Person`), the `HaveBirthday` method updates the original `Age` field directky.

## Recap
- A **struct** is a way to group related data into a single type.
- You define a struct with the **type** keyword, followed by field names and types.
- Structs allow you to model complex entities by comibining multiple fields.
- You can access, modify, and use pointers to modify struct fields, and even associate methods with structs to give them behaviors.
