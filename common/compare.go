package common

func CompareStringArrays(arr1, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	freq1 := make(map[string]int)
	freq2 := make(map[string]int)
	for _, item := range arr1 {
		freq1[item]++
	}
	for _, item := range arr2 {
		freq2[item]++
	}
	if len(freq1) != len(freq2) {
		return false
	}
	for key, val1 := range freq1 {
		if val2, ok := freq2[key]; !ok || val1 != val2 {
			return false
		}
	}
	return true
}
