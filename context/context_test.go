package context

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	contextE := context.WithValue(contextB, "d", "D")
	contextD := context.WithValue(contextB, "e", "E")

	contextF := context.WithValue(contextC, "f", "F")

	contextG := context.WithValue(contextF, "g", "G")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
	fmt.Println(contextF)
	fmt.Println(contextG)

	assert.Equal(t, "C", contextG.Value("c"))
}

func CreateCounterWithoutContext() chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0
		for {
			destination <- counter
			counter++
			time.Sleep(1 * time.Second)
		}
	}()
	return destination
}
func TestContextWithGoroutineLeak(t *testing.T) {
	fmt.Println("Before Total goroutine :", runtime.NumGoroutine())
	destination := CreateCounterWithoutContext()
	for i := range destination {
		fmt.Println("Counter : ", i)
		if i == 5 {
			break
		}
	}

	fmt.Println("After Total goroutine :", runtime.NumGoroutine())
}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second) // Slow simulation
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("Before Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()

	child, cancel := context.WithCancel(parent)
	destination := CreateCounter(child)
	fmt.Println("On Process Total goroutine :", runtime.NumGoroutine())
	for i := range destination {
		fmt.Println("Counter : ", i)
		if i == 5 {
			break
		}
	}
	cancel()
	time.Sleep(1 * time.Second)

	fmt.Println("After Total goroutine :", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Before Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()

	child, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(child)

	for i := range destination {
		fmt.Println("Counter : ", i)
	}
	time.Sleep(1 * time.Second)

	fmt.Println("After Total goroutine :", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("Before Total goroutine :", runtime.NumGoroutine())
	parent := context.Background()

	child, cancel := context.WithDeadline(parent, time.Now().Add(3*time.Second))
	defer cancel()

	destination := CreateCounter(child)

	for i := range destination {
		fmt.Println("Counter : ", i)
	}
	time.Sleep(1 * time.Second)

	fmt.Println("After Total goroutine :", runtime.NumGoroutine())
}
