package main

import (
	"fmt"
)

func helloWorld(f func())  {

	fmt.Println("In helloWorld")
	f()
}

func f1(a string) func() {
	return func() {
		fmt.Printf("Hello World from %s\n", a)
	}
}

func f2(a string) func() {
	return func() {
		fmt.Printf("Hello World from %s\n", a)
	}
}

func main() {

    helloWorld(f1("f1"))
	helloWorld(f2("f2"))
}
