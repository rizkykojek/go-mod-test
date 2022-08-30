package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Before UnitTEST ......")

	m.Run()

	fmt.Println("After UnitTEST ......")
}

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

func TestSayHelloTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{"asep", "asep", "Hello asep"},
		{"john", "john", "Hello john"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SayHello(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("edo")
	}
}

func BenchmarkSayHelloLongParams(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("jojojojojojojojojojojojojojojojojojojojoj")
	}
}

func BenchmarkSayHelloTableTest(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{"asep", "asep"},
		{"john", "john"},
		{"jojojojojojojojojojojojojojojojojojojojoj", "jojojojojojojojojojojojojojojojojojojojoj"},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			SayHello(benchmark.request)
		})
	}
}
