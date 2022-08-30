package concurency

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func() {
			time.Sleep(3 * time.Second)
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU", totalCpu)

	// runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
}
