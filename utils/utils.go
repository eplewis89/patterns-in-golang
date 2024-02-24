package utils

import (
	"math/rand"
	"strconv"
)

func GenerateIntArray(len, maxn int, ran bool) []int {
	arr := []int{}

	for i := 0; i < len; i++ {
		val := i

		if ran {
			val = rand.Intn(maxn)
		}

		arr = append(arr, val)
	}

	return arr
}

func PrintIntArray(arr []int) {
	if len(arr) == 0 {
		println("[]")
	} else {
		str := "[" + strconv.Itoa(arr[0])

		for i := 1; i < len(arr); i++ {
			str += ", " + strconv.Itoa(arr[i])
		}

		str += "]"

		println(str)
	}
}
