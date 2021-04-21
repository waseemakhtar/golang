// main.go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	env := os.Getenv("BPLOGGER_STDERR")
	env = strings.ToLower(strings.TrimSpace(env))
	switch env {
	case "1", "y", "yes", "stderr", "true":
		fmt.Printf("%s", env)
	}
}
