package main

import "fmt"

func getMostVisited(m int, arr []int) int {
	// Dynamic programming
	var ans int
	array := [10]int{}
	var left int
	var right int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			left = i
			right = i + 1
		} else {
			left = i + 1
			right = i
		}

		// Instead of increasing every portion of the array by 1 and get O(n^2),
		// you increase the left most part of the array by 1 and decrease the
		// right most part of the array that isn't in between the array so that only
		// this segment will have it's value be the same as the leftmost with + 1 and
		// - 1 for any start or end tracks covered.

		array[left]++
		array[right+1]--
	}

	max := 0
	s := 0

	for i := 0; i < len(array); i++ {
		array[i] += s

		if array[i] > max {
			max = array[i]
			ans = i + 1
		}

		s = array[i]
	}

	return ans
}

func main() {
	m := 10
	arr := []int{4, 2, 1, 5, 2, 8, 9, 1, 10}

	a := getMostVisited(m, arr)

	fmt.Println(a)
}
