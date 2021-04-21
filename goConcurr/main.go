package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range a {
		waitGroup.Add(1)
		i := i
		go func() {
			defer waitGroup.Done()
			printInt(i)
		}()
	}
	waitGroup.Wait()
}

func printInt(number int) {
	fmt.Println(number)
}
