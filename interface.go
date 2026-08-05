package main

import "fmt"

type MyInterface interface {
	String() string
	Print()
}

type MyString string

func (s MyString) String() string {
	return string(s)
}

func (s MyString) Print() {
	fmt.Println(s)
}

type MyStruct struct {
	name string
}

func (s *MyStruct) String() string {
	return s.name
}

func (s *MyStruct) Print() {
	fmt.Println(s.name)
}

func main() {
	s := newObj()
	s.Print()

	myFunction(&MyStruct{name: "struct"})
	myFunction(s)
}

func newObj() MyInterface {
	var str MyString = "abc"
	return str
}

func myFunction(m MyInterface) {
	str, ok := m.(MyString)
	if ok {
		fmt.Println("this is sting ", str)
	} else {
		fmt.Println("this is something else", m)
	}
}
