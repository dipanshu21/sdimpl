package main

import (
	"fmt"
	"strings"
)

func main() {
	//its tricky to iterate over chars in a string in Go, as
	// strings are not directly iterable like in some other languages.
	str := "Hello Ë, World!"
	for i := 0; i < len(str); i++ {
		fmt.Printf("Character at index %d: %c\n", i, str[i])
	}

	//as E is a multi-byte character, it will not be printed correctly
	//using range to iterate over the string will handle multi-byte characters correctly
	for i, char := range str {
		fmt.Printf("Character at index %d: %c\n", i, char)
	}

	//using rune slice to iterate over the string
	var str2 = []rune("Hello Ë, World!")
	for i, char := range str2 {
		fmt.Printf("Character at index %d: %c\n", i, char)
	}

	//String building in Go is done using the strings.Builder type
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(" World!")
	fmt.Println("Built String:", builder.String())
}
