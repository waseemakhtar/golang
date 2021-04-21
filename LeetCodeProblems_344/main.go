package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l'}

	fmt.Println(string(s))
	reverseString(s)
	fmt.Println(string(s))

}

func reverseString(s []byte)  {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}
}