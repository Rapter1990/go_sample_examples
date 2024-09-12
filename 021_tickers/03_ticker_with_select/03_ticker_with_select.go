package main

import (
	"fmt"
	"time"
)

func main() {

	// Using Ticker with Select
	// Use a ticker inside a select statement to manage multiple concurrent events

	ticker := time.NewTicker(1 * time.Second)
	stop := make(chan bool)

	go func() {
		time.Sleep(3 * time.Second)
		stop <- true
	}()

	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
		case <-stop:
			fmt.Println("Stop signal received.")
			ticker.Stop()
			return
		}
	}

	// The select statement listens for ticks and a stop signal
	// When the stop channel receives a signal, the ticker is stopped, and the program exits the loop

}
