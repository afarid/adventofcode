package task1

import (
	"testing"
)

func TestGetComplement2(t *testing.T) {
	tables := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{500, 900, 400, 20, 171}, 0},
		{[]int{1000, 100, 900, 1020}, 1000 *1020  },
	}

	for _, table := range tables {
		result := GetComplement2(table.input, 2020)
		expectedResult := table.expectedResult
		if result != expectedResult {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
		}
	}

}

func TestGetComplement3(t *testing.T) {
	tables := []struct {
		input          []int
		expectedResult int
	}{
		{[]int{500, 900, 400, 20, 171}, 0},
		{[]int{1000, 100, 900, 1020, 500, 520}, 1000 * 500 * 520  },
	}
	for _, table := range tables{
		result := GetComplement3(table.input, 2020)
		expectedResult := table.expectedResult
		if result != expectedResult {
			t.Errorf("Result was incorrect, got: %d, want: %d.", result, expectedResult)
		}
	}

}
