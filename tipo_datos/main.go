package main

import "fmt"

func main() {
	var entero32 uint32
	var entero64 uint64

	entero32 = 25
	entero64 = 85

	fmt.Println(entero32 + uint32(entero64))
}
