package patterns

import (
	"github.com/eplewis89/patterns-in-golang/utils"
)

// this sliding window algorithm creates subsets of
// int arrays and returns them
func SlidingWindow(len, max, window int) [][]int {
	// hold the return values here
	ret := [][]int{}

	// we can't do a sliding window with zero values
	if len < 1 || max < 1 || window < 1 {
		println("cannot use window or sum less than 1")
		return ret
	}

	// window can be equal but not greater than length of array
	if window > len {
		println("window cannot be greater than length of array")
		return ret
	}

	// this is in the "utils" library and generates
	// a random int array of size len, with max value
	arr := utils.GenerateIntArray(len, max, true)

	// iterate through the array only up to the length
	// minus window so we don't run into out of index issue
	for i := 0; i <= len-window; i++ {
		// hold current int array
		cur := []int{arr[i]}

		// iterate through the window
		for w := 1; w < window; w++ {
			cur = append(cur, arr[w+i])
		}

		ret = append(ret, cur)
	}

	return ret
}

// this algorithm matches a subarray of length window
// to the declared value sum and adds it to an array
// if the sum matches the values of the subarray
func FindSubarraysForSum(arr []int, sum, window int) [][]int {
	// hold the return values here
	ret := [][]int{}

	// we can't do a sliding window with zero values
	if sum < 1 || window < 1 {
		println("cannot use window or sum less than 1")
		return ret
	}

	// window can be equal but not greater than length of array
	if window > len(arr) {
		println("window cannot be greater than length of array")
		return ret
	}

	// iterate through the array only up to the length
	// minus window so we don't run into out of index issue
	for i := 0; i <= len(arr)-window; i++ {
		// hold current sum
		cur_sum := arr[i]

		// hold current int array
		cur := []int{arr[i]}

		// iterate through the window
		for w := 1; w < window; w++ {
			// add val to sum
			cur_sum += arr[w+i]
			// add val to current array
			cur = append(cur, arr[w+i])
		}

		// if our sum matches append the current arr
		if sum == cur_sum {
			ret = append(ret, cur)
		}
	}

	return ret
}

// this algorithm matches a subarray of length window
// to the declared value sum and adds it to an array
// if the sum matches the values of the subarray
func FindSubarraysForSumSlidingWindow(arr []int, sum, window int) [][]int {
	// hold the return values here
	ret := [][]int{}

	// we can't do a sliding window with zero values
	if sum < 1 || window < 1 {
		println("cannot use window or sum less than 1")
		return ret
	}

	// window can be equal but not greater than length of array
	if window > len(arr) {
		println("window cannot be greater than length of array")
		return ret
	}

	// hold the sum of the window
	window_sum := 0

	// hold current int array
	cur := []int{}

	// first, check the sum of the first {window size}
	// elements
	for i := 0; i < window; i++ {
		window_sum += arr[i]
		cur = append(cur, arr[i])
	}

	// if our sum matches append the current arr
	if window_sum == sum {
		ret = append(ret, cur)
	}

	// starting at window index, go to end of list
	// here we'll remove the zero-th element and append
	// the new element to the sum and array
	for i := window; i < len(arr); i++ {
		// add the current index value and subtract
		// the previous zero-th element from the sum
		window_sum += arr[i] - arr[i-window]

		// remove the zero-th element from the previous
		// window array and append new value to end
		cur = cur[1:]
		cur = append(cur, arr[i])

		// if our sum matches append the current arr
		if window_sum == sum {
			ret = append(ret, cur)
		}
	}

	return ret
}
