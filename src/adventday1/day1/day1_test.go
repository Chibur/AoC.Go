package day1

import (
	"testing"
)

func TestFrequencyCalculation(t *testing.T) {
	in := getFrequencies()
	want := 8

	got := CalculateFrequency(in)

	if want != got {
		t.Errorf("Failed to calculate correct frequency. Expected: %d, Actual: %d", want, got)
	}
}

func TestGetFirstFrequencyDuplicate(t *testing.T) {
	in := getFrequencies()
	want := 1

	got := GetFirstFrequencyDuplicate(in)

	if want != got {
		t.Errorf("Failed to find correct duplicate frequency. Expected: %d, Actual: %d", want, got)
	}
}

func getFrequencies() []string {
	return []string{"+2", "-1", "+3", "+1", "-4", "4", "-2", "2", "3"}
}
