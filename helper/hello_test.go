package helper

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	result := SayHello("kojek")

	if result != "Hello kojek" {
		t.Error("Result must be 'Hello kojek")
	}

	fmt.Println("TestSayHello Done")
}

func TestSayHelloFailed(t *testing.T) {
	result := SayHello("rizky")

	if result != "Hello rizky " {
		t.Error("Result must be 'Hello rizky")
	}

	fmt.Println("TestSayHelloFailed Done")
}
