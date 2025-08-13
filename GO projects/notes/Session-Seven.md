

---

# Golang Bootcamp - Session 7

**Duration:** 1.5 hours
**Topic:** Error Handling in Go

---

## 1. Topics Covered

* The `error` type in Go
* Returning errors from functions
* Creating custom errors
* Using `errors.New` and `fmt.Errorf`
* The `panic` function
* The `recover` function
* Best practices for error handling

---

## 2. Detailed Notes

### The `error` Type

* In Go, `error` is a built-in interface:

  ```go
  type error interface {
      Error() string
  }
  ```
* Any type that implements `Error()` satisfies `error`.

---

### Returning Errors from Functions

* Common Go pattern: return `(value, error)`
* Caller checks if error is not `nil`.

```go
import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

---

### Creating Custom Errors

* Using `errors.New`:

  ```go
  err := errors.New("something went wrong")
  ```
* Using `fmt.Errorf` for formatted errors:

  ```go
  err := fmt.Errorf("file %s not found", "data.txt")
  ```

---

### Defining Your Own Error Type

```go
type MyError struct {
    Code int
    Msg  string
}

func (e MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}

func doSomething() error {
    return MyError{Code: 404, Msg: "Resource not found"}
}

func main() {
    err := doSomething()
    fmt.Println(err)
}
```

---

### `panic` Function

* Used for unrecoverable errors.
* Stops normal execution and begins unwinding the stack.

```go
func main() {
    panic("something went very wrong")
}
```

⚠ **Best practice**: Use `panic` only for truly unexpected conditions, not normal errors.

---

### `recover` Function

* Used inside a `defer` to regain control after a panic.

```go
func safeDivision(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println(a / b)
}

func main() {
    safeDivision(10, 0)
    fmt.Println("Program continues after panic recovery")
}
```

---

## 3. Key Takeaways

* Errors in Go are values — always check them before using results.
* Use `errors.New` or `fmt.Errorf` for standard error messages.
* Create custom error types when you need more information.
* Reserve `panic` for unrecoverable states.
* Use `recover` only when you must handle a panic and continue execution.

---

## 4. Exercises

1. Write a function `readFile(filename string) (string, error)` that returns the file’s contents or an error if the file doesn’t exist.
2. Implement a `LoginError` struct with `Username` and `Reason`, and make it satisfy the `error` interface.
3. Write a function that triggers a panic, and recover from it to print `"Recovered and safe"`.

---


