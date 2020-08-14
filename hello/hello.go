package main

import (
	"fmt"

	"github.com/godsman-yang/hello/morestrings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	fmt.Println("Hello, World.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
