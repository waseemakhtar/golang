package main

import "fmt"

func main() {

	fmt.Println(rotateString("hello", "llohe"))

}

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		return true
	}

	var r []byte
	l := len(A)

	for l > 0 {
		temp := A[0]
		r = []byte(A[1:])
		r = append(r, temp)
		A = string(r)
		fmt.Println(l, A, B)
		if A == B {
			return true
		}
		l = l-1
	}
	return false
}