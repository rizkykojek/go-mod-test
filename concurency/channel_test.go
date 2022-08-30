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

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizky Kojek"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}
