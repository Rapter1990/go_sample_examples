package main

import (
	"fmt"
)

func main() {
	// For loop - standard for loop
	for i := 0; i < 5; i++ {
		fmt.Println("For loop iteration:", i)
	}

	// For loop as a while loop
	j := 0
	for j < 5 {
		fmt.Println("While loop iteration:", j)
		j++
	}

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

	// If-else statement
	num := 7
	if num%2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// If with initialization statement
	if n := 10; n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}

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

	// Switch with multiple cases
	letter := 'a'
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("The letter", string(letter), "is a vowel")
	default:
		fmt.Println("The letter", string(letter), "is a consonant")
	}

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
}
