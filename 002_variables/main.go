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
