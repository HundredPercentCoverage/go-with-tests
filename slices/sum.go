package slices

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(numberSlices ...[]int) []int {
	var sums []int

	for _, numberSlice := range numberSlices {
		sums = append(sums, Sum(numberSlice))
	}

	return sums
}

func SumTails(numberSlices ...[]int) []int {
	var sums []int

	for _, numberSlice := range numberSlices {
		if len(numberSlice) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(numberSlice[1:]))
		}
	}

	return sums
}
