# Go Sample Example - String Manipulations

This repository demonstrates various string manipulation techniques in Go, including checking substrings, counting characters, and transforming strings. It showcases how to perform common string operations using the `strings` package in Go.

## 📖 Information

<ul style="list-style-type:disc">
  <li>This example covers various string manipulation functions available in the Go `strings` package.</li>
  <li>It includes examples for checking substrings, counting occurrences, and replacing characters, as well as splitting and joining strings.</li>
</ul>

## 💻 Code Example

```go
package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	// String Manipulation

	// Contains
	p("Contains 'world' in 'Hello, world!': ", strings.Contains("Hello, world!", "world"))

	// Count
	p("Count occurrences of 'l' in 'Hello, world!': ", strings.Count("Hello, world!", "l"))

	// HasPrefix
	p("HasPrefix 'Hello' in 'Hello, world!': ", strings.HasPrefix("Hello, world!", "Hello"))

	// HasSuffix
	p("HasSuffix '!' in 'Hello, world!': ", strings.HasSuffix("Hello, world!", "!"))

	// Index
	p("Index of 'o' in 'Hello, world!': ", strings.Index("Hello, world!", "o"))
	p("Index of 'z' in 'Hello, world!': ", strings.Index("Hello, world!", "z"))

	// Join
	p("Join ['Hello', 'world'] with ', ': ", strings.Join([]string{"Hello", "world"}, ", "))

	// Repeat
	p("Repeat 'Go' 3 times: ", strings.Repeat("Go", 3))

	// Replace
	p("Replace 'l' with 'L' in 'Hello, world!' (-1 for all): ", strings.Replace("Hello, world!", "l", "L", -1))
	p("Replace 'l' with 'L' in 'Hello, world!' (first occurrence): ", strings.Replace("Hello, world!", "l", "L", 1))

	// ReplaceAll
	p("ReplaceAll 'foo' with 'bar' in 'foofoofoo': ", strings.ReplaceAll("foofoofoo", "foo", "bar"))

	// Split
	p("Split 'a-b-c-d-e' by '-': ", strings.Split("a-b-c-d-e", "-"))
	p("Split 'a,b,c,d,e' by ',': ", strings.Split("a,b,c,d,e", ","))

	// ToLower
	p("ToLower 'HELLO, WORLD!': ", strings.ToLower("HELLO, WORLD!"))

	// ToUpper
	p("ToUpper 'hello, world!': ", strings.ToUpper("hello, world!"))

	// Trim
	p("Trim spaces from '  Hello, world!  ': ", strings.Trim("  Hello, world!  ", " "))

	// TrimPrefix
	p("TrimPrefix 'Hello' from 'Hello, world!': ", strings.TrimPrefix("Hello, world!", "Hello"))

	// TrimSuffix
	p("TrimSuffix '!' from 'Hello, world!': ", strings.TrimSuffix("Hello, world!", "!"))

	// Fields (splits on whitespace)
	p("Fields in 'Hello world  from  Go!': ", strings.Fields("Hello world  from  Go!"))

	// Title
	p("Title case 'go is awesome': ", strings.Title("go is awesome"))

	// Map (custom mapping function)
	p("Map function to replace 'i' with 'I' in 'this is a test': ",
		strings.Map(func(r rune) rune {
			if r == 'i' {
				return 'I'
			}
			return r
		}, "this is a test"))

}
```

### 🏃 How to Run

1. Make sure you have Go installed. If not, you can download it from [here](https://golang.org/dl/).
2. Clone this repository:

   ```bash
   git clone https://github.com/Rapter1990/go_sample_examples.git
   ```

3. Navigate to the `028_string_manipulations` directory:

   ```bash
   cd go_sample_examples/028_string_manipulations
   ```

4. Run the Go program:

   ```bash
   go run main.go
   ```

### 📦 Output

When you run the program, you should see the following output:

```
Contains 'world' in 'Hello, world!':  true
Count occurrences of 'l' in 'Hello, world!':  3                           
HasPrefix 'Hello' in 'Hello, world!':  true                               
HasSuffix '!' in 'Hello, world!':  true                                   
Index of 'o' in 'Hello, world!':  4                                       
Index of 'z' in 'Hello, world!':  -1                                      
Join ['Hello', 'world'] with ', ':  Hello, world                          
Repeat 'Go' 3 times:  GoGoGo                                              
Replace 'l' with 'L' in 'Hello, world!' (-1 for all):  HeLLo, worLd!      
Replace 'l' with 'L' in 'Hello, world!' (first occurrence):  HeLlo, world!
ReplaceAll 'foo' with 'bar' in 'foofoofoo':  barbarbar                    
Split 'a-b-c-d-e' by '-':  [a b c d e]                                    
Split 'a,b,c,d,e' by ',':  [a b c d e]                                    
ToLower 'HELLO, WORLD!':  hello, world!                                   
ToUpper 'hello, world!':  HELLO, WORLD!                                   
Trim spaces from '  Hello, world!  ':  Hello, world!                      
TrimPrefix 'Hello' from 'Hello, world!':  , world!                        
TrimSuffix '!' from 'Hello, world!':  Hello, world                        
Fields in 'Hello world  from  Go!':  [Hello world from Go!]               
Title case 'go is awesome':  Go Is Awesome                                
Map function to replace 'i' with 'I' in 'this is a test':  thIs Is a test
```
