package concurency

import (
	"fmt"
	"github.com/rizkykojek/go-mod-test/v2/helper"
	"testing"
	"time"
)

func RunSayHello() {
	fmt.Println(helper.SayHello("kojek"))
}

func TestCreateGoroutine(t *testing.T) {
	go RunSayHello()
	fmt.Println("Opps")

	time.Sleep(1 * time.Second)
}

func DisplayNumber(i int) {
	fmt.Println(i)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
