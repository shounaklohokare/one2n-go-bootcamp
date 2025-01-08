// run go test -v to get test results
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
		if n%2 != 0 && isPrime(n) {
			out = append(out, n)
		}
	}

	return out

}

type CustomCondition func(num int) bool

func getAllConditions(nums []int, customConditions ...CustomCondition) []int {

	var out []int
	for _, num := range nums {

		conditionFulfill := true
		for _, customCondition := range customConditions {
			if !customCondition(num) { // if any one condition is not getting fulfilled then we don't want that num
				conditionFulfill = false
				break
			}
		}

		if conditionFulfill {
			out = append(out, num)
		}

	}

	return out

}

func getAnyCondition(nums []int, customConditions ...CustomCondition) []int {

	var out []int
	for _, num := range nums {

		for _, customCondition := range customConditions {
			if customCondition(num) {
				out = append(out, num) // if any one condition is getting fulfilled then we accept that number
				break
			}
		}

	}

	return out

}
