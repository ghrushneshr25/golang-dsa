package common

func CompareStringArrays(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	freq := make(map[string]int)
	for _, item := range arr1 {
		freq[item]++
	}
	for _, item := range arr2 {
		freq[item]--
		if freq[item] < 0 {
			return false
		}
	}
	for _, count := range freq {
		if count != 0 {
			return false
		}
	}
	return true
}
