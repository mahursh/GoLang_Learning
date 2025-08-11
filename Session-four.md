
---

# Golang Bootcamp - Session 4

**Duration:** 1.5 hours
**Topic:** Arrays, Slices & Maps

---

## 1. Topics Covered

* Arrays in Go
* Declaring & initializing arrays
* Iterating over arrays
* Slices: creation & usage
* Slice operations: append, copy, slicing
* `make` function for slices
* Maps: creation, reading, writing, deleting
* Checking for key existence in maps

---

## 2. Detailed Notes

### Arrays

* Fixed-size collection of elements of the same type.
* Size is part of the type.

```go
var numbers [3]int // [0 0 0]
numbers[0] = 10
numbers[1] = 20
numbers[2] = 30
fmt.Println(numbers) // [10 20 30]
```

---

#### Array Initialization

```go
var arr = [3]int{1, 2, 3}
short := [...]string{"Go", "Java", "Python"} // size inferred
```

---

#### Iterating Over Arrays

```go
for i := 0; i < len(arr); i++ {
    fmt.Println(arr[i])
}

for index, value := range arr {
    fmt.Println(index, value)
}
```

---

### Slices

* More flexible than arrays (dynamic size).
* A slice points to an underlying array.

```go
numbers := []int{1, 2, 3}
fmt.Println(numbers)
```

---

#### Creating Slices

1. From an array:

   ```go
   arr := [5]int{10, 20, 30, 40, 50}
   slice := arr[1:4] // [20 30 40]
   ```
2. Using `make`:

   ```go
   s := make([]int, 3)       // length 3, capacity 3
   s2 := make([]int, 3, 5)   // length 3, capacity 5
   ```

---

#### Slice Operations

* Append:

  ```go
  nums := []int{1, 2}
  nums = append(nums, 3, 4)
  fmt.Println(nums) // [1 2 3 4]
  ```

* Copy:

  ```go
  a := []int{1, 2, 3}
  b := make([]int, len(a))
  copy(b, a)
  fmt.Println(b)
  ```

* Slicing:

  ```go
  s := []int{10, 20, 30, 40, 50}
  fmt.Println(s[1:4]) // [20 30 40]
  fmt.Println(s[:3])  // [10 20 30]
  fmt.Println(s[2:])  // [30 40 50]
  ```

---

### Maps

* Key-value pairs, unordered.
* Keys must be comparable (`string`, `int`, etc.).

```go
m := make(map[string]int)
m["Alice"] = 25
m["Bob"] = 30
fmt.Println(m) // map[Alice:25 Bob:30]
```

---

#### Reading from a Map

```go
age := m["Alice"]
fmt.Println("Alice's age:", age)
```

---

#### Checking if Key Exists

```go
value, exists := m["Charlie"]
if exists {
    fmt.Println("Charlie is", value)
} else {
    fmt.Println("Charlie not found")
}
```

---

#### Deleting from a Map

```go
delete(m, "Bob")
fmt.Println(m)
```

---

#### Map Literal Initialization

```go
scores := map[string]int{
    "Math":    90,
    "Science": 85,
}
```

---

## 3. Key Takeaways

* Arrays are fixed-size, slices are dynamic.
* Slices share the same underlying array → modifying one affects others referencing the same data.
* Use `append` to grow slices and `copy` to duplicate.
* Maps are Go’s built-in hash tables for key-value storage.
* Always check if a map key exists before using its value.

---

## 4. Exercises

1. Create an array of 5 integers, assign values, and print them using a loop.
2. Create a slice of strings, append two more strings, and print it.
3. Copy a slice into a new slice and modify one to see if it affects the other.
4. Create a map of countries and capitals, add two entries, delete one, and check if a key exists.

---


