package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestComparisonSort(t *testing.T) {
	tests := []struct {
		input          []int32
		expectedResult []int32
	}{
		{
			input:          []int32{503, 87, 512, 61, 908, 170, 897, 275, 897, 653, 426, 154, 509, 612, 677, 765, 703},
			expectedResult: []int32{61, 87, 154, 170, 275, 426, 503, 509, 512, 612, 653, 677, 703, 765, 897, 897, 908},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-case-%d", i), func(t *testing.T) {
			sortedListTable := countingSort(test.input)
			result := make([]int32, len(test.input))

			for i, pos := range sortedListTable {
				result[pos] = test.input[i]
			}

			require.Equal(t, test.expectedResult, result)
		})
	}
}

func TestDistrubutionSort(t *testing.T) {
	tests := []struct {
		u              int
		v              int
		input          []int32
		expectedResult []int32
	}{
		{
			u:              1,
			v:              5,
			input:          []int32{3, 5, 1, 4, 2},
			expectedResult: []int32{1, 2, 3, 4, 5},
		},
		{
			u:              1,
			v:              5,
			input:          []int32{3, 5, 5, 5, 4, 3, 2, 1, 1, 4, 2},
			expectedResult: []int32{0, 1, 0, 2, 0, 3, 0, 4, 0, 0, 5},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-case-%d", i), func(t *testing.T) {
			result := distrubutionSort(test.u, test.v, test.input)

			require.Equal(t, test.expectedResult, result)
		})
	}
}

func TestStraightInsertionSort(t *testing.T) {
	tests := []struct {
		input          []int32
		expectedResult []int32
	}{
		{
			input:          []int32{3, 5, 1, 4, 2},
			expectedResult: []int32{1, 2, 3, 4, 5},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-case-%d", i), func(t *testing.T) {
			result := straightInsertionSort(test.input)

			require.Equal(t, test.expectedResult, result)
		})
	}
}

func TestModdedStraightInsertionSort(t *testing.T) {
	tests := []struct {
		jump           int
		input          []int32
		expectedResult []int32
	}{
		{
			jump:           2,
			input:          []int32{3, 5, 1, 4, 2},
			expectedResult: []int32{1, 4, 2, 5, 3},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-case-%d", i), func(t *testing.T) {
			result := moddedStraightInsertionSort(2, test.input)

			require.Equal(t, test.expectedResult, result)
		})
	}
}

func TestShellSort(t *testing.T) {
	tests := []struct {
		input          []int32
		expectedResult []int32
	}{
		{
			input:          []int32{1, 15, 5, 2, 4, 10, 3, 21},
			expectedResult: []int32{1, 2, 3, 4, 5, 10, 15, 21},
		},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("test-case-%d", i), func(t *testing.T) {
			result := shellSort(test.input)
			require.Equal(t, test.expectedResult, result)
		})
	}
}
