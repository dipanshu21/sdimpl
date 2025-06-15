package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//Available in specific byte sizes
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	// float32, float64
	var intNum int = 32767
	intNum = intNum + 1
	fmt.Println("Integer Number:", intNum)

	var floatNum float32 = 3.14
	fmt.Println("Float Number:", floatNum)

	var floatNum64 float64 = 3.141592653589793
	var intNum1 int64 = 1
	/*
		Not allowed to do math between different types directly.
		var result float64 = intNum64 + floatNum64
		To perform operations between different types, you need to explicitly convert one type to another.
	*/
	var result float64 = float64(intNum1) + floatNum64
	fmt.Println("Result of float64 and int64 addition:", result)

	var intNum2 int = 3
	var intNum3 int = 2
	fmt.Println("Sum of two integers:", intNum2+intNum3)
	fmt.Println("Division of two integers:", intNum2/intNum3)

	var str string = "Hello, Go!"
	fmt.Println("String:", str)

	fmt.Println("String with escape characters:", "Hello, \nGo!")

	/** getting the length of a string is tricky in Go
	 * because it is not a simple count of characters.
	 * It counts the number of bytes in the string.
	 * This is because Go uses UTF-8 encoding, where some characters may take up more than one byte.
	 * To get the number of characters (runes) in a string, you can use the utf8 package.
	 * The utf8.RuneCountInString function counts the number of runes (characters) in a string
	 */
	fmt.Println("Length of string:", utf8.RuneCountInString(str))

	//Runes are used to represent Unicode characters in Go.
	var myRune rune = 'A'
	fmt.Println("Rune:", myRune)

	var myBool bool = true
	fmt.Println("Boolean:", myBool)

	// Can skip the var keyword if you want to use the short variable declaration syntax
	myShortVar := "This is a short variable declaration"
	fmt.Println("Short Variable Declaration:", myShortVar)

	//can define multiple variables at once
	var a, b, c int = 1, 2, 3
	fmt.Println("Multiple Variables:", a, b, c)

	a, b, c = 4, 5, 6 // reassigning values to multiple variables
	fmt.Println("Reassigned Multiple Variables:", a, b, c)

	//Constants are immutable values that cannot be changed after they are defined.
	const myConst = "This is a constant"
	fmt.Println("Constant:", myConst)
}
