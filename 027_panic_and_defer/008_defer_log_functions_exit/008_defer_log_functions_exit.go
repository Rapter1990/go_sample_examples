package main

import (
	"fmt"
)

func main() {
	defer exitLog("main")
	fmt.Println("Main function running")
	performTask()
}

func performTask() {
	defer exitLog("performTask")
	fmt.Println("Task performed")
	panic("Panic in performTask")
}

func exitLog(funcName string) {
	fmt.Printf("Exiting %s\n", funcName)
}
