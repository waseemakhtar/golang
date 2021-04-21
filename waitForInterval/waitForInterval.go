package waitForInterval

import (
	"fmt"
	"time"
)

func main() {
	max := 5 * time.Second
	interval := 1 * time.Second

	waitForInterval(max, interval)
	fmt.Printf("Waited for %s\n", max)

}

func waitForInterval(max, interval time.Duration) string {
	timeout := time.After(max)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		// Got a timeout! fail with a timeout error
		case <-timeout:
			fmt.Println("timeout, return")

			return "waited"
		// Got a tick, we should check on orchState active for the resource
		case <-ticker.C:
			fmt.Println("ticked")
			//return "ticked"
		}
	}
}
