
---

# Golang Bootcamp - Session 2

**Duration:** 1.5 hours
**Topic:** Control Flow in Go (if, else, switch, loops)

---

## 1. Topics Covered

* `if` and `else` statements
* Comparison & logical operators
* `switch` statements (including expressionless switch)
* `for` loops (traditional, while-style, infinite)
* Loop control (`break`, `continue`)
* `range` for iterating over collections

---

## 2. Detailed Notes

### `if` and `else` Statements

* Basic form:

  ```go
  if condition {
      // do something
  } else if anotherCondition {
      // do something else
  } else {
      // default case
  }
  ```

* Example:

  ```go
  package main

  import "fmt"

  func main() {
      age := 20

      if age >= 18 {
          fmt.Println("You are an adult.")
      } else {
          fmt.Println("You are underage.")
      }
  }
  ```

* You can declare variables inside an `if`:

  ```go
  if x := 10; x > 5 {
      fmt.Println("x is greater than 5")
  }
  ```

---

### Comparison & Logical Operators

| Operator | Meaning          |
| -------- | ---------------- |
| `==`     | Equal to         |
| `!=`     | Not equal to     |
| `<`      | Less than        |
| `>`      | Greater than     |
| `<=`     | Less or equal    |
| `>=`     | Greater or equal |
| `&&`     | Logical AND      |
| `\|\|`   | Logical OR       |
| `!`      | Logical NOT      |

---

### `switch` Statements

* Example with values:

  ```go
  package main

  import "fmt"

  func main() {
      day := "Tuesday"

      switch day {
      case "Monday":
          fmt.Println("Start of the week.")
      case "Friday":
          fmt.Println("Almost weekend!")
      default:
          fmt.Println("Midweek day.")
      }
  }
  ```

* Example without expression (acts like `if`/`else` chain):

  ```go
  switch {
  case 5 > 3:
      fmt.Println("5 is greater than 3")
  case 2 > 1:
      fmt.Println("2 is greater than 1")
  }
  ```

* You can group multiple cases:

  ```go
  switch day {
  case "Saturday", "Sunday":
      fmt.Println("Weekend!")
  }
  ```

---

### `for` Loops

Go has only one loop keyword: `for`.

1. **Traditional loop**

   ```go
   for i := 0; i < 5; i++ {
       fmt.Println(i)
   }
   ```

2. **While-style loop**

   ```go
   i := 0
   for i < 5 {
       fmt.Println(i)
       i++
   }
   ```

3. **Infinite loop**

   ```go
   for {
       fmt.Println("Loop forever!")
       break // without this, it runs forever
   }
   ```

---

### Loop Control

* **`break`** → exit the loop early
* **`continue`** → skip to next iteration

Example:

```go
for i := 1; i <= 5; i++ {
    if i == 3 {
        continue // skip 3
    }
    fmt.Println(i)
}
```

---

### `range` for Iteration

`range` lets you iterate over slices, arrays, maps, strings, channels.

Example with slice:

```go
numbers := []int{10, 20, 30}
for index, value := range numbers {
    fmt.Println(index, value)
}
```

Example with string:

```go
for i, char := range "Go" {
    fmt.Printf("%d: %c\n", i, char)
}
```

---

## 3. Key Takeaways

* Go uses **`if`**, **`switch`**, and **`for`** for all control flow.
* `switch` can be used without an expression for flexible condition checks.
* Go’s `for` loop can behave like traditional, while-style, or infinite loops.
* `range` is a clean way to loop over collections.
* `break` and `continue` help control loop execution.

---

## 4. Exercises

1. Write a program that checks if a number is even or odd.
2. Write a program that:

   * Reads a day of the week from the user.
   * Prints if it’s a weekday or weekend using `switch`.
3. Print all numbers from 1 to 20, skipping multiples of 3 using `continue`.
4. Iterate over a string using `range` and print each character on a new line.

---

