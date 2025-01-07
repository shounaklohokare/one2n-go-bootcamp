package main

import (
	"math"
)

func getEvenNumbers(nums []int) []int {

	var out []int
	for _, n := range nums {
		if n%2 == 0 {
			out = append(out, n)
		}
	}

	return out

}

func getOddNumbers(nums []int) []int {

	var out []int
	for _, n := range nums {
		if n%2 != 0 {
			out = append(out, n)
		}
	}

	return out
}

func getPrimeNumbers(nums []int) []int {

	var out []int
	for _, n := range nums {
		if isPrime(n) {
			out = append(out, n)
		}
	}

	return out
}

func isPrime(n int) bool {

	if n == 1 {
		return false
	}

	var i = 2
	for i <= int(math.Floor(math.Sqrt(float64(n)))) {
		if n%i == 0 {
			return false
		}
		i += 1
	}
	return true

}

func getOddPrimeNumbers(nums []int) []int {

	var out []int
	for _, n := range nums {
		if n%2 == 0 && isPrime(n) {
			out = append(out, n)
		}
	}

	return out

}
