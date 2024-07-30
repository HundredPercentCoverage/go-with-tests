package iteration

// Repeat takes a given string and repeats it a given number of times
func Repeat(str string, times int) string {
	var result string

	for i := 0; i < times; i++ {
		result += str
	}

	return result
}
