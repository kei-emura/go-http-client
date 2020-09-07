package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

func main() {
	src := "テスト"
	ascii, err := idna.ToASCII(src)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s -> %s\n", src, ascii)
}
