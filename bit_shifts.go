package main

import "fmt"

func main() {
	var unsignedLeft uint8 = 1
	fmt.Printf("%08b\n", unsignedLeft)
	fmt.Printf("%08b\n", unsignedLeft<<3)

	var unsignedRight uint8 = 1 << 7
	fmt.Printf("%08b\n", unsignedRight)
	fmt.Printf("%08b\n", unsignedRight>>3)

	var signedLeft int8 = 1
	fmt.Printf("%08b\n", signedLeft)
	fmt.Printf("%08b\n", signedLeft<<3)

	var signedRight int8 = -128
	fmt.Printf("%08b\n", byte(signedRight))
	fmt.Printf("%08b\n", byte(signedRight>>3))
}
