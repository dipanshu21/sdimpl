package main

import "fmt"

func main() {
	// This is a placeholder for the arrays program.
	// You can implement your logic here or call other functions as needed.

	// Example of declaring and using an array in Go
	var myArray [5]int8 // Declare an array of integers with a fixed size of 5
	myArray[0] = 10     // Assign a value to the first element
	myArray[1] = 20     // Assign a value to the second element
	myArray[2] = 30     // Assign a value to the third element
	myArray[3] = 40     // Assign a value to the fourth element
	myArray[4] = 50     // Assign a value to the fifth element

	println("Array Length:", len(myArray))
	println("First Element:", myArray[0])
	println("First two Element:", myArray[1:3])

	//print memory address of the array and its elements
	println("Memory Address of Array:", &myArray)
	println("Memory Address of First Element:", &myArray[0])
	println("Memory Address of Second Element:", &myArray[1])
	println("Memory Address of Third Element:", &myArray[2])
	println("Memory Address of Fourth Element:", &myArray[3])

	// Print the array elements
	for i, value := range myArray {
		println("Element at index", i, "is", value)
	}

	//different ways to declare and initialize an array
	var anotherArray = [3]int{1, 2, 3} // Declare and initialize an array with values
	println("Another Array Length:", len(anotherArray))

	var yetAnotherArray = [...]int{4, 5, 6} // Declare and initialize an array with values using ellipsis
	println("Yet Another Array Length:", len(yetAnotherArray))

	// Slices are more flexible than arrays and can grow or shrink in size
	//Basically they are a wrapper around arrays
	var mySlice []int = []int{1, 2, 3, 4, 5} // Declare and initialize a slice
	println("Slice Length:", len(mySlice))
	println("First Element of Slice:", mySlice[0])
	fmt.Println("First two Elements of Slice:", mySlice[1:3])

	//add elements to a slice
	mySlice = append(mySlice, 6) // Append an element to the slice
	fmt.Println("Slice after appending an element:", mySlice)

	var anotherSlice = make([]int, 3) // Create a slice with a length of 3
	anotherSlice[0] = 10
	anotherSlice[1] = 20
	anotherSlice[2] = 30
	fmt.Println("First two Elements of Another Slice:", anotherSlice[1:3])

	//Slices are basically dynamically grown arrays
	var dynamicSlice = []int{4, 5, 6} // Declare a slice without initializing it
	fmt.Printf("Dynamic Slice Length: %v and Capacity is: %v", len(dynamicSlice), cap(dynamicSlice))
	dynamicSlice = append(dynamicSlice, 7)
	fmt.Println("Dynamic Slice after appending elements:", dynamicSlice)
	fmt.Printf("Dynamic Slice Length: %v and Capacity is: %v", len(dynamicSlice), cap(dynamicSlice))
}
