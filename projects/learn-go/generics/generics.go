package generics

// This package demonstrates the use of generics in Go.

func iterateOverSlice[T any](value []T) []T {
	for _, v := range value {
		// Perform some operation on each element
		// For demonstration, we will just print the value
		println(v)
	}
	return value
}
