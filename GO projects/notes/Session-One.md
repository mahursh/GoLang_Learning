
---

# Golang Bootcamp - Session 1

**Duration:** 1.5 hours
**Topic:** Introduction to Go & Basic Syntax

---

## 1. Topics Covered

* Why Go?
* Go Installation & Setup
* Go workspace & `GOPATH` vs Go Modules
* Basic program structure (`package main`, `func main`)
* Printing output with `fmt.Println`
* Variables & Constants
* Basic Data Types (`string`, `int`, `float`, `bool`)
* Simple I/O

---

## 2. Detailed Notes

### Why Go?

* Created by Google in 2009 for efficiency, scalability, and concurrency.
* Compiles to machine code → very fast.
* Strongly typed, garbage collected, built-in concurrency.
* Popular for backend services, networking, DevOps tools.

### Installation & Setup

1. Download from: [https://go.dev/dl/](https://go.dev/dl/)
2. Verify install:

   ```bash
   go version
   ```
3. Create a workspace folder and enable modules:

   ```bash
   mkdir go-bootcamp
   cd go-bootcamp
   go mod init example.com/bootcamp
   ```

### Go Program Structure

* Every executable Go program must have:

  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, Go!")
  }
  ```
* `package main` → entry point
* `import` → brings in standard libraries or external packages
* `main()` → starting function

### Variables

* Declared with `var`, or short declaration `:=`

  ```go
  var name string = "Alice"
  age := 25 // type inferred
  ```
* Constants:

  ```go
  const Pi = 3.14159
  ```

### Basic Data Types

* Integers: `int`, `int8`, `int64`
* Floating point: `float32`, `float64`
* Strings: `"Hello"`
* Boolean: `true` / `false`

### Simple I/O

```go
package main

import "fmt"

func main() {
    var name string
    fmt.Print("Enter your name: ")
    fmt.Scanln(&name)
    fmt.Println("Hello,", name)
}
```

---

## 3. Key Takeaways

* Go is simple, fast, and great for scalable apps.
* The `main` package & `main` function are required for executables.
* Use `var` or `:=` for variable declaration.
* Go has only a few basic types, but they are strongly typed.
* Input & output use the `fmt` package.

---

## 4. Exercises

1. Install Go and set up a workspace.
2. Write a program that:

   * Asks for your name and age.
   * Prints `"Hello, <name>, you are <age> years old!"`.
3. Create a constant `pi` and print the area of a circle for radius `5`.

---

