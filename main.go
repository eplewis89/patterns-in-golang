package main

import "github.com/eplewis89/patterns-in-golang/patterns"

func main() {
	SlidingWindows()
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
