package main

import (
	"fmt"
)

// Function to swap two strings
func swap(a, b string) (string, string) {
	return b, a
}

// Function to calculate the sum of integers
func sum(a, b int) int {
	return a + b
}

// Variadic function to calculate the sum of multiple integers
func variadicSum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// Function to create a closure
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Recursive function to calculate factorial
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Using the swap function
	a, b := "hello", "world"
	swappedA, swappedB := swap(a, b)
	fmt.Printf("Swapped values: %s, %s\n", swappedA, swappedB)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Using the sum function
	fmt.Printf("Sum of 5 and 7: %d\n", sum(5, 7))

	fmt.Println("-----------------------------------------------------------------------------------")

	// Using the variadicSum function
	fmt.Printf("Sum of 1, 2, 3, 4, 5: %d\n", variadicSum(1, 2, 3, 4, 5))

	fmt.Println("-----------------------------------------------------------------------------------")

	// Using the closure
	multiplierBy2 := makeMultiplier(2)
	fmt.Printf("5 multiplied by 2: %d\n", multiplierBy2(5))

	fmt.Println("-----------------------------------------------------------------------------------")

	// Using the factorial function
	fmt.Printf("Factorial of 5: %d\n", factorial(5))

	fmt.Println("-----------------------------------------------------------------------------------")

}
