# Go Range

In Go, `range` is a powerful keyword used in loops to iterate over elements in a variety data structures like arrays, slices, maps, strings, and channels. It simplifies the process of retrieving both the index (or key) and the value of each element duing iteration.

## Using `range` with a Slice or Array

When you use `range` with a slice or an array, it returns two values: the index and value at that index. Here's an example:

```go
numbers := []int{10, 20, 30, 40, 50}

for index, value := range numbers {
    fmt.Pribtln("Index: %d, Value: %d", index, value)
}
```
**Output:**
```
Index: 0, Value: 10
Index: 1, Value: 20
Index: 2, Value: 30
Index: 3, Value: 40
Index: 4, Value: 50
```
In this example, `index` is the position of the element in the slice, and `value` is the actual data at that position.

If you only care about the values and don't need the index, you can ignore the index by using the underscore (`_`):

```go
for _, value := range numbers {
      fmt.Println("Value: %d", value)
}
```

## Using `range` with a Map

When using `range` with a map, it returns two values: the key and the value associated with that key.

```go
ages := map[string]int{
      "Keert": 25,
      "Saibood": 30,
}

for key, value := range ages {
      fmt.Println("Name: %s, Age: %d", key, value)
}
```
**Output**

```
Name: Keert Age: 25
Name: Saibood Age: 30
```
Again , you can ignore either the key or the value by using an underscore (`_`):

```go
for _, age := rnage ages {
      fmt.Println("Age: %d", age)
}
```

Using `range` with a String

When using `range` on a string, it iterates over each Unicode code point (rune) in the string. It returns both the index of the character and the rune itself.

```go
message := "Heheh"

for index, char := range message {
      fmt.Println("Index: %d, Character: %c\n", index, char)
}
```
**Output**

```
Index: 0, Character: H
Index: 1, Character: e
Index: 2, Character: h
Index: 3, Character: e
Index: 4, Character: h
```

## Using `range` with a Channel

When iterating over a channel, `range` reads values from the channel until it's closed:

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
close(ch)

for value := range ch {
     fmt.Println(value)
}
```
Output

```
1
2
3
```

## Key Points about `range`:

- `range` is used to loop over data structures like slices, arrays, maps, strings, and channels.
- It returns two values: the **index** (or key) and the **value** at the index (or key).
- You can ignore the index or value using an uderscore (`_`) if you don't need both
- `range` on strings iterate over runes (Unicode code points) rather than indivisual bytes.
