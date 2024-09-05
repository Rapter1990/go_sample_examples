package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by zero", a)
	}

	// Returning nil when there's no error
	return a / b, nil // nil -> there is no value // nothing here.
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

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic Error Creation

	// Creating an error
	err := errors.New("something went wrong")

	// Checking and handling the error
	if err != nil {
		fmt.Println("Error occurred:", err)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Custom Error Messages

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	// Case 1: Both values are provided
	a, b := 10, 2
	result, err = divideNil(&a, &b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result) // Output: Result: 5
	}

	// Case 2: One value is nil
	var c *int
	result, err = divideNil(&a, c)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: one or both operands are nil
	} else {
		fmt.Println("Result:", result)
	}

	// Case 3: Both values are nil
	result, err = divideNil(nil, nil)
	if err != nil {
		fmt.Println("Error:", err) // Output: Error: one or both operands are nil
	} else {
		fmt.Println("Result:", result)
	}

	a, b = 10, 0
	result, err = divideNil(&a, &b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Custom Error Types

	err1 := riskyOperation()
	if err1 != nil {
		fmt.Println("Custom error occurred:", err1)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Wrapping Errors

	_, err2 := findItem(0)
	if errors.Is(err2, ErrNotFound) {
		fmt.Println("Specific error:", err2)
	} else if err != nil {
		fmt.Println("General error:", err2)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Unwrapping Errors

	_, err3 := findItem(0)
	if err3 != nil {
		fmt.Println("Error:", err3)

		// Unwrapping the error
		if unwrappedErr := errors.Unwrap(err3); unwrappedErr != nil {
			fmt.Println("Unwrapped error:", unwrappedErr)
		}
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Checking for Specific Errors

	err4 := fmt.Errorf("operation failed: %w", ErrNotFound)

	if errors.Is(err4, ErrNotFound) {
		fmt.Println("The error is ErrNotFound")
	} else {
		fmt.Println("Unknown error:", err4)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Asserting Custom Error Types

	err5 := &MyError{Code: 404, Message: "Not Found"}

	// Wrapping the error
	wrappedErr := fmt.Errorf("something went wrong: %w", err5)

	var myErr *MyError
	if errors.As(wrappedErr, &myErr) {
		fmt.Printf("Custom error detected: %v (Code: %d)\n", myErr.Message, myErr.Code)
	} else {
		fmt.Println("Different error:", wrappedErr)
	}

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println("-----------------------------------------------------------------------------------")

	// Combining Errors

	err6 := process()
	if err != nil {
		fmt.Println("Error chain:", err6)
	}

	fmt.Println("-----------------------------------------------------------------------------------")

}
