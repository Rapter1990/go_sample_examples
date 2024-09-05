package main

import (
	"fmt"
	"time"
)

func main() {

	// Timer with Select ( Use select to handle timers alongside other concurrent events )
	// You can use timers within a select statement to wait for multiple events simultaneously,
	// including the timer firing

	timer1 := time.NewTimer(2 * time.Second) // create timer in 2 seconds
	timer2 := time.NewTimer(4 * time.Second) // create timer in 4 seconds

	select {
	case <-timer1.C:
		fmt.Println("Timer 1 fired")
	case <-timer2.C:
		fmt.Println("Timer 2 fired")
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout! No timer fired in 3 seconds.")
	}

}
