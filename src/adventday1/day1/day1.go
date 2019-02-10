package day1

import (
	"strconv"

	"../input"
)

func CalculateFrequency() int {
	frequencies := input.GetFrequencies()
	summedFrequency := 0

	for i := 0; i < len(frequencies); i++ {
		frequency := parseStringToInt(frequencies[i])
		summedFrequency += frequency
	}
	return summedFrequency
}

func GetFirstFrequencyDuplicate() int {
	frequencies := input.GetFrequencies()
	summedFrequency := 0
	duplicateFound := false
	var summedFrequencies = []int{}

	for !duplicateFound {
		for i := 0; i < len(frequencies); i++ {
			frequencyInt := parseStringToInt(frequencies[i])
			summedFrequency += frequencyInt
			summedFrequencies = append(summedFrequencies, summedFrequency)

			if containsDuplicates(summedFrequencies, summedFrequency) {
				duplicateFound = true
				break
			}
		}
	}
	return summedFrequency
}

func parseStringToInt(inputString string) int {
	inputInt, _ := strconv.Atoi(inputString)
	return inputInt
}

func containsDuplicates(s []int, e int) bool {
	count := 0
	for _, a := range s {
		if a == e {
			count++
		}
		if count > 1 {
			return true
		}
	}
	return false
}
