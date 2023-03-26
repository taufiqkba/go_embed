package go_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

// Embed string
//
//go:embed version.txt
var version string

func TestString(t *testing.T) {
	fmt.Println(version)
}

// Embed file to []byte type (from image)
//
//go:embed logo.jpg
var logo []byte

func TestByteArray(t *testing.T) {
	err := ioutil.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}

// Embed multiple files
//
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt

var files embed.FS

func TestMultipleFile(t *testing.T) {
	a, err := files.ReadFile("files/a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}
