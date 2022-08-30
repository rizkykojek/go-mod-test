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

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizky Kojek"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "test 1"
	channel <- "test 2"
	channel <- "test 3"

	fmt.Println(<-channel)
	fmt.Println(<-channel)
	fmt.Println(<-channel)

	time.Sleep(2 * time.Second)
	fmt.Println("Finished...")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
		defer close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}

	fmt.Println("Finished...")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data from channel 1, ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data from channel 2, ", data)
			counter++
		default:
			fmt.Println("Waiting data...")
		}

		if counter == 2 {
			break
		}
	}
	fmt.Println("Finished...")
}
