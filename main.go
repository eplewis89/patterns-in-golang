package main

import (
	"github.com/eplewis89/patterns-in-golang/patterns"
	"github.com/eplewis89/patterns-in-golang/utils"
)

func main() {
	SlidingWindows()
	TwoPointers()
}

func SlidingWindows() {
	// should error out
	patterns.SlidingWindow(1, -1, 2)

	// should error out
	patterns.SlidingWindow(1, 1, 2)

	// should print the whole array
	patterns.SlidingWindow(10, 50, 10)

	// should print subsections
	patterns.SlidingWindow(100, 50, 3)
}

func TwoPointers() {
	arr := []int{0, 1, 15, 16, 25, 30, 37}

	// this should work and print "[0, 30]"
	utils.PrintIntArray(patterns.TwoPointers(arr, 30))

	// this should work and print "[30, 37]"
	utils.PrintIntArray(patterns.TwoPointers(arr, 67))

	// this should work and print "[0, 1]"
	utils.PrintIntArray(patterns.TwoPointers(arr, 1))

	// this should not work and print "[]"
	utils.PrintIntArray(patterns.TwoPointers(arr, 12))

	// this should not work and print "[]"
	utils.PrintIntArray(patterns.TwoPointers([]int{}, 12))
}
