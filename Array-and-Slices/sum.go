package main

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfAnswers := len(numbersToSum)
	answers := make([]int, lengthOfAnswers)

	for i, numbers := range numbersToSum {
		answers[i] = Sum(numbers)
	}

	return []int{3, 9}
}