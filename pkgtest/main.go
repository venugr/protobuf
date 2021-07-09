package main

import (
	"fmt"

	"github.com/venugr/protobuf/pkgtest/stuff"
)

func main() {
	fmt.Println("Hello, Package Test")
	fmt.Println(stuff.Foo())
}