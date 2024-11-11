package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @author: Aakash B
 * @date: 2024/11/11 20:15
 * @description: Demonstrates concurrency in Go.
 * @file: concurrency.go
 * @summary: This example shows how to use concurrency in Go.
 * @link: https://go.dev/tour/concurrency

Concept/Topic: Concurrency in Go
--------------------------------
	Purpose:
		Go is designed with concurrency in mind, using goroutines and channels as primary tools.
		Concurrency enables multiple tasks to execute independently, making Go suitable for high-performance,
		parallel applications. Concurrency primitives like goroutines, channels, and synchronization mechanisms
		(mutex and sync packages) help in managing safe data sharing and coordination across multiple tasks.

	Core Components:
		Goroutines:
			Lightweight threads managed by the Go runtime for concurrent task execution.
		Channels:
			Mechanism for communication and synchronization between goroutines,
			available as unbuffered or buffered.
		Select Statements:
			Control structure to manage multiple channel operations and coordinate between them.
		Mutex and sync Packages:
			Tools for locking mechanisms to ensure safe access to shared data in concurrent environments.

Explanation:
------------
	Goroutines:
		go printMessage("Hello from goroutine!"): Launches a new goroutine.
		go keyword initiates a goroutine, which runs concurrently with the main function.

	WaitGroup (sync.WaitGroup):
		Used to wait for multiple goroutines to complete. wg.Add(2) sets the number of goroutines
		to wait for, and defer wg.Done() marks completion.

	Unbuffered Channel:
		msgChan := make(chan string):
			Creates an unbuffered channel. Unbuffered channels block the sender until a receiver
			is ready to receive.
	The main function receives the message sent by the goroutine through msgChan.

	Buffered Channel:
		bufferedChan := make(chan int, 2):
			Creates a buffered channel with capacity 2.
			Buffered channels allow sending and receiving without immediate synchronization
			if they’re not full or empty.
			Both values sent to bufferedChan are retrieved later without blocking.

	Select Statements:
		select enables waiting on multiple channel operations.
		The select statement executes the first case with an available value.
		time.After implements a timeout, which is triggered if no channels are ready within a set duration.

	Mutex and sync Package:
		mutex.Lock() and mutex.Unlock():
			Ensure that only one goroutine accesses the shared variable (counter) at a time.
			Mutex is used here to prevent race conditions when multiple goroutines modify the shared counter variable.

Use Cases in Data Structures and Algorithms:
-------------------------------------------
	Parallel Processing:
		Execute tasks like sorting subarrays or performing calculations in parallel.
	Producer-Consumer Problem:
		Channels and select statements help coordinate producers and consumers effectively.
	Resource Management:
		Mutexes provide controlled access to shared resources (e.g., caches or databases).

Key Points:
-----------
	Goroutines:
		Lightweight and efficient, enabling highly concurrent applications.

	Channels:
		Allow safe communication and synchronization, especially between goroutines.

	Select Statements:
		Provide flexibility in managing multiple channel operations,
		useful for coordinating concurrent tasks.

	Mutexes:
		Crucial for synchronizing access to shared data in concurrent data structures
		and avoiding race conditions.

Practical Usage:
----------------
	Web Servers:
		Handle multiple client requests concurrently using goroutines and channels.
	Task Queues:
		Use buffered channels for managing tasks in worker pools.
	Shared State Management:
		Employ mutexes to protect shared state in cache implementations or databases.

Notes:
------
	Concurrency in Go emphasizes simplicity and efficiency, with goroutines providing lightweight
	threads and channels enabling straightforward inter-goroutine communication.
	Understanding these concurrency tools and best practices helps avoid common pitfalls,
	like deadlocks and race conditions, while building responsive and scalable applications.

*/

func main() {
	// 1. Goroutines
	go printMessage("Hello from goroutine!")

	// Using WaitGroup to wait for goroutines to complete
	// 2. WaitGroup
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		printMessage("Hello from goroutine 1!")
	}()

	go func() {
		defer wg.Done()
		printMessage("Hello from goroutine 2!")
	}()

	wg.Wait()

	// Using Buffered Channel
	msgChan := make(chan string)
	defer close(msgChan)

	go func() {
		msgChan <- "Hello from goroutine!"
	}()

	msg := <-msgChan
	if msg != "" {
		fmt.Println(msg)
	}

	// 3. Buffered Channel
	bufferedChan := make(chan int, 2)

	bufferedChan <- 1
	bufferedChan <- 2

	fmt.Println(<-bufferedChan)
	fmt.Println(<-bufferedChan)

	// 4. Select Statements
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "One"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Two"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}

	// 5. Mutex and sync Package
	var counter int = 0
	var mutex sync.Mutex
	for i := 0; i < 5; i++ {
		go func() {
			mutex.Lock()
			defer mutex.Unlock()
			counter++
			fmt.Println("Counter:", counter)
		}()
	}

	wg.Wait()
	time.Sleep(2 * time.Second)

}

// printMessage Prints a given string to the console.
func printMessage(s string) {
	fmt.Println(s)
}
