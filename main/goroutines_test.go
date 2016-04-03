package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var counter int
var atomicCounter int64
var mutex sync.Mutex

func ExampleRunParallelWithMutex() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)

	go RunParallelWithMutex("Foo:")
	go RunParallelWithMutex("Goo:")

	wg.Wait()
	fmt.Println("Final Counter:", counter)

	// Output: Final Counter: 40
}

func RunParallelWithMutex(s string) {
	for i := 0; i < 20; i++ {
		mutex.Lock()
		x := counter
		x++
		time.Sleep(3 * time.Millisecond)
		counter = x
		mutex.Unlock()
		// fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

func ExampleRunParallelRaceCondition() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)

	go RunParallelRaceCondition("Foo:")
	go RunParallelRaceCondition("Goo:")

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}

func RunParallelRaceCondition(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(3 * time.Millisecond)
		counter = x
	}
	wg.Done()
}

func ExampleAtomicIncrement() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg.Add(2)

	go AtomicIncrement("Foo:")
	go AtomicIncrement("Goo:")

	wg.Wait()
	fmt.Println("Final Counter:", atomicCounter)

	// Output: Final Counter: 40
}

func AtomicIncrement(s string) {
	for i := 0; i < 20; i++ {
		atomic.AddInt64(&atomicCounter, 1)
	}
	wg.Done()
}
