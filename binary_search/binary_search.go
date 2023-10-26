package binarysearch

import "math"

func binarysearch(arrayNumbers []int, value int) bool {
	low := 0
	high := len(arrayNumbers)
	middle := int(math.Floor(float64(low + (high-low)/2)))
	for range arrayNumbers {
		valueMiddle := arrayNumbers[middle]
		if valueMiddle == value {
			return true
		} else if valueMiddle > value {
			high = valueMiddle
		} else {
			low = valueMiddle + 1

		}
	}
	return false
}
