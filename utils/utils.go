package utils

import "math/rand"

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
