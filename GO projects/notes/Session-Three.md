
---

# Golang Bootcamp - Session 3

**Duration:** 1.5 hours
**Topic:** Functions in Go

---

## 1. Topics Covered

* Function basics & syntax
* Parameters & return values
* Multiple return values
* Named return values
* Variadic functions
* Passing by value vs passing by reference (pointers)
* Functions as values & anonymous functions
* Defer statements

---

## 2. Detailed Notes

### Function Basics

A function in Go has:

* A **name**
* Optional **parameters**
* Optional **return values**
* A **body** with code

Basic example:

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello from a function!")
}

func main() {
    greet()
}
```

---

### Parameters & Return Values

* Parameters must have a type.
* Return type is specified after parameters.

```go
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 4)
    fmt.Println("Sum:", result)
}
```

---

### Multiple Return Values

Go can return multiple values.

```go
func divide(a, b int) (int, int) {
    quotient := a / b
    remainder := a % b
    return quotient, remainder
}

func main() {
    q, r := divide(10, 3)
    fmt.Println("Quotient:", q, "Remainder:", r)
}
```

---

### Named Return Values

You can name return variables in the function signature.

```go
func rectangle(width, height int) (area int, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return // automatically returns area, perimeter
}
```

---

### Variadic Functions

Functions can take an arbitrary number of arguments.

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

func main() {
    fmt.Println(sum(1, 2, 3, 4, 5))
}
```

---

### Passing by Value vs Reference

* By default, Go passes variables **by value**.
* Use **pointers** (`*T`) to pass by reference.

```go
func increment(x *int) {
    *x++
}

func main() {
    num := 10
    increment(&num)
    fmt.Println(num) // 11
}
```

---

### Functions as Values & Anonymous Functions

Functions can be assigned to variables.

```go
func main() {
    sayHello := func(name string) {
        fmt.Println("Hello,", name)
    }
    sayHello("Go")
}
```

---

### `defer` Statements

* **`defer`** postpones execution until the surrounding function returns.
* Often used for cleanup (e.g., closing files).

Example:

```go
func main() {
    fmt.Println("Start")
    defer fmt.Println("This will run at the end")
    fmt.Println("Middle")
}
```

Output:

```
Start
Middle
This will run at the end
```

---

## 3. Key Takeaways

* Functions are first-class citizens in Go.
* Multiple return values are common (e.g., `(value, error)`).
* Use variadic functions for flexible arguments.
* Go passes arguments by value unless using pointers.
* `defer` is great for cleanup tasks.

---

## 4. Exercises

1. Write a function `max(a, b int) int` that returns the greater of two numbers.
2. Create a `swap(a, b *int)` function that swaps the values of two integers using pointers.
3. Write a variadic function `average(numbers ...float64) float64` that calculates the average.
4. Use an anonymous function to print `"Hello from anonymous function!"`.

---

If you want, I can move to **Session 4: Arrays, Slices & Maps** next so your bootcamp README keeps growing.
Do you want me to continue with **Session 4** now?
