package main

import (
	// "bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run aws-extract-gz-file-example.go file.gz")
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	zr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	f2, err := os.Create("ungzipped-file.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f2.Close()

	if _, err := io.Copy(f2, zr); err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("file has been written")
}
