Certainly! Here's a revised `README.md` for the "Variables" example based on the provided Go code:

---

# Go Sample Example - Variables

This repository contains a Go (Golang) program that demonstrates the use of variables, constants, and enumerations in Go. The example showcases how to declare and use different types of variables, constants, and enums in Go.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example introduces various types of variables including integers, floating-point numbers, strings, and booleans.</li>
  <li>It covers the declaration of constants and enumerations using `iota` for creating successive integer constants.</li>
  <li>Shows the use of both explicit and inferred types for variable declarations.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
	"math"
)

// Constants
const Pi = 3.14159

// Enumerations using iota
const (
	Red = iota // The iota keyword represents successive integer constants 0, 1, 2,…
	Green
	Blue
)

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Integer variables
	var intVar int = 10
	var intAuto = 20 // Type inference

	// Long integer (int64)
	var longVar int64 = 10000000000

	// Floating-point variables
	var floatVar float32 = 3.14
	var doubleVar float64 = math.Pi // Double precision floating-point

	// String variables
	var strVar string = "Hello, Go!"
	var strAuto = "Welcome to Go programming!" // Type inference

	// Boolean variable
	var boolVar bool = true

	// Using short variable declaration
	name := "John Doe"
	age := 30
	salary := 55000.50
	isEmployed := true

	// Printing variables
	fmt.Println("Integer variable:", intVar)
	fmt.Println("Auto-inferred integer:", intAuto)
	fmt.Println("Long variable (int64):", longVar)
	fmt.Println("Float variable:", floatVar)
	fmt.Println("Double variable:", doubleVar)
	fmt.Println("String variable:", strVar)
	fmt.Println("Auto-inferred string:", strAuto)
	fmt.Println("Boolean variable:", boolVar)

	// Printing short-declared variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Salary:", salary)
	fmt.Println("Is Employed:", isEmployed)

	// Constants
	fmt.Println("Constant Pi:", Pi)

	// Enumerations
	fmt.Println("Red:", Red)
	fmt.Println("Green:", Green)
	fmt.Println("Blue:", Blue)

	fmt.Println("-----------------------------------------------------------------------------------")

}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:
   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```
3. Navigate to the `002_variables` directory:
   ```bash
   cd go_sample_examples/002_variables
   ```
4. Run the Go program:
   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
Integer variable: 10
Auto-inferred integer: 20
Long variable (int64): 10000000000
Float variable: 3.14
Double variable: 3.141592653589793
String variable: Hello, Go!
Auto-inferred string: Welcome to Go programming!
Boolean variable: true
Name: John Doe
Age: 30
Salary: 55000.5
Is Employed: true
Constant Pi: 3.14159
Red: 0
Green: 1
Blue: 2
-----------------------------------------------------------------------------------
```