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
