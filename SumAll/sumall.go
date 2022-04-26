package sumall

import "github.com/loqutus/go-tdd/sum"

// returns slice with sums of all numbers in each slice
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, sum.Sum(numbers))
	}
	return sums
}
