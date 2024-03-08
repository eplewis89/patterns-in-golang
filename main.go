package main

import (
	"fmt"

	"github.com/eplewis89/patterns-in-golang/patterns"
	"github.com/eplewis89/patterns-in-golang/utils"
)

func main() {
	//SlidingWindows()
	SlidingWindowsTest()
	//TwoPointers()
}

func SlidingWindows() {
	fmt.Println("============= Sliding Windows ==================")

	// should error out
	vals := patterns.SlidingWindow(-1, -1, 2)

	for _, val := range vals {
		fmt.Println(val)
	}

	// should error out
	vals = patterns.SlidingWindow(1, 1, 2)

	for _, val := range vals {
		fmt.Println(val)
	}

	// should print the whole array
	vals = patterns.SlidingWindow(10, 50, 10)

	for _, val := range vals {
		fmt.Println(val)
	}

	// should print subsections
	vals = patterns.SlidingWindow(10, 5, 3)

	for _, val := range vals {
		fmt.Println(val)
	}

	fmt.Println("================================================")
}

func SlidingWindowsTest() {
	fmt.Println("============= Sliding Windows Test =============")

	// test the sliding windows algorithms
	arr := utils.GenerateIntArray(500, 20, true)
	vals := patterns.FindSubarraysForSum(arr, 15, 3)

	for _, val := range vals {
		fmt.Println(val)
	}

	vals = patterns.FindSubarraysForSumSlidingWindow(arr, 15, 3)

	for _, val := range vals {
		fmt.Println(val)
	}

	fmt.Println("================================================")
}

func TwoPointers() {
	fmt.Println("============= Two Pointers =====================")

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

	fmt.Println("================================================")
}
