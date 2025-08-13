

---

# Golang Bootcamp - Session 6

**Duration:** 1.5 hours
**Topic:** Interfaces in Go

---

## 1. Topics Covered

* What is an interface?
* Declaring and implementing interfaces
* Implicit interface satisfaction in Go
* Empty interface (`interface{}`)
* Type assertions
* Type switches
* Embedding interfaces

---

## 2. Detailed Notes

### What is an Interface?

* An **interface** is a type that defines a set of method signatures.
* Any type that implements those methods **automatically satisfies** the interface (no explicit `implements` keyword like Java).

---

### Declaring an Interface

```go
type Speaker interface {
    Speak() string
}
```

---

### Implementing an Interface

```go
type Person struct {
    Name string
}

func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

func main() {
    var s Speaker
    s = Person{Name: "Alice"}
    fmt.Println(s.Speak())
}
```

---

### Implicit Satisfaction

* You **don’t need** to explicitly declare that a type implements an interface.
* If a type has all the required methods, it satisfies the interface automatically.

---

### Empty Interface (`interface{}`)

* Special interface with **zero methods**.
* Can hold values of **any type**.
* Useful for generic containers or when type is not known at compile time.

```go
func printAnything(value interface{}) {
    fmt.Println(value)
}

func main() {
    printAnything(42)
    printAnything("Go")
    printAnything([]int{1, 2, 3})
}
```

---

### Type Assertions

* Extract the underlying value from an interface variable.

```go
var i interface{} = "Hello"
str, ok := i.(string)
if ok {
    fmt.Println("String value:", str)
} else {
    fmt.Println("Not a string")
}
```

---

### Type Switch

* Switch on the **type** stored in an interface.

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    describe(10)
    describe("Go")
    describe(3.14)
}
```

---

### Embedding Interfaces

* One interface can embed another, combining method sets.

```go
type Reader interface {
    Read() string
}

type Writer interface {
    Write(data string)
}

type ReadWriter interface {
    Reader
    Writer
}
```

---

## 3. Key Takeaways

* Interfaces define **behavior**, not data.
* Go uses **implicit implementation** — no explicit keywords.
* The empty interface can hold any type.
* Use type assertions and type switches to work with interface values.
* Interfaces can be embedded to create more complex behaviors.

---

## 4. Exercises

1. Create a `Shape` interface with a method `Area() float64`.
   Implement it for `Circle` and `Rectangle` structs.
2. Write a function that takes an `interface{}` and:

   * Prints the type using a type switch.
   * Prints the value.
3. Create an embedded interface `File` that combines `Reader` and `Writer`,
   then implement it for a `Document` struct.

---


