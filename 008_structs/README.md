# Go Sample Example - Structs

This repository demonstrates various types of structs in Go, including basic struct usage, anonymous structs, pointers to structs, methods on structs, struct embedding, and struct tags. It showcases how to define and use structs to represent and manage data in Go programs.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic struct definitions, initialization, and usage in Go.</li>
  <li>It includes examples of methods on structs, pointers to structs, embedded structs (composition), and struct tags for JSON serialization.</li>
</ul>

## 💻 Code Example

```go
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
	person1 := Person{Name: "John Doe", Age: 30}
	fmt.Println("Person 1:", person1)

	person2 := Person{"Jane Doe", 25}
	fmt.Println("Person 2:", person2)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Pointers to Structs
	car := &Car{Brand: "Toyota", Year: 2021}
	fmt.Println("Car Brand:", car.Brand)
	car.Year = 2022
	fmt.Println("Updated Car Year:", car.Year)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Structs with Methods
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println("Area:", rect.Area())
	rect.Resize(20, 10)
	fmt.Println("New Area:", rect.Area())

	fmt.Println("-----------------------------------------------------------------------------------")

	// Embedded Structs (Composition)
	emp := Employee{Name: "John Doe", Age: 30, Address: Address{City: "New York", State: "NY", Country: "USA"}}
	fmt.Println("Employee Name:", emp.Name, "City:", emp.City)

	fmt.Println("-----------------------------------------------------------------------------------")

	// Struct Tags
	person := Person1{Name: "John Doe", Age: 30}
	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))

	employee := Employee1{FirstName: "Jane", LastName: "Doe", Age: 28, Position: "Developer"}
	jsonData1, _ := json.Marshal(employee)
	fmt.Println(string(jsonData1))

	user := User{Username: "johndoe", Password: "secret", Email: "johndoe@example.com"}
	jsonData2, _ := json.Marshal(user)
	fmt.Println(string(jsonData2))

	fmt.Println("-----------------------------------------------------------------------------------")
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `008_structs` directory:

```bash
cd go_sample_examples/008_structs
```

4. Run the Go program:

```bash
go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```
Person 1: {John Doe 30}
Person 2: {Jane Doe 25}
-----------------------------------------------------------------------------------
Car Brand: Toyota
Updated Car Year: 2022
-----------------------------------------------------------------------------------
Area: 50
New Area: 200
-----------------------------------------------------------------------------------
Employee Name: John Doe City: New York
-----------------------------------------------------------------------------------
{"name":"John Doe","age":30}
{"first_name":"Jane","last_name":"Doe","age":28,"position":"Developer"}
{"username":"johndoe","email":"johndoe@example.com"}
-----------------------------------------------------------------------------------
```