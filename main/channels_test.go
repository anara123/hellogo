package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func ExampleSimpleChannel() {
	var c = make(chan int)
	var done = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		close(c)
	}()

	for k := range c {
		fmt.Println(k)
	}
}

// Multiple Workers style
// The message is not emited to all the readers,
// instead only single reader gets the message
// That is why it is look more like workers
func ExampleMultipleReadersFromSingleChannel() {
	var c = make(chan int)
	var done = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	go func() {
		for n := range c {
			fmt.Println("Foo:", n)
			time.Sleep(time.Millisecond)
		}
		done <- true
	}()

	go func() {
		for n := range c {
			fmt.Println("Moo:", n)
			time.Sleep(time.Millisecond)
		}
		done <- true
	}()

	<-done
	<-done
}

func ExampleChanPassedAsArgument() {
	var c = incrementor()
	var cSum = puller(c)

	for sum := range cSum {
		fmt.Println(sum)
	}

	// Output:
	// 0
	// 1
	// 3
	// 6
	// 10
}

func incrementor() chan int {
	var out = make(chan int)

	go func() {
		for i := 0; i <= 4; i++ {
			out <- i
		}
		close(out)
	}()

	return out
}

func puller(c chan int) chan int {
	var out = make(chan int)
	go func() {
		var sum int
		for i := range c {
			sum += i
			out <- sum
		}
		close(out)
	}()

	return out
}

func TestFuctorial(t *testing.T) {
	var assert = assert.New(t)

	assert.Equal(24, <-fuctorial(4))
}

func TestFuctorialInParallel(t *testing.T) {
	// var assert = assert.New(t)
	var c = gen()
	var out2 = make(chan int)
	go func() {
		for n := range c {
			out2 <- <-fuctorial(n)
		}
		close(out2)
	}()

	for fuc := range out2 {
		fmt.Println(fuc)
	}
}

func gen() chan int {
	var out = make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 3; j < 13; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func fuctorial(n int) chan int {
	var out = make(chan int)

	go func() {
		var fuc = 1

		for i := 1; i <= n; i++ {
			fuc *= i
		}
		out <- fuc
		close(out)
	}()

	return out
}
