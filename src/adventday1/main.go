package main

import (
	"fmt"

	"./day1"
)

func main() {
	frequencySum := day1.CalculateFrequency()
	duplicateFrequency := day1.GetFirstFrequencyDuplicate()
	fmt.Println("Frequency Sum: ", frequencySum)
	fmt.Println("Duplicate Frequency: ", duplicateFrequency)
}
