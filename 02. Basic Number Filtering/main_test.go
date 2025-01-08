package main

import (
	"reflect"
	"testing"
)

func TestGetEvenNumbers(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := []int{2, 4, 6, 8, 10}

	out := getEvenNumbers(nums)

	if !reflect.DeepEqual(expected, out) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}

}

func TestGetOddNumbers(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := []int{1, 3, 5, 7, 9}

	out := getOddNumbers(nums)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}
}

func TestGetPrimeNumbers(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := []int{2, 3, 5, 7}

	out := getPrimeNumbers(nums)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}

}

func TestGetOddPrimeNumbers(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	expected := []int{3, 5, 7}

	out := getOddPrimeNumbers(nums)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}

}

func TestGetEvenMultiplesOfFive(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	expected := []int{10, 20}

	isEven := func(num int) bool {
		return num%2 == 0
	}

	isMultipleOfFive := func(num int) bool {
		return num%5 == 0
	}

	out := getAllConditions(nums, isEven, isMultipleOfFive)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}
}

func TestOddMultiplesOfThreeGreaterThanTen(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	expected := []int{15}

	isOdd := func(num int) bool {
		return num%2 != 0
	}

	isMultipleOfThree := func(num int) bool {
		return num%3 == 0
	}

	isGreaterThanTen := func(num int) bool {
		return num > 10
	}

	out := getAllConditions(nums, isOdd, isMultipleOfThree, isGreaterThanTen)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}

}

func TestOddGreaterThanFiveAndMultipleOfSeven(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	expected := []int{7, 21}

	isOdd := func(num int) bool {
		return num%2 != 0
	}

	isGreaterThanFive := func(num int) bool {
		return num > 5
	}

	isMultipleOfSeven := func(num int) bool {
		return num%7 == 0
	}

	out := getAllConditions(nums, isOdd, isGreaterThanFive, isMultipleOfSeven)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}

}

func TestGetAnyGreaterThanTwelvePrimeEvenNumber(t *testing.T) {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}

	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 22, 23, 24, 25}

	isGreaterThanTwentyOne := func(num int) bool {
		return num > 21
	}

	isMultipleOfNineteen := func(num int) bool {
		return num%19 == 0
	}

	out := getAnyCondition(nums, isPrime, isGreaterThanTwentyOne, isMultipleOfNineteen)

	if !reflect.DeepEqual(out, expected) {
		t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", expected, out)
	}

}

func TestMatchAllConditionsNumberFilter(t *testing.T) {

	isOdd := func(n int) bool { return n%2 != 0 }
	isEven := func(n int) bool { return n%2 == 0 }
	getGreaterThanN := func(n int) CustomCondition { return func(m int) bool { return m > n } }
	greaterThan8 := getGreaterThanN(8)
	getMultiplesOf := func(n int) CustomCondition { return func(m int) bool { return m%n == 0 } }
	getMultiplesOf3 := getMultiplesOf(3)
	getLessThanN := func(n int) CustomCondition { return func(m int) bool { return m < n } }
	getLessThan15 := getLessThanN(15)

	tt := []struct {
		nums     []int
		conds    []CustomCondition
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds:    []CustomCondition{isOdd, greaterThan8, getMultiplesOf3},
			expected: []int{9, 15},
		},
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds:    []CustomCondition{isEven, getLessThan15, getMultiplesOf3},
			expected: []int{6, 12},
		},
	}

	for _, tc := range tt {
		out := getAllConditions(tc.nums, tc.conds...)

		if !reflect.DeepEqual(tc.expected, out) {
			t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", tc.expected, out)
		}
	}
}

func TestMatchAnyConditionsNumberFilter(t *testing.T) {

	greaterThanN := func(n int) CustomCondition { return func(m int) bool { return m > n } }
	greaterThan13 := greaterThanN(13)
	multiplesOf := func(n int) CustomCondition { return func(m int) bool { return m%n == 0 } }
	multiplesOf6 := multiplesOf(6)
	multiplesOf3 := multiplesOf(3)
	lessThanN := func(n int) CustomCondition { return func(m int) bool { return m < n } }
	lessThan8 := lessThanN(8)

	tt := []struct {
		nums     []int
		conds    []CustomCondition
		expected []int
	}{
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13, 14, 15, 16, 17, 18, 19, 20},
			conds:    []CustomCondition{isPrime, greaterThan13, multiplesOf6},
			expected: []int{2, 3, 5, 6, 7, 11, 13, 14, 15, 16, 17, 18, 19, 20},
		},
		{
			nums:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			conds:    []CustomCondition{lessThan8, multiplesOf3},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 9, 12, 15, 18},
		},
	}

	for _, tc := range tt {
		out := getAnyCondition(tc.nums, tc.conds...)

		if !reflect.DeepEqual(tc.expected, out) {
			t.Errorf("Test failed!\nExpected:- %v\toutput:- %v", tc.expected, out)
		}
	}
}
