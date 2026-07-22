package main

import "fmt"

func main() {
	// var a uint8 = 1           // 00000001
	// var b uint8 = 1 << 3      // 00001000
	// fmt.Printf("%08b\n", a&b) // 00001001

	var three uint8 = 3
	fmt.Printf("%08b\n", three)
	fmt.Printf("%08b\n", ^three)
	// var seven uint8 = 7
	// fmt.Printf("%08b\n", seven)
	// fmt.Printf("%08b\n", three&seven)
	// fmt.Printf("%08b\n", three|seven)

}
