package linearsearch

func linearsearch(arrayNumbers []int, value int) bool {

	for _, v := range arrayNumbers {
		if v == value {
			return true
		}
	}
	return false
}
