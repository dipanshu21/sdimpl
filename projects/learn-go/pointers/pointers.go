package main

import (
	"fmt"
)

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", x)
	var squaredArr [5]int = squareArray(x)
	fmt.Println("Original array after calling square:", x)
	fmt.Println("Squared array:", squaredArr)

	var y [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Original array with pointer:", y)
	squareArrayWithPointer(&y)
	fmt.Println("Original array after calling square:", y)

}

func simplePointerExample() {
	var p *int32 = new(int32)
	*p = 42                           // Assign a value to the pointer
	fmt.Println("Pointer value:", *p) // Dereference the pointer to get the value
	// You can also use the pointer to modify the value
	*p = 100
	fmt.Println("Modified pointer value:", *p) // Dereference again to see the modified value
	// You can also create a pointer to a variable
	var x int32 = 10
	var px *int32 = &x                          // Create a pointer to x
	fmt.Println("Value of x:", x)               // Print the value of x
	fmt.Println("Pointer to x:", px)            // Print the pointer to x
	fmt.Println("Value pointed to by px:", *px) // Dereference px to get the value of x
}

func squareArray(arr [5]int) [5]int {
	for i := range arr {
		arr[i] *= arr[i] // Square each element in the array
	}

	return arr // Return the modified array
}

func squareArrayWithPointer(arr *[5]int) {
	for i := range *arr {
		(*arr)[i] *= (*arr)[i] // Square each element in the array using pointer dereferencing
	}
	// No need to return, as the original slice is modified
}
