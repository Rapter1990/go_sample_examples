package main

import (
	"fmt"
)

// Function that does not modify the original value
func zeroVal(val int) {
	val = 0
}

// Function that modifies the value through a pointer
func zeroPtr(ptr *int) {
	*ptr = 0
}

// Define a struct
type Person struct {
	Name string
	Age  int
}

// Function that modifies the struct via a pointer
func updateAge(p *Person, newAge int) {
	p.Age = newAge
}

// Function that appends a value to a slice using a pointer
func appendValue(slice *[]int, value int) {
	*slice = append(*slice, value)
}

// Function that adds a key-value pair to a map using a pointer
func addEntry(m *map[string]int, key string, value int) {
	(*m)[key] = value
}

func main() {

	// --------------------------------------------------------------------------------------------

	// Example with value
	x := 5
	fmt.Println("Before zeroVal:", x) // Output: Before zeroVal: 5
	zeroVal(x)
	fmt.Println("After zeroVal:", x) // Output: After zeroVal: 5

	// Example with pointer
	y := 5
	fmt.Println("Before zeroPtr:", y) // Output: Before zeroPtr: 5
	zeroPtr(&y)
	fmt.Println("After zeroPtr:", y) // Output: After zeroPtr: 0

	// --------------------------------------------------------------------------------------------

	// Create an instance of Person
	person := Person{Name: "Alice", Age: 30}

	// Print the original Person
	fmt.Println("Original person:", person) // Output: {Alice 30}

	// Update age using a pointer
	updateAge(&person, 35)

	// Print the updated Person
	fmt.Println("Updated person:", person) // Output: {Alice 35}

	// --------------------------------------------------------------------------------------------

	// Declare a slice
	numbers := []int{1, 2, 3}

	// Print the original slice
	fmt.Println("Original slice:", numbers) // Output: [1 2 3]

	// Append a value using a pointer
	appendValue(&numbers, 4)

	// Print the updated slice
	fmt.Println("Updated slice:", numbers) // Output: [1 2 3 4]

	// --------------------------------------------------------------------------------------------

	// Declare a map
	data := make(map[string]int)

	// Print the original map
	fmt.Println("Original map:", data) // Output: map[]

	// Add entries using a pointer
	addEntry(&data, "a", 1)
	addEntry(&data, "b", 2)

	// Print the updated map
	fmt.Println("Updated map:", data) // Output: map[a:1 b:2]

	// --------------------------------------------------------------------------------------------

	// Declare an integer variable
	a := 10
	// Declare a pointer to the variable
	p := &a
	// Declare a pointer to the pointer
	pp := &p

	// Print values and addresses
	fmt.Println("Value of x:", a)                // Output: 10
	fmt.Println("Value pointed to by p:", *p)    // Output: 10
	fmt.Println("Value pointed to by pp:", **pp) // Output: 10
	fmt.Println("Address of p:", p)              // Output: Address of x
	fmt.Println("Address of pp:", pp)            // Output: Address of p

	// Assign a new value to the variable through the pointer-to-pointer
	**pp = 20

	// Print values after assignment through pp
	fmt.Println("\nAfter assigning a new value through pp:")
	fmt.Println("Value of a:", a)                // Output: 20
	fmt.Println("Value pointed to by p:", *p)    // Output: 20
	fmt.Println("Value pointed to by pp:", **pp) // Output: 20

	// Directly assign a new value to p
	*p = 30

	// Print values after direct assignment to p
	fmt.Println("\nAfter directly assigning a new value to p:")
	fmt.Println("Value of a:", a)                // Output: 30
	fmt.Println("Value pointed to by p:", *p)    // Output: 30
	fmt.Println("Value pointed to by pp:", **pp) // Output: 30

}
