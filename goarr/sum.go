package arrays

// Sum function
func Sum(numbers []int) int {

	sum := 0

	for _, value := range numbers {
		sum += value
	}
	return sum
}

// SumAllTails function
func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int
	for _, value := range numbersToSum {

		if len(value) == 0 {
			sums = append(sums, 0)
		} else {
			tail := value[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

// Multiply function
func Multiply(numbers []int) int {

	mult := 1

	for _, value := range numbers {
		mult = mult * value
	}

	return mult
}
