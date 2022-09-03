package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T){
	t.Run("got new slice containing the totals for each slice passed in", func(t *testing.T){
		got := SumAll([]int{1,2}, []int{0,9})
		expected := []int{3,9}

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, expected %v", got, expected)
		}
	})
}