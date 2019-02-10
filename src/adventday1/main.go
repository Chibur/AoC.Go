package main

import (
	"fmt"

	"./day1"
	"./input"
)

func main() {
	frequencies := input.GetFrequencies()
	frequencySum := day1.CalculateFrequency(frequencies)
	duplicateFrequency := day1.GetFirstFrequencyDuplicate(frequencies)
	fmt.Println("Frequency Sum: ", frequencySum)
	fmt.Println("Duplicate Frequency: ", duplicateFrequency)
}
