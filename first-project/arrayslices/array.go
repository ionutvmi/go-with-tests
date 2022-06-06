package arrayslices

func Sum(numbers []int) int {
	result := 0

	// _ is the blank identifier
	// helps us avoid the `unused variable` errors
	// index, value
	for _, val := range numbers {
		result += val
	}

	return result
}
