# Go Arrays

InGo, an **array** is a fixed-size sequence of elements, all of the same type. Once an array is defined, its size cannot be changed. Arrays are useful when you know exactly how may items you need to store, but they are less commonly used than slices, which offer more flexibility.

## Defining an Array

To define a array in Go, you specify its size (like number of elements) and type, like this:

```go
var arr [5]int  // An array of 5 integers
```
This creates an array named `arr` with 5 elements, all initialized to zero since the default value for `int` is `0`.

You can also initailize the array with values when you declare it:

```go
arr := [5]int{10, 20, 30, 40, 50}    // Array with values
```
This creates an array with the values `10, 20, 30, 40, 50`. If you remove some values, the remaining elements will be intialized to the zero value of their type:

```go
arr := [int]{1, 2}    // Remaining elements are zero (e.g., {1, 2, 0, 0, 0})
```

## Accessing Array Elemets

You can access indivisual elements of an array using thei index. Array indexes start at 0, so the first element is at index `0` and the last elements is at `size-1`.

```go
arr := [3]int{10, 20, 30}
fmt.Println(arr[0])    // Output: 10
fmt.Println(arr[2])    // Output: 30
```

You can also update the elementsby accessing through their index:

```go
arr [1] = 25    // Saying that update second element which 20 to 25.
```

## Iteratiung Over an Array

You can use a `for` loop to iterate over all elements of an array:

```go
arr := [3]int{10, 20, 30}

for i := 0, i < len(arr); i++ {
      fmt.Println(arr[i])
}
```

Or, more commonly, you can use a `for range` loop, which iterates over each element:

```go
for index, value := range arr {
      fmt.Println("Index: %d, Value: %d", index, value)
}
```

## Array Length

The length of an array is fixed and can be obtained using the built-inlen() function:

```go
arr := [5]int{10, 20, 30, 40, 50}
fmt.Println(len(arr))    // Output will be: 5
```

## Multidimensional Arrays

You can define mulitdimensional arrays, such as a 2D array (which is an array of arrays):

```go
var matrix [2][3]int    // A 2x3 matrix (2 rows, 3 columns)
matrix := [2][3]int {
      {1, 2, 3},
      {4, 5, 6},
}

fmt.Println(matrix[1][2])    // Output: 6 (Element at row 2, column 3)
```

## Limitations of Arrays

Arrays in Go are not resizable. If you need a dynamic, resizable collection, Go provides **slices**, which are more flexible and powerful. Slices are built on top of arrays but allow you to change their size, which is why slices ar more commonly used in Go.

## Recap
- An array is aa fixed-size, ordered collection of elements of same type.
- You can declare, initialize, and access arrays using indesxes.
- Arrays are less commonly used than slices in Go due to their fixed size.
- You can get the length of an array with `len()` and iterate over its elements using a loop.
