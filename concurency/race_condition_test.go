package concurency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				counter++
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Last Counter : ", counter)
}

func TestRaceConditionWithLocking(t *testing.T) {
	counter := 0
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	assert.Equal(t, 100000, counter)
	fmt.Println("Last Counter : ", counter)
}

func TestRaceConditionWithAtomic(t *testing.T) {
	var counter int64 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			defer group.Done()
			group.Add(1)
			for i := 0; i < 100; i++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	group.Wait()
	assert.Equal(t, int64(100000), counter)
	fmt.Println("Last Counter : ", counter)
}
