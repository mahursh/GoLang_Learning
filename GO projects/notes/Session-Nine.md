

---

# Session 9 – Error Handling and Custom Errors in Go

## 📌 Overview

Error handling in Go is explicit — errors are values.
Instead of exceptions like in some languages, Go uses the `error` type to represent issues.

---

## 1️⃣ The `error` Type

* The `error` type is an **interface** in Go.
* Defined in the `builtin` package:

```go
type error interface {
    Error() string
}
```

* Any type that has an `Error()` method returning a string implements the `error` interface.

---

## 2️⃣ Returning Errors from Functions

* Go’s convention: **last return value is the error**.

```go
package main
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
    result, err := divide(4, 0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}
```

✅ **Key point:** Always check `err != nil`.

---

## 3️⃣ Using `errors.New()` vs `fmt.Errorf()`

* **`errors.New()`**: Simple static error messages.
* **`fmt.Errorf()`**: Formatted error messages with variables.

```go
import (
    "errors"
    "fmt"
)

err1 := errors.New("something went wrong")
err2 := fmt.Errorf("file %s not found", filename)
```

---

## 4️⃣ Custom Error Types

You can define your own error type for richer error information.

```go
package main
import (
    "fmt"
)

type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
    return &MyError{Code: 404, Message: "Resource not found"}
}

func main() {
    err := doSomething()
    if err != nil {
        fmt.Println(err)
    }
}
```

---

## 5️⃣ Wrapping and Unwrapping Errors

From **Go 1.13**, you can wrap errors using `%w` and unwrap using `errors.Is` or `errors.As`.

```go
package main
import (
    "errors"
    "fmt"
)

var ErrNotFound = errors.New("not found")

func findItem() error {
    return fmt.Errorf("database query failed: %w", ErrNotFound)
}

func main() {
    err := findItem()
    if errors.Is(err, ErrNotFound) {
        fmt.Println("Item not found in DB")
    }
}
```

---

## 6️⃣ Panic and Recover (Advanced)

* **`panic()`**: Stops normal execution.
* **`recover()`**: Regains control of a panicking goroutine.
* Use sparingly — Go prefers explicit error returns.

```go
func mayPanic() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()
    panic("Something bad happened!")
}
```

---

## 🔑 Key Takeaways

* In Go, **errors are values**, not exceptions.
* Always handle `err` explicitly — ignoring errors is a bad practice.
* Create **custom error types** for more context.
* Use **error wrapping** for better debugging.
* Avoid panic except for truly unrecoverable situations.

---

