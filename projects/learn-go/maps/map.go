package main

import "fmt"

func main() {

	var myMap = map[string]int{"one": 1, "two": 2, "three": 3} // Declare and initialize a map
	fmt.Println("Map Length:", len(myMap))
	fmt.Println("Value for key 'one':", myMap["one"])
	fmt.Println("Value for key 'two':", myMap["two"])
	fmt.Println("Value for key 'three':", myMap["three"])

	var val, exists = myMap["four"] // Check if a key exists in the map
	if exists {
		fmt.Println("Value for key 'four':", val)
	} else {
		fmt.Println("Key 'four' does not exist in the map.")
	}

	//iterate over the map
	// The order of iteration is not guaranteed to be the same every time
	fmt.Println("Iterating over the map:")
	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	//while loop in go is not a separate construct like in some other languages.
	// Instead, you can use a for loop with a condition to achieve the same effect.
	var i int = 0

	for i < 5 {
		fmt.Println("While Loop Iteration:", i)
		i++
	}

	/**
	 Another way to do this is
	 for {
		if i >= 5 {
			break
		}
		fmt.Println("While Loop Iteration:", i)
		i++
	}
	*/
}
