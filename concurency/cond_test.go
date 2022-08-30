package concurency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done ", value)
	cond.L.Unlock()
}

func TestCondSignal(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal() // one by one
		}
	}()

	group.Wait()
}

func TestCondBroadcast(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}

	cond.Broadcast()

	group.Wait()
}
