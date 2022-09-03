package main

func Sum(numbers []int) int {

	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var answers []int

	for _, numbers := range numbersToSum {
		answers = append(answers, Sum(numbers))
	}

	return answers
}