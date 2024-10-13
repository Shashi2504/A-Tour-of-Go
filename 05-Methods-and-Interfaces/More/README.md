# Go Errors

In Go, **errors, reader, and images** are important concepts that help manage different aspects of programming, especially when dealing with input/ouput and error handling. Let's get into each one:

## Errors

In Go, errors are represented by the built-in `error` type, which is an interface that includes a single method:

```go
type error interface {
    Error() string
}
```

**Creating and Handling Errors**

You can create custom errors using the `errors` package. Here's how to create and handle errors:
  1. Creating an Error: You can create a new error using `errors.New()` or by defining a custom type that implements the `error` interface.

     ```go
     import errors

     func doSmthg() error {
          return errors.New("an error occured")
     } 
     ```
  2. **Handling Errors:** When calling functions that retrun an error, you should check the error value.

     ```go
     err := doSmthg()
     if err @= nil {
          fmt.Println("Error:", err)
     }
     ```

**Example**

Check the below one you can get clear understanding of this:

```go
package main

import (
      "fmt"
      "errors"
)

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, error.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}
```

# Go Readers

In Go, **Readers** are used for input/output operations, especailly when dealing with streams of data. The `io.Reader` interaface is defined as:

```go
type Reader interface {
    Read(p []bye) (n int, err error)
}
```

**Reading Data**

The `Read` method reads upto `len(p)` bytes into `p` and returns the number of bytes read and any error encountered. There are several built-in types tha timplement the `Reader` interface, such as `os.File`, `bytes.Buffer`, and `strings.Reader`.

**Example of Using a Reader**

Here's an example of reading from a string using `strings.Reader`:

```go
package main

import (
      "fmt"
      "io"
      "strings"
)

func main() {
      s := "Hello, Go!!"
      reader := strings.NewReader(s)

      buf := make([]byte, 4)    // Buffer to hold read data
      for {
        n ,err := reader.Read(buf)
        if err == i.EOF {
            break
        }
        if err != nil {
        fmt.Println("Error:", err)
        }
        fmt.Println(string(buf[:n]))    // Print the read bytes
      }
}
```

# Go Images

Go provide the `image` package, which includes types and functions to work with images. You can manipulate images, decode them from various formats (like PNG, JPEG) and even create new images programatically.

**Working with Images**

1. **Importing the Image Package:** You usually import `image`, along with a specific image format package (like `image/png`).
2. **Decoding an Image:** You can read an iage from a file and decode it.

   ```go
   package main

   import (
      "fmt"
       "image"
       "image/png"
       "os"
   )

   func main() {
       file, err := os.Oepn("image.png")
       if err != nil {
          fmt.Println("Error:", err)
           return
       }
       defer file.Close()

       img, err := png.Decode(file)
       if err != nil {
           fmt.Println("Error:", err)
           return
       }

       fmt.Println("Image size:", img.Bounds().Size())
   }
   ```
   3. Creating a New Image: You can also create a new image from scratch using `image.NewRGBA`.
  
   ```go
   package main

   import (
        "image"
        "image/color"
        "image/png"
        "os"
   )

   func main() {
        img := image.NewRGBA(image.Rect(0, 0, 100, 100))

       // Set the pixels to red
       for y := 0, y < 100; y++ {
            for x := 0; x < 100; x++ {
                img.Set(x, y, color.RGBA{255, 0, 0, 255})    // Red
           }
       }

       // Save the image to a file
       file, _ := os.Create(red_image.png)
       defer file.Close()
       png.Encode(file, img)
   }
   ```

   ## Key Points

   - **Errors:** Use the `errors` type to handle and represent errors in your Go programs.
   - **Reader:** The `io.Reader` interface allows you to read data from various sources, making it easier to handle streams of input.
   - **Images:** The `image` package provides toos for working with images, allowing you to decode, manipulate, and create images in Go.
