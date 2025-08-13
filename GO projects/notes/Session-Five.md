

---

# Golang Bootcamp - Session 5

**Duration:** 1.5 hours
**Topic:** Structs & Methods

---

## 1. Topics Covered

* Struct basics & declaration
* Initializing structs
* Accessing & modifying struct fields
* Anonymous structs
* Nested structs
* Methods on structs
* Pointer receivers vs value receivers
* Method chaining basics

---

## 2. Detailed Notes

### Struct Basics

A **struct** is a collection of fields (like an object with properties in other languages).

```go
type Person struct {
    Name string
    Age  int
}
```

---

### Initializing Structs

1. **Named field initialization**:

   ```go
   p1 := Person{Name: "Alice", Age: 25}
   ```

2. **Positional initialization** (not recommended for large structs):

   ```go
   p2 := Person{"Bob", 30}
   ```

3. **Zero-value struct**:

   ```go
   var p3 Person
   p3.Name = "Charlie"
   p3.Age = 40
   ```

---

### Accessing & Modifying Struct Fields

```go
fmt.Println(p1.Name) // Alice
p1.Age = 26
```

---

### Anonymous Structs

You can define a struct without giving it a name.

```go
student := struct {
    ID    int
    Grade string
}{ID: 1, Grade: "A"}
```

---

### Nested Structs

```go
type Address struct {
    City, Country string
}

type Employee struct {
    Name    string
    Address Address
}

emp := Employee{
    Name:    "John",
    Address: Address{City: "Paris", Country: "France"},
}
fmt.Println(emp.Address.City)
```

---

### Methods on Structs

A method is a function with a **receiver**.

**Value receiver example**:

```go
func (p Person) Greet() {
    fmt.Println("Hello, my name is", p.Name)
}
```

---

**Pointer receiver example**:

```go
func (p *Person) HaveBirthday() {
    p.Age++
}
```

---

### Pointer vs Value Receiver

* **Value receiver** → method gets a copy of the struct (changes don’t affect original).
* **Pointer receiver** → method gets the struct’s memory address (changes affect original).
* Pointer receivers are preferred if:

  * The method modifies the struct.
  * The struct is large (avoid copying).

---

### Method Chaining

Return the struct pointer from methods to allow chaining.

```go
func (p *Person) SetName(name string) *Person {
    p.Name = name
    return p
}

func (p *Person) SetAge(age int) *Person {
    p.Age = age
    return p
}

func main() {
    p := &Person{}
    p.SetName("Diana").SetAge(28).Greet()
}
```

---

## 3. Key Takeaways

* Structs group related data together.
* Methods make structs more powerful, allowing behavior to be attached.
* Use pointer receivers to modify the struct or for performance.
* Anonymous and nested structs are useful for quick or hierarchical data models.
* Method chaining improves readability for consecutive operations.

---

## 4. Exercises

1. Define a `Car` struct with fields `Brand`, `Model`, and `Year`.
   Create a method `DisplayInfo()` to print the car details.
2. Create a `BankAccount` struct with fields `Owner` and `Balance`.
   Add methods `Deposit(amount float64)` and `Withdraw(amount float64)`.
   Ensure withdrawals don’t allow negative balances.
3. Create a nested struct `Library` with `Name` and a slice of `Book` structs.
   Add a method to display all books in the library.

---

 
