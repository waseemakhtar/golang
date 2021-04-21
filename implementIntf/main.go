package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)


func Print(r io.Reader) io.Reader {
	b, _ := ioutil.ReadAll(r)
	fmt.Printf(string(b))
	return
}

func main() {
	Print(strings.NewReader("helloWorld"))
}

