package concurency

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Data Privacy"
		fmt.Println("Finished send data to channel...")
	}()

	fmt.Println("Ready to get Data...")
	data := <-channel
	fmt.Println("Received : " + data)
	time.Sleep(5 * time.Second)
}
