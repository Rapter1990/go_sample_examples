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
