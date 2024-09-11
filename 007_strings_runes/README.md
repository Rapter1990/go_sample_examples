# Go Sample Example - Strings and Runes

This repository demonstrates how to work with strings and runes in Go, covering basic string operations, substring extraction, concatenation, string manipulation (upper/lower case), comparison, splitting and joining strings, and iterating over runes. Additionally, it shows how to work with string encoding and decoding using UTF-8.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers fundamental operations with strings and runes in Go.</li>
  <li>It includes examples for basic string manipulation, substring operations, string comparison, rune handling, and encoding/decoding of strings and runes.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	fmt.Println("-----------------------------------------------------------------------------------")

	// Basic string creation
	str := "Hello, Go!"

	// String length
	fmt.Println("Length of string:", len(str)) // Output: 10

	// Substring
	substr := str[7:10]
	fmt.Println("Substring:", substr) // Output: Go!

	// Concatenation
	concatStr := str + " Welcome to Go!"
	fmt.Println("Concatenated string:", concatStr) // Output: Hello, Go! Welcome to Go!

	// String to uppercase
	upperStr := strings.ToUpper(str)
	fmt.Println("Uppercase string:", upperStr) // Output: HELLO, GO!

	// String to lowercase
	lowerStr := strings.ToLower(str)
	fmt.Println("Lowercase string:", lowerStr) // Output: hello, go!

	fmt.Println("-----------------------------------------------------------------------------------")

	// String Comparison
	str1 := "Hello"
	str2 := "World"

	// Equality check
	fmt.Println("str1 == str2:", str1 == str2) // Output: false

	// Compare strings
	fmt.Println("str1 < str2:", str1 < str2) // Output: true (lexicographical comparison)

	fmt.Println("-----------------------------------------------------------------------------------")

	// String Splitting and Joining
	strOneTwoThree := "one,two,three"

	// Split a string
	parts := strings.Split(strOneTwoThree, ",")
	fmt.Println("Splitted parts:", parts) // Output: [one two three]

	// Join a slice of strings
	joinedStr := strings.Join(parts, " | ")
	fmt.Println("Joined string:", joinedStr) // Output: one | two | three

	fmt.Println("-----------------------------------------------------------------------------------")

	strGo := "Hello, Go!"

	// Check if a substring is present
	fmt.Println("Contains 'Go':", strings.Contains(strGo, "Go")) // Output: true

	// Find the index of a substring
	index := strings.Index(strGo, "Go")
	fmt.Println("Index of 'Go':", index) // Output: 7

	// Replace a substring
	replacedStr := strings.ReplaceAll(strGo, "Go", "Golang")
	fmt.Println("Replaced string:", replacedStr) // Output: Hello, Golang!

	fmt.Println("-----------------------------------------------------------------------------------")

	// Rune Usage
	strRuneUsage := "Hello, 世界" // "世界" is "world" in Chinese

	// Iterate over runes
	fmt.Println("Runes:")
	for i, r := range strRuneUsage {
		fmt.Printf("Index %d: %c\n", i, r)
	}

	// Length in bytes
	fmt.Println("Length in bytes:", len(strRuneUsage)) // Output: 13 (because Chinese characters are 3 bytes each)

	// Length in runes
	fmt.Println("Length in runes:", utf8.RuneCountInString(strRuneUsage)) // Output: 9

	// Accessing specific rune
	firstRune := []rune(strRuneUsage)[0]
	fmt.Printf("First rune: %c\n", firstRune) // Output: H

	fmt.Println("-----------------------------------------------------------------------------------")

	// Rune Conversion
	// Convert a character to rune
	char := 'a'
	fmt.Printf("Rune of '%c': %d\n", char, char) // Output: Rune of 'a': 97

	// Convert a rune to string
	runeToStr := string(char)
	fmt.Println("String from rune:", runeToStr) // Output: a

	fmt.Println("-----------------------------------------------------------------------------------")

	// String and Rune Encoding/Decoding
	strEncodeDecode := "Hello, 🌏"

	// Encode string to bytes
	bytes := []byte(strEncodeDecode)
	fmt.Println("Bytes:", bytes) // Output: [72 101 108 108 111 44 32 240 159 144 141]

	// Decode bytes to string
	decodedStr := string(bytes)
	fmt.Println("Decoded string:", decodedStr) // Output: Hello, 🌏

	// Print UTF-8 encoded values of runes
	fmt.Println("UTF-8 encoded values:")
	for i, r := range str {
		fmt.Printf("Index %d: %c (%X)\n", i, r, r)
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

3. Navigate to the `007_strings_runes` directory:

```bash
   cd go_sample_examples/007_strings_runes
```

4. Run the Go program:

```bash
   go run main.go
```

### 📦 Output

When you run the program, you should see the following output:

```
-----------------------------------------------------------------------------------
Length of string: 10
Substring: Go!
Concatenated string: Hello, Go! Welcome to Go!
Uppercase string: HELLO, GO!
Lowercase string: hello, go!
--------------------------------------------------------------------------------
---
str1 == str2: false
str1 < str2: true
--------------------------------------------------------------------------------
---
Splitted parts: [one two three]
Joined string: one | two | three
--------------------------------------------------------------------------------
---
Contains 'Go': true
Index of 'Go': 7
Replaced string: Hello, Golang!
--------------------------------------------------------------------------------
---
Runes:
Index 0: H
Index 1: e
Index 2: l
Index 3: l
Index 4: o
Index 5: ,
Index 6:
Index 7: 世
Index 10: 界
Length in bytes: 13
Length in runes: 9
First rune: H
--------------------------------------------------------------------------------
---
Rune of 'a': 97
String from rune: a
--------------------------------------------------------------------------------
---
Bytes: [72 101 108 108 111 44 32 240 159 140 143]
Decoded string: Hello, 🌏
UTF-8 encoded values:
Index 0: H (48)
Index 1: e (65)
Index 2: l (6C)
Index 3: l (6C)
Index 4: o (6F)
Index 5: , (2C)
Index 6:   (20)
Index 7: G (47)
Index 8: o (6F)
Index 9: ! (21)
```