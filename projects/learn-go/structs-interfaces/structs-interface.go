package main

import (
	"fmt"
)

const (
	// Size defines the number of elements to insert
	Size = 1000000
)

// BenchmarkResult holds the timing results for a benchmark
type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type electricEngine struct {
	mpkwh     uint8
	kwh       uint8
	ownerInfo owner
}

func (engine gasEngine) milesLeft() uint16 {
	return uint16(engine.mpg * engine.gallons)
}

func (engine electricEngine) milesLeft() uint16 {
	return uint16(engine.kwh * engine.mpkwh)
}

type engine interface {
	milesLeft() uint16
}

func canMakeIt(e engine, distance uint16) bool {
	if distance <= e.milesLeft() {
		fmt.Println("You can make it!")
		return true
	} else {
		fmt.Println("You cannot make it!")
		return false
	}
}

type owner struct {
	name string
}

func main() {

	// Create an instance of gasEngine
	engine1 := gasEngine{
		mpg:     25,
		gallons: 10,
		ownerInfo: owner{
			name: "John Doe",
		},
	}

	engine2 := electricEngine{
		mpkwh: 4,
		kwh:   50,
		ownerInfo: owner{
			name: "Jane Smith",
		},
	}

	printEngineDetails(engine1)
	printEngineDetails(engine2)

	canMakeIt(engine1, 300)
	canMakeIt(engine2, 100)
}

func printEngineDetails(e engine) {
	fmt.Printf("Engine Details:\n")
	fmt.Printf("Miles Left: %d\n", e.milesLeft())
}
