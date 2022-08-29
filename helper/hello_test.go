package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestSayHelloAssert(t *testing.T) {
	result := SayHello("kojek")
	assert.Equal(t, "Hello kojek ", result)
	fmt.Println("TestSayHelloAssert Done")
}

func TestSayHelloRequire(t *testing.T) {
	result := SayHello("kojek")
	require.Equal(t, "Hello kojek ", result)
	fmt.Println("TestSayHelloRequire Done")
}
