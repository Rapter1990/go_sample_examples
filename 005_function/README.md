# Go Sample Example - Functions

This repository demonstrates various types of functions in Go, including simple functions, variadic functions, closures, and recursive functions. It showcases how to define and use functions to perform common programming tasks.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic function definitions and usages in Go.</li>
  <li>It includes functions for swapping values, calculating sums, handling variable-length arguments, creating closures, and computing factorials.</li>
</ul>

## 💻 Code Example

```golang
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
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:
   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```
3. Navigate to the `005_function` directory:
   ```bash
   cd go_sample_examples/005_function
   ```
4. Run the Go program:
   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Swapped values: world, hello
Sum of 5 and 7: 12
Sum of 1, 2, 3, 4, 5: 15
5 multiplied by 2: 10
Factorial of 5: 120
```
