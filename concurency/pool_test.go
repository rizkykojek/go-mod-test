package concurency

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{New: func() interface{} {
		return "New"
	}}

	pool.Put("A")
	pool.Put("B")
	pool.Put("C")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(10 * time.Nanosecond) // Some of pool will get default because still used
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Finished...")
}
