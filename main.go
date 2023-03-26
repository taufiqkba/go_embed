package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed test/version.txt
var version string

//go:embed test/logo.jpg
var logo []byte

//go:embed test/files/*.txt
var path embed.FS

func main() {
	fmt.Println(version)

	err := ioutil.WriteFile("logo_new.jpg", logo, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	dirEntry, _ := path.ReadDir("test/files")
	for _, entry := range dirEntry {
		if !entry.IsDir() {
			fmt.Println(entry.Name())
			file, _ := path.ReadFile("test/files/" + entry.Name())
			fmt.Println(string(file))
		}
	}
}
