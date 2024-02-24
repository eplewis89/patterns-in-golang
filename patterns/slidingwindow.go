package patterns

import (
	"strconv"

	"github.com/eplewis89/patterns-in-golang/utils"
)

// sliding window is described as an algorithm
// which performs an operation on a subset of an
// array, such as finding the max value from a
// subarray of specified length
func SlidingWindow(len, max, window int) {
	// we can't do a sliding window with zero values
	if len < 1 || max < 1 || window < 1 {
		println("cannot use number less than 1")
		return
	}

	// window can be equal but not greater than length of array
	if window > len {
		println("window cannot be greater than length of array")
		return
	}

	// generate a random int array
	arr := utils.GenerateIntArray(len, max, true)

	// iterate through the array only up to the length
	// minus window so we don't run into out of index issue
	for i := 0; i <= len-window; i++ {
		str := "[" + strconv.Itoa(arr[i])

		// iterate through the window
		for w := 1; w < window; w++ {
			str += "," + strconv.Itoa(arr[w+i])
		}

		str += "]"

		println(str)
	}
}
