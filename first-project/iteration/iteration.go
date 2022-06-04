package iteration

// Repeats a substring the specified amount of times
func Repeat(substring string, repeatCount int) string {
	result := ""

	for i := 0; i < repeatCount; i++ {
		result += substring
	}

	return result
}
