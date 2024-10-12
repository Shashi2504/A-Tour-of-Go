# Go Maps

In Go, a **maps** is a build-in data type that stores key-value pairs. It is like a dictionary in other languages. The keys in a map are unique, and each key maps to a specific value. Maps provide fast lookups, adds, and deletes, amking them useful for scenarious where you need to associate one of values (keys) with another (values).

## Declaring a Map

To create a map, you need to specify the key type and the value type. You can declare a map using the `make` function or a map literal.

  1. Using `make` function:
```go
ages := make(map[string]int)
```
This creates a map named `ages` where the keys are of type `string` and the values are of type `int`.

  2. Using a `map` literal:
```go
ages := map[string]int{
    "Memiii": 25,
    "Beeru": 30,
}
```
This is another way of creating and intializing a map with a key-value pairs at the same time.

## Adding and Accessing Elements

You can add or update elements in a map by using the key inside square brackets `[]` and assigning a value:

```go
ages["Suiii"] = 34,    // Adding a new key-value pair
ages[""Beeru] = 44,    // Updating an existing key
```

To access the value associated with a key, you use the same syntax:

```go
age := ages["Beeru"]
fmt.Println(age)    // Output: 44
```

## Checking if a Key exists

To check if a key exists in a map, you can use the **comma ok idiom**. This will return two values: the value associated with the key, and a boolean indicatimg whether the key exists.

```go
age, exists := ages["Beeru"]
if exists {
    fmt.Println("Beeru ages is:", age)
} else {
    fmt.Println("Beeru age is unknown")
}
```

## Deleting Elements

To remove key-value pair from a map, you can use the `delete` function:

```go
delete(ages, "Beeru")
```
This removes the key `"Beeru"` and its associated value from the map.

## Iterating Over a Map

You can use the `range` keyword to loop through all key-values pairs in a map:

```go
foe name, age := range ages {
    fmt.Println(name, "is", age, "years old")
}
```
**Output**
```
Beeru age is 44 years old
Memiii age is 25 years old
```

## Zero Value for a Map

Teh zero value of map is `nil`, meaning it hasn't been initialized yet. If you try to read from a nil map, it will behave as if it's empty, but adding elements to a nil map will cause the run-time error.

## Key Points about Maps

- Maps sotre key-value pairs, and the kyes must be unique.
- You can add, accessm update, and delete elements using the key.
- You can check if a key exists using the comma ok idiom.
- Iterating over maps with `range` gives you both the key and value.
- A map's zero value is `nil`, and you must initialize it before adding elements.
