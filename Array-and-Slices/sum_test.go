package main

import "testing"

func TestSum(t *testing.T){
	//take an array of numbers and return the total.

	t.Run("collection of any numbers", func(t *testing.T) {
		numbers := []int{1,2,3}

		got := Sum(numbers)
		expected := 6
	
		if got != expected {
			t.Errorf("got %d, expected %d, given %v", got, expected, numbers)
		}
	})	
}