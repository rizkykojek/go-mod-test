package embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed VERSION.txt
var version string

//go:embed VERSION.txt
var version2 string

func TestString(t *testing.T) {
	fmt.Println(version)
	fmt.Println(version2)
}

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	err := os.WriteFile("logo_next.png", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

func TestMultipleFiles(t *testing.T) {
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestPath(t *testing.T) {
	directoryEntries, _ := files.ReadDir("files")
	for _, entry := range directoryEntries {
		if !entry.IsDir() {
			fmt.Print(entry.Name())
			file, _ := files.ReadFile("files/" + entry.Name())
			fmt.Println(", Contect " + string(file))
		}
	}
}
