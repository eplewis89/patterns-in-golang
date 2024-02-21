package patterns

import (
	"strconv"

	"github.com/eplewis89/patterns-in-golang/utils"
)

func SlidingWindow(len, max, window int) {
	if len < 1 || max < 1 || window < 1 {
		println("cannot use number less than 1")
		return
	}

	arr := utils.GenerateIntArray(len, max, true)

	if window > len {
		println("window cannot be greater than length of array")
		return
	}

	for i := 0; i <= len-window; i++ {
		str := "[" + strconv.Itoa(arr[i])

		for w := 1; w < window; w++ {
			str += "," + strconv.Itoa(arr[w+i])
		}

		str += "]"

		println(str)
	}
}
