package main

import (
	"errors"
	"fmt"
)

func main() {
	// This is a placeholder for the func-control program.
	// You can implement your logic here or call other functions as needed.
	fmt.Println("Function Example")

	// Call the printMe function to demonstrate its functionality
	printMe()

	// Call the printMe2 function with a string argument
	printMe2("Hello, Go Functions!")

	printIntDivAndRemainder(10, 2)
	printIntDivAndRemainder(10, 0)
	printIntDivAndRemainder(13, 3)

	printIntDivAndRemainderCaseVersion(10, 2)
	printIntDivAndRemainderCaseVersion(10, 0)
	printIntDivAndRemainderCaseVersion(13, 3)
}

func printMe() {
	fmt.Println("This is a function that prints a message.")
	// You can add more functionality here as needed.
}

func printMe2(value string) {
	fmt.Println("Value passed to printMe2:", value)
	// You can add more functionality here as needed.
}

func printIntDivAndRemainder(a int, b int) {
	result, remainder, err := intDivisionAndRemainder(a, b)

	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	} else if remainder == 0 {
		fmt.Println("The result is an integer.")
	} else {
		fmt.Println("The result is a float.")
		fmt.Println("Integer Division Remainder:", remainder)
	}

	fmt.Println("Integer Division Result:", result)
}

func printIntDivAndRemainderCaseVersion(a int, b int) {
	result, remainder, err := intDivisionAndRemainder(a, b)

	switch {
	case err != nil:
		fmt.Println("Error:", err.Error())
		return
	case remainder == 0:
		fmt.Println("The result is an integer.")
	default:
		fmt.Println("The result is a float.")
		fmt.Println("Integer Division Remainder:", remainder)
	}
	fmt.Println("Integer Division Result:", result)
}

func intDivisionAndRemainder(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("division by zero")
		return 0, 0, err
	}
	return a / b, a % b, err
}
