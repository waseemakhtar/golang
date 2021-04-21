package main

import (
	"fmt"
)

func main() {

	sortedStr := sortStr([]rune("Hello World"))
	fmt.Println(string([]rune("Hello World")))
	fmt.Println(string(sortedStr))
}

func sortStr(str []rune) []rune {

	j := 0
	for i, _ := range str {
		j = i+1
		for j < len(str) {
			if str[i] > str[j] {
				tmp := str[i]
				str[i] = str[j]
				str[j] = tmp
			}
			j++
		}
	}
	return str

}