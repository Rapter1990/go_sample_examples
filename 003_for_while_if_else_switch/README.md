# Go Sample Example - For, While, If-Else, Switch

This repository contains a Go (Golang) program that demonstrates various control flow structures including `for` loops, `if-else` statements, and `switch` statements. The example showcases how to use these structures for looping and conditional execution.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example introduces different types of loops in Go, including `for` loops used as standard loops and while-like loops.</li>
  <li>It covers the use of infinite loops with `break` statements.</li>
  <li>It demonstrates `if-else` statements with and without initialization statements, as well as nested conditions.</li>
  <li>It shows the use of `switch` statements with multiple cases, and a switch without an expression for alternative conditional logic.</li>
</ul>

## 💻 Code Example

```golang
package main

import (
	"fmt"
)

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// For loop - standard for loop
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// For loop as a while loop
	j := 0
	for j < 5 {
		fmt.Println("While loop iteration:", j)
		j++
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Infinite loop with break
	k := 0
	for {
		if k == 3 {
			fmt.Println("Breaking the infinite loop at iteration:", k)
			break
		}
		fmt.Println("Infinite loop iteration:", k)
		k++
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// If-else statement
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// If with initialization statement
	if n := 10; n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Nested if-else
	age := 25
	if age < 13 {
		fmt.Println("Child")
	} else if age < 20 {
		fmt.Println("Teenager")
	} else if age < 30 {
		fmt.Println("Young Adult")
	} else {
		fmt.Println("Adult")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Switch statement
	day := "Tuesday"
	switch day {
	case "Monday":
		fmt.Println("Today is Monday")
	case "Tuesday":
		fmt.Println("Today is Tuesday")
	case "Wednesday":
		fmt.Println("Today is Wednesday")
	default:
		fmt.Println("Today is another day")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Switch with multiple cases
	letter := 'a'
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("The letter", string(letter), "is a vowel")
	default:
		fmt.Println("The letter", string(letter), "is a consonant")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

	// Switch without an expression (alternative to if-else)
	score := 85
	switch {
	case score >= 90:
		fmt.Println("Grade: A")
	case score >= 80:
		fmt.Println("Grade: B")
	case score >= 70:
		fmt.Println("Grade: C")
	default:
		fmt.Println("Grade: F")
	}

	fmt.Println("-----------------------------------------------------------------------------------")

}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:
   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```
3. Navigate to the `003_for_while_if_else_switch` directory:
   ```bash
   cd go_sample_examples/003_for_while_if_else_switch
   ```
4. Run the Go program:
   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
For loop iteration: 0
For loop iteration: 1
For loop iteration: 2
For loop iteration: 3
For loop iteration: 4
-----------------------------------------------------------------------------------
While loop iteration: 0
While loop iteration: 1
While loop iteration: 2
While loop iteration: 3
While loop iteration: 4
-----------------------------------------------------------------------------------
Infinite loop iteration: 0
Infinite loop iteration: 1
Infinite loop iteration: 2
Breaking the infinite loop at iteration: 3
-----------------------------------------------------------------------------------
7 is odd
-----------------------------------------------------------------------------------
10 is even
-----------------------------------------------------------------------------------
Young Adult
-----------------------------------------------------------------------------------
Today is Tuesday
-----------------------------------------------------------------------------------
The letter a is a vowel
-----------------------------------------------------------------------------------
Grade: B
-----------------------------------------------------------------------------------
```