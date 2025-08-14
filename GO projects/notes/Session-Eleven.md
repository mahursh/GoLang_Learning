
# Go Bootcamp – Session 10

## Topic: Concurrency Patterns in Go

In this session, we explored advanced **concurrency patterns** in Go, focusing on how to build scalable and efficient concurrent programs.

---

## 1. Review: Concurrency in Go
- **Concurrency** is about dealing with lots of things at once (not the same as parallelism).
- Go provides **goroutines** and **channels** to manage concurrency safely.

---

## 2. Common Concurrency Patterns

### 2.1 Worker Pool Pattern
**Purpose:** Limit the number of goroutines working simultaneously and process jobs in parallel.

**Example:**
```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // simulate work
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Receive results
	for a := 1; a <= numJobs; a++ {
		fmt.Println(<-results)
	}
}
````

---

### 2.2 Fan-Out, Fan-In

* **Fan-Out:** Multiple goroutines read from the same channel and process data.
* **Fan-In:** Multiple goroutines send to the same channel, merging results.

**Example:**

```go
func main() {
	in := make(chan int)
	out := make(chan int)

	// Producer
	go func() {
		for i := 0; i < 10; i++ {
			in <- i
		}
		close(in)
	}()

	// Fan-Out (multiple workers)
	for w := 0; w < 3; w++ {
		go func(id int) {
			for num := range in {
				out <- num * 2
			}
		}(w)
	}

	// Fan-In (single consumer)
	go func() {
		for result := range out {
			fmt.Println("Result:", result)
		}
	}()

	time.Sleep(time.Second * 2)
}
```

---

### 2.3 Pipeline Pattern

* Chain multiple stages of processing using channels.

**Example:**

```go
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3, 4)
	out := square(in)

	for result := range out {
		fmt.Println(result)
	}
}
```

---

## 3. Context for Cancellation

* Use `context` to stop goroutines gracefully.

```go
import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
}
```

---

## 4. Best Practices

* Avoid goroutine leaks: always ensure goroutines can exit.
* Close channels when no longer sending.
* Use buffered channels to prevent blocking when appropriate.
* Use `sync.WaitGroup` for coordination.

---

## Homework

1. Implement a **worker pool** that downloads multiple URLs concurrently and stores their status codes.
2. Create a **pipeline** that:

   * Reads integers from a slice.
   * Squares them.
   * Filters only even results.
3. Add **context-based cancellation** to your worker pool.



---


