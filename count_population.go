package main

import (
	"fmt"

	"github.com/gennadyterekhov/tribal-relations-core-abstraction-layer/count_population"
)

func main() {
	fertility := 0
	currentPopulation := 0
	availableFood := 0

	_, err := fmt.Scanf("What is the number on the die: %d", fertility)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	_, err = fmt.Scanf("What is the number on available food: %d", availableFood)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	_, err = fmt.Scanf("What is the current population: %d", currentPopulation)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	newPopulation := count_population.CountPopulation(fertility, availableFood, currentPopulation)

	fmt.Println(newPopulation)
}
