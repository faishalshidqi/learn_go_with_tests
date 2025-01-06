package arraysAndSlice

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := make([]int, 0)
	for _, num := range numbers {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	sums := make([]int, 0)
	for _, arr := range numbers {
		if len(arr) > 0 {
			tail := arr[1:]
			sums = append(sums, Sum(tail))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
