# Go Sample Example - Error Handling

This repository demonstrates various ways to handle errors in Go. It covers the creation, checking, and wrapping of errors, as well as custom error types and handling nil pointers.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers basic error handling concepts in Go.</li>
  <li>It includes examples of creating and handling errors, custom error types, wrapping errors, and working with nil values.</li>
  <li>The example also demonstrates the usage of <code>errors.Is</code>, <code>errors.As</code>, and <code>errors.Unwrap</code> for advanced error handling.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}
	return a / b, nil
}

func divideNil(a, b *int) (int, error) {
	if a == nil || b == nil {
		return 0, fmt.Errorf("one or both operands are nil")
	}
	if *b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", *a)
	}
	return *a / *b, nil
}

// Custom error type
type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func riskyOperation() error {
	return &MyError{
		Code:    404,
		Message: "Resource not found",
	}
}

var ErrNotFound = errors.New("item not found")

func findItem(id int) (string, error) {
	if id == 0 {
		return "", fmt.Errorf("findItem failed: %w", ErrNotFound)
	}
	return "Item", nil
}

var ErrStep1 = errors.New("step 1 failed")
var ErrStep2 = errors.New("step 2 failed")

func process() error {
	if err := step1(); err != nil {
		return fmt.Errorf("process failed: %w", err)
	}
	if err := step2(); err != nil {
		return fmt.Errorf("process failed: %w", err)
	}
	return nil
}

func step1() error {
	return ErrStep1
}

func step2() error {
	return ErrStep2
}

func main() {
	// Example of basic error handling
	err := errors.New("something went wrong")
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	// Example of custom error message
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Handling nil pointer errors
	a, b := 10, 2
	result, err = divideNil(&a, &b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Using custom error types
	err1 := riskyOperation()
	if err1 != nil {
		fmt.Println("Custom error occurred:", err1)
	}

	// Wrapping and unwrapping errors
	_, err2 := findItem(0)
	if errors.Is(err2, ErrNotFound) {
		fmt.Println("Specific error:", err2)
	}

	// Combining multiple errors
	err6 := process()
	if err6 != nil {
		fmt.Println("Error chain:", err6)
	}
}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
```

3. Navigate to the `010_error` directory:

```bash
   cd go_sample_examples/010_error
```

4. Run the Go program:

```bash
   go run main.go
```

### 📦 Output

When you run the program, you should see output similar to this:

```
Error occurred: something went wrong
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Error: cannot divide 10 by zero
Result: 5
Error: one or both operands are nil
Error: one or both operands are nil
Error: cannot divide 10 by zero
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Custom error occurred: Error 404: Resource not found
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Specific error: findItem failed: item not found
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Error: findItem failed: item not found
Unwrapped error: item not found
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
The error is ErrNotFound
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Custom error detected: Not Found (Code: 404)
--------------------------------------------------------------------------------
---
--------------------------------------------------------------------------------
---
Error chain: process failed: step 1 failed
```