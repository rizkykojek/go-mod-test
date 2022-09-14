package embed

import (
	_ "embed"
	"fmt"
	"testing"
)

//go:embed VERSION.txt
var version string

//go:embed VERSION.txt
var version2 string

func TestEmbedString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}
