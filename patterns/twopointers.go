package patterns

func TwoPointers(arr []int, sum int) []int {
	ret := []int{}

	if len(arr) < 2 {
		return ret
	}

	// set a counter so we don't exceed arr len
	c := 0

	// set a sum value here
	x := 0

	pointerA := 0
	pointerB := len(arr) - 1

	// while loop
	for x != sum {
		x := arr[pointerA] + arr[pointerB]

		if x == sum {
			ret = []int{arr[pointerA], arr[pointerB]}
			break
		}

		if x < sum {
			pointerA++
		}

		if x > sum {
			pointerB--
		}

		c += 1

		if c > len(arr) {
			break
		}
	}

	return ret
}
