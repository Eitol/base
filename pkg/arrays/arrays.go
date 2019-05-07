package arrays

func InStrArray(array []string, item interface{}) bool {
	for i := 0; i < len(array); i++ {
		if array[i] == item {
			return true
		}
	}
	return false
}
