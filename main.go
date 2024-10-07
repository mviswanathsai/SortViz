package main

import (
	"math"
)

// Compare each number with every other number.
// The algorithm iterates i and compares it with every element that comes before i.
// This way, we don't do the same comparisons twice.
// The comparison of i with elements that come after i will happen in subsequent iterations.

// Returns an auxilary table which the sorted order of the input array.
func countingSort(input []int32) []int32 {
	count := make([]int32, len(input))

	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[i] < input[j] {
				count[j] += 1
			} else {
				count[i] += 1
			}
		}
	}

	return count
}

// In my opinion this is just a counting sort where you have a few restrictions that help
// make it more efficient.
// Essentially, we consider that all the keys lie between a certain range (smaller the range, faster the sort)
// This works really well if there is a lot of equality as well.
// We build a frequncy table and organize the keys into "families" and then generate a result table
// with only the last member of each family populated in it.
// Though it seems really restrictive, you may apply this logic to just the leading values of the keys
// and obtain a partially sorted table which is pretty neat.
// Also, unlike the auxilary table you got from the countingSort, here we get the actual result.
func distrubutionSort(u int, v int, input []int32) []int32 {
	count := make([]int32, v-u+1)
	result := make([]int32, len(input))

	// 0 would be u, 1 would be u+1 and so on.
	// it is essentially, u, u+i...v
	for _, value := range input {
		count[value-int32(u)] += 1
	}

	// now count is a frequency table.
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	// now count[i] is the number of keys that are less than, or equal to u+i
	// in the case where there are no equalities, it would be a perfect sort
	// we will loop through the frequency table in reverse order and generate a result
	// table where we populate only the positions of the last members of a given family.
	// for example, if the result is supposed to be 1, 1, 1, 2, 2, 2, 2, 2
	// the result table we generate would look like 0, 0, 1, 0, 0, 0, 0, 2
	for j := len(count); j > 0; j-- {
		i := count[j-1]
		result[i-1] = int32(j)
		count[j-1] -= 1
	}

	return result
}

// This sort assumes that everything to the left of the element under consideration, Kj is already sorted.
// We look for where we can place Kj, for this we go from the left of Kj and try to find a Ki such that
// Ki<Kj.
// The first instance where this is true, is where Kj is to be placed.
// Until we find that Ki<=Kj, we keep shifting Ki with Ki+1(Kj)
// Note, due to the nature of the algorithm, input[i+1] will always be the element under consideration Kj.
func straightInsertionSort(input []int32) []int32 {
	for j := 1; j < len(input); j++ {
		i := j - 1

		for i >= 0 {
			if input[i+1] >= input[i] {
				break
			} else {
				// keep swapping input[i] with input[i+1] (Kj)
				temp := input[i+1]
				input[i+1] = input[i]
				input[i] = temp
				i--
			}
		}
	}

	return input
}

func moddedStraightInsertionSort(jump int, input []int32) []int32 {
	for j := jump; j < len(input); j++ {
		i := j - jump

		for i >= 0 {
			if input[i+jump] >= input[i] {
				break
			} else {
				// keep swapping input[i] with input[i+jump] (Kj)
				temp := input[i+jump]
				input[i+jump] = input[i]
				input[i] = temp
				i -= jump
			}
		}
	}

	return input
}

// We find that the bulk of the time complexity arises from moving one element at a
// time. We realize that, if we find a way through which we can move elements over a
// larger distance we might be able to write a more efficient algorithm.
// This is where Shell sort comes in. aka, "diminishing increment sort"
func shellSort(input []int32) []int32 {
	var increments []int
	// This computation actually slows the overall function down a bit, but it is mentioned in TAOCP Vol.3,
	// so I decided to add it here.
	if len(input) < 100 {
		for i := 0; i < 3*len(input); i++ {
			increments = append(increments, 3*i+1)
		}
	} else {
		for i := 1; i < 3*len(input); i++ {
			increments = append(increments, int(math.Pow(float64(2), float64(i))))
		}
	}

	// we need two loops: 1. for the increments, 2. for the numbers themselves (Happens
	// inside the moddedStraightInsertionSort)
	// first, lets loop on the increments
	for i := len(increments) - 1; i > -1; i-- {
		// now, we need to sort elements with this increment for each number i
		moddedStraightInsertionSort(increments[i], input)
	}
	return input
}

// Further, we realize that the insertion sort can be improved by settling for a Datastructure
// that is better suited for random insertion operations. And ofcourse, this would be the linked
// list. The linked list is much better than a conventional array when it comes to insertions
// because we need not shift everything that comes after.
