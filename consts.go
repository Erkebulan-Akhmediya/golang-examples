package main

import "fmt"

type Weekday int

const (
	Monday = iota
	Tuesday
	Thursday
	Wednesday
	Friday
)

type Flags uint8

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func main() {
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Thursday)
	fmt.Println(Wednesday)
	fmt.Println(Friday)

	fmt.Printf("%08b\n", FlagUp)
	fmt.Printf("%08b\n", FlagBroadcast)
	fmt.Printf("%08b\n", FlagLoopback)
	fmt.Printf("%08b\n", FlagPointToPoint)
	fmt.Printf("%08b\n", FlagMulticast)
}
