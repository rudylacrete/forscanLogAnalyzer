package utils

// ArrayIndex return true if the given index of the string or -1 if not found
func ArrayIndex(arr []string, val string) int {
	for i, v := range arr {
		if v == val {
			return i
		}
	}
	return -1
}
