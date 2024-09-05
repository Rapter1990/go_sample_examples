package main

import (
	"encoding/json"
	"fmt"
)

// Define a struct type
type Person struct {
	Name string
	Age  int
}

// Define a struct type
type Car struct {
	Brand string
	Year  int
}

// Define a struct type
type Rectangle struct {
	Width  int
	Height int
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

// Method to change the dimensions of the rectangle (by pointer)
func (r *Rectangle) Resize(width, height int) {
	r.Width = width
	r.Height = height
}

// Define a struct type
type Address struct {
	City    string
	State   string
	Country string
}

// Define a struct that embeds another struct
type Employee struct {
	Name    string
	Age     int
	Address // Embedded struct
}

// Define a struct with JSON tags
type Person1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city,omitempty"` // The `omitempty` tag omits the field if it's empty or zero-valued
}

// Define a struct with custom JSON field names
type Employee1 struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Position  string `json:"position"`
}

// Define a struct with a field to be ignored in JSON serialization
type User struct {
	Username string `json:"username"`
	Password string `json:"-"` // The `-` tag tells the JSON encoder to ignore this field
	Email    string `json:"email"`
}

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Struct Declaration and Initialization

	// Create an instance of the struct and initialize fields
	var person1 Person
	person1.Name = "John Doe"
	person1.Age = 30

	// Print the struct fields
	fmt.Println("Person 1:", person1)           // Output: {John Doe 30}
	fmt.Println("Person 1 Name:", person1.Name) // Output: John Doe
	fmt.Println("Person 1 Age:", person1.Age)   // Output: 30

	// Initialize struct with a literal
	person2 := Person{Name: "Jane Doe", Age: 25}
	fmt.Println("Person 2:", person2) // Output: {Jane Doe 25}

	// Initialize struct with zero values
	person3 := Person{}
	fmt.Println("Person 3:", person3) // Output: { 0}

	// Initialize struct without specifying field names
	person4 := Person{"Alice", 28}
	fmt.Println("Person 4:", person4) // Output: {Alice 28}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Anonymous Structs

	// Declare and initialize an anonymous struct
	student := struct {
		Name  string
		Grade int
	}{
		Name:  "Tom",
		Grade: 90,
	}

	// Print the anonymous struct
	fmt.Println("Student:", student)             // Output: {Tom 90}
	fmt.Println("Student Name:", student.Name)   // Output: Tom
	fmt.Println("Student Grade:", student.Grade) // Output: 90

	fmt.Println("-----------------------------------------------------------------------------------")

	// Pointers to Structs

	// Create a pointer to a struct
	car := &Car{Brand: "Toyota", Year: 2021}

	// Access struct fields through the pointer
	fmt.Println("Car Brand:", car.Brand) // Output: Toyota
	fmt.Println("Car Year:", car.Year)   // Output: 2021

	// Modify struct fields through the pointer
	car.Year = 2022
	fmt.Println("Updated Car Year:", car.Year) // Output: 2022

	fmt.Println("-----------------------------------------------------------------------------------")

	// Structs with Methods

	// Initialize a Rectangle struct
	rect := Rectangle{Width: 10, Height: 5}

	// Call the Area method
	area := rect.Area()
	fmt.Println("Area:", area) // Output: 50

	// Resize the rectangle using the Resize method
	rect.Resize(20, 10)
	fmt.Println("Resized Rectangle:", rect) // Output: {20 10}

	// Call the Area method again after resizing
	area = rect.Area()
	fmt.Println("New Area:", area) // Output: 200

	fmt.Println("-----------------------------------------------------------------------------------")

	// Embedded Structs (Composition)

	// Initialize an Employee struct with embedded Address
	emp := Employee{
		Name:    "John Doe",
		Age:     30,
		Address: Address{City: "New York", State: "NY", Country: "USA"},
	}

	// Access fields of the embedded struct directly
	fmt.Println("Employee Name:", emp.Name)       // Output: John Doe
	fmt.Println("Employee City:", emp.City)       // Output: New York
	fmt.Println("Employee State:", emp.State)     // Output: NY
	fmt.Println("Employee Country:", emp.Country) // Output: USA

	fmt.Println("-----------------------------------------------------------------------------------")

	// Struct Tags

	// Create an instance of the struct
	person := Person1{
		Name: "John Doe",
		Age:  30,
	}

	// Convert struct to JSON
	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData)) // Output: {"name":"John Doe","age":30}

	// Create an instance of the struct
	employee := Employee1{
		FirstName: "Jane",
		LastName:  "Doe",
		Age:       28,
		Position:  "Developer",
	}

	// Convert struct to JSON
	jsonData1, _ := json.Marshal(employee)
	fmt.Println(string(jsonData1)) // Output: {"first_name":"Jane","last_name":"Doe","age":28,"position":"Developer"}

	// Create an instance of the struct
	user := User{
		Username: "johndoe",
		Password: "secret", // This field will be ignored in the JSON output
		Email:    "johndoe@example.com",
	}

	// Convert struct to JSON
	jsonData2, _ := json.Marshal(user)
	fmt.Println(string(jsonData2)) // Output: {"username":"johndoe","email":"johndoe@example.com"}

	fmt.Println("-----------------------------------------------------------------------------------")

}
