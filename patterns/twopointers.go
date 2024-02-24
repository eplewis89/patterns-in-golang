package patterns

// in this two pointers example, we are
// taking an input array and finding the
// given sum of the input array, and
// return an array of the two values who
// add up to the target sum
func TwoPointers(arr []int, target int) []int {
	ret := []int{}

	if len(arr) < 2 {
		return ret
	}

	// set a counter so we don't exceed arr len
	counter := 0

	// pointerA points to the beginning
	// and is incremented when a sum is less
	// than the target
	pointerA := 0

	// pointerB points to the end of the list
	// and is decremented when the sum is
	// greater than the target
	pointerB := len(arr) - 1

	// while loop
	for counter < len(arr) {
		sum := arr[pointerA] + arr[pointerB]

		if sum == target {
			ret = []int{arr[pointerA], arr[pointerB]}
			break
		}

		// we need to move pointerA right
		// when the sum is less than target
		if sum < target {
			pointerA++
		}

		// we need to move pointerB left
		// when the sum is greater than target
		if sum > target {
			pointerB--
		}

		counter++
	}

	return ret
}
