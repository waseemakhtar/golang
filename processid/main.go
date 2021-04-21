package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	pid := os.Getpid()
	fmt.Println("Own process identifier: ", pid)
	fmt.Println("Own process identifier: ", strconv.Itoa(pid))

	p, _ := os.FindProcess(pid)
	i := 0
	configureSignalHandling()
	for i <= 10 {
		_ = p.Signal(syscall.SIGUSR1)
		i++
		time.Sleep(1000000000)
	}
}

func configureSignalHandling() {

	var sigChan chan os.Signal
	sigChan = make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGUSR1)

	go func() {

		for {

			select {
			case usrSignal := <-sigChan:
				if usrSignal != syscall.SIGUSR1 {
					continue
				}

				fmt.Println("Got SIGNAL")
			}
		}

	}()
}
