
---

# Golang Bootcamp - Session 8

**Duration:** 1.5 hours
**Topic:** Concurrency in Go – Goroutines & Channels

---

## 1. Topics Covered

* Introduction to concurrency in Go
* Goroutines
* Channels
* Buffered vs Unbuffered channels
* Channel directions (`chan<-` and `<-chan`)
* Closing channels
* `select` statement
* Common concurrency patterns

---

## 2. Detailed Notes

### Introduction to Concurrency

* **Concurrency**: dealing with many tasks at once (not necessarily at the same time).
* Go’s concurrency is built around **goroutines** and **channels**.

---

### Goroutines

* A goroutine is a lightweight thread managed by the Go runtime.
* Created by prefixing a function call with `go`.

```go
package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello from goroutine")
}

func main() {
    go sayHello()
    time.Sleep(time.Second) // Give goroutine time to run
}
```

* Goroutines are **cheap**: thousands can run in a single program.

---

### Channels

* Channels are pipes that connect goroutines.
* Use them to send and receive values between goroutines.

```go
package main

import "fmt"

func main() {
    ch := make(chan string)

    go func() {
        ch <- "Hello Channel"
    }()

    msg := <-ch
    fmt.Println(msg)
}
```

---

### Buffered vs Unbuffered Channels

**Unbuffered Channel**:

* Sends block until another goroutine receives.
* Receives block until another goroutine sends.

**Buffered Channel**:

* Allows sending multiple values before blocking, up to its capacity.

```go
ch := make(chan int, 2) // buffered channel
ch <- 1
ch <- 2
fmt.Println(<-ch)
fmt.Println(<-ch)
```

---

### Channel Directions

* Function parameters can restrict channel direction.

```go
func sendData(ch chan<- string) {
    ch <- "data"
}

func receiveData(ch <-chan string) {
    fmt.Println(<-ch)
}

func main() {
    ch := make(chan string)
    go sendData(ch)
    receiveData(ch)
}
```

---

### Closing Channels

* Use `close(channel)` to signal no more values will be sent.

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
close(ch)

for val := range ch {
    fmt.Println(val)
}
```

---

### `select` Statement

* Allows waiting on multiple channel operations.

```go
ch1 := make(chan string)
ch2 := make(chan string)

go func() { ch1 <- "from ch1" }()
go func() { ch2 <- "from ch2" }()

select {
case msg1 := <-ch1:
    fmt.Println(msg1)
case msg2 := <-ch2:
    fmt.Println(msg2)
}
```

---

## 3. Common Concurrency Patterns

**Fan-out, Fan-in**:

* Fan-out: multiple goroutines receive data from one channel.
* Fan-in: multiple goroutines send data into one channel.

**Worker Pools**:

* Fixed number of worker goroutines that process jobs from a channel.

---

## 4. Key Takeaways

* Goroutines are the foundation of concurrency in Go.
* Channels are the safe way to communicate between goroutines.
* Always consider channel buffering and blocking behavior.
* `select` allows you to multiplex channel operations.
* Close channels when no more values will be sent.

---

## 5. Exercises

1. Write a program with two goroutines: one sends numbers to a channel, and one prints them.
2. Implement a buffered channel to queue messages and print them in order.
3. Create a worker pool that processes a list of integers by squaring them.

---

