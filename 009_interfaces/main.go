package main

import (
	"fmt"
)

// Define an interface
type Speaker interface {
	Speak() string
}

// Define a struct
type Dog struct {
	Name string
}

// Implement the Speak method for the Dog type
func (d Dog) Speak() string {
	return "Woof!"
}

// Define a Cat struct
type Cat struct {
	Name string
}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Function that accepts any type
func PrintValue(val interface{}) {
	fmt.Println(val)
}

func PrintValue1(val interface{}) {
	// Type assertion to extract the underlying value as an int
	if v, ok := val.(int); ok {
		fmt.Printf("Integer value: %d\n", v)
	} else {
		fmt.Println("Not an integer")
	}
}

func PrintValue2(val interface{}) {
	// Type switch to handle different types
	switch v := val.(type) {
	case int:
		fmt.Printf("Integer value: %d\n", v)
	case string:
		fmt.Printf("String value: %s\n", v)
	case float64:
		fmt.Printf("Float64 value: %f\n", v)
	default:
		fmt.Println("Unknown type")
	}
}

// Define an interface with multiple methods
type Vehicle interface {
	Drive() string
	Horn() string
}

// Define a Car struct
type Car struct {
	Brand string
}

// Implement the Vehicle interface for Car
func (c Car) Drive() string {
	return "Driving the car"
}

func (c Car) Horn() string {
	return "Beep beep!"
}

// Define an interface
type Flyer interface {
	Fly() string
}

// Define a struct with a method that satisfies the Flyer interface
type Bird struct {
	Species string
}

func (b Bird) Fly() string {
	return "Flying high!"
}

// Define another struct that embeds Bird and adds new methods
type Eagle struct {
	Bird
}

func (e Eagle) Hunt() string {
	return "Hunting prey!"
}

// Implement the HuntPrey method that accepts a prey name as a parameter
func (e Eagle) HuntPrey(prey string) string {
	return fmt.Sprintf("The eagle is hunting a %s!", prey)
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Interface Declaration

	// Create an instance of Dog
	dog := Dog{Name: "Buddy"}

	// Declare a variable of type Speaker and assign it to a Dog instance
	var speaker Speaker = dog

	// Call the Speak method
	fmt.Println(speaker.Speak()) // Output: Woof!

	fmt.Println("-----------------------------------------------------------------------------------")

	// Multiple Types Implementing the Same Interface

	cat := Cat{Name: "Whiskers"}

	var speaker2 Speaker = cat

	fmt.Println(speaker2.Speak()) // Output: Meow!

	fmt.Println("-----------------------------------------------------------------------------------")

	// Empty Interface (interface{})

	// Using the PrintValue function with different types
	PrintValue(42)           // Output: 42
	PrintValue("Hello, Go!") // Output: Hello, Go!
	PrintValue(3.14)         // Output: 3.14

	fmt.Println("-----------------------------------------------------------------------------------")

	// Type Assertion
	// Type assertion is a way to extract the underlying value of an interface

	PrintValue1(42)           // Output: Integer value: 42
	PrintValue1("Hello, Go!") // Output: Not an integer

	fmt.Println("-----------------------------------------------------------------------------------")

	// Type Switch

	PrintValue(42)           // Output: Integer value: 42
	PrintValue("Hello, Go!") // Output: String value: Hello, Go!
	PrintValue(3.14)         // Output: Float64 value: 3.140000

	fmt.Println("-----------------------------------------------------------------------------------")

	// Interfaces with Multiple Methods

	// Create an instance of Car
	car := Car{Brand: "Toyota"}

	// Declare a variable of type Vehicle and assign it to a Car instance
	var vehicle Vehicle = car

	// Call the Drive and Horn methods
	fmt.Println(vehicle.Drive()) // Output: Driving the car
	fmt.Println(vehicle.Horn())  // Output: Beep beep!

	fmt.Println("-----------------------------------------------------------------------------------")

	// Interfaces with Multiple Methods

	// Create an instance of Eagle
	eagle := Eagle{Bird{Species: "Eagle"}}

	// The Eagle can use methods from Bird as well as its own methods
	fmt.Println(eagle.Fly())              // Output: Flying high!
	fmt.Println(eagle.Hunt())             // Output: Hunting prey!
	fmt.Println(eagle.HuntPrey("rabbit")) // Output: The eagle is hunting a rabbit!

	fmt.Println("-----------------------------------------------------------------------------------")

}
