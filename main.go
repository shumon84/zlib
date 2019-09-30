package main

import (
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
)

func Compose() {
	zw := zlib.NewWriter(os.Stdout)
	defer zw.Close()
	io.Copy(zw, os.Stdin)
}

func Decompose() {
	zr, err := zlib.NewReader(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()
	io.Copy(os.Stdout, zr)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage :", os.Args[0], "[c|d]")
		return
	}
	if os.Args[1] == "c" {
		Compose()
	} else if os.Args[1] == "d" {
		Decompose()
	}
}
