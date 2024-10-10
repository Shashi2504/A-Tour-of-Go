# Go Slices

In Go, **slices** are a more flexible and powerful way to work with sequence of elements compared to arrays. While arrays have a fixed size, slices allow you to dynamically change their size, amking them the preferred way to manage collection of elemeent in Go.

## Defining a Slice

A slice is essentially a reference to an underlying array. You can create a slice using a slice literal, like this:

```go
numbers := []int{1, 2, 3, 4, 5}    // A slice of integers
```

This creates a slice `numbers` containing the values `1, 2, 3, 4, 5`. Unlike arrays, slices do not require you to specify their size; Go automatically handles the size and capacity for you.

You can also create a slice from an existing array by "slicing" it.

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4]    // Creates a slice of arr with elements {20, 30, 40}
```
In the example above, `slice: arr[1:4]` creates a slice from the second element to the fourth element of `arr`. Slices use the syntax `arr[startIndexd:endIndex]` where `startIndex` is inclusive and `endIndex` is exclusive.

## Length and Capacity

A slice has both a **length** (the number of elements it currently holds) and a **capacity** (the meximum number of elemeents it can hold before it needs to resize the underlying array).

You can check the length and capacity of a slice using the `len()` and `cap()` functions:

```go
numbers := []int{1, 2, 3, 4, 5}
fmt.Println(len(numbers))    // Output: 5 (length)
fmt.Println(cap(numbers))    // Output: 5 (capacity)
```

If a slice is created from an array, the capacity depends on how much of the array is available from the starting point of the slice:

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4]    // Slice with elements {20, 30, 40}
fmt.Println(len(slice))    // Output: 3
fmt.Println(cap(slice))    // Output: 4 (baecause the underlying array has more elements from index 1 onwards)
```

## Modifying a Slice

You can change elements of a slice just like you would with an array:

```go
slice := []int{1, 2, 3}
slice[1] = 10    // Change the second element to 10
fmt.Println(slice)    // Output: {1, 20, 3}
```

## Appending to a Slice

Slices can be dynamically resized by appending new elements using the `append()` the function:

```go
numbers := []int{1, 2, 3}
numbers := append(numbers, 4, 5)    // Adds two new elements to the slice
fmt.Println(numbers)    // Output: [1, 2, 3, 4, 5]
```
If the capacity of the slice is not enought o accomodate the new elements. Go automatically resizes the underlying array and allocates more space:

## Copying a Slice

You can use the `copy()` function to copy elements from one slice to another:

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))    // Craetes a destination slice with the same length
copy(dst, src)    // Copy elements from src to dst
fmt.Println(dst)  // Output: [1, 2, 3]
```

## Slicing a Slice

You can create sub-slices from a slice, just like you can slice an array:

```go
numbers := []int{1, 2, 3, 4, 5}
subSlice := numbers[1:4]    // Creates a sub-slice {2,3, 4}
fmt.Println(subSlice)    // Output: [2, 3, 4]
```

## Zero Value of a Slice

The zero value of a slice is `nil`. A `nil` slice behaves like an empty slice but has a length and capacity of `0`:

```go
var s []int
fmt.Println(s)    // Output: []
fmt.Println(len(s))    // Output: 0
fmt.Println(cap(s))    // Output: 0
fmt.Println(s == nil)    // Output: true
```

## Key Points about Slices

- **Slices** are dynamically sized, unlike arrays.
- They reference an unerlying array and have both a length (number of elements) and a capactiy (space for elements).
- **Appending** to slices automatically resizes them if necessary.
- **Slicing** a slice creates a sub-slice without copying the underlying elements.
- `len` gives the length of the slice, and `cap` gives the capacity.
- A `nil` slice has no elements and zero capacity but is still valid.
