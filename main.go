package main

import (
	"fmt"

	"goji.io"
)

func main() {
	foo := goji.NewMux()
	foo.Handle()
	fmt.Println("Hello world!")
}
