package helper

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Intersect(arr1 []int, arr2 []int) []int {
	intersected := make([]int, 0, Max(len(arr1), len(arr2)))
	for _, val := range arr2 {
		if Contains(arr1, val) {
			intersected = append(intersected, val)
		}
	}
	return intersected
}
