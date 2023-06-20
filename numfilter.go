// https://playbook.one2n.in/go-bootcamp/go-projects/basic-number-filtering-exercise

package main

import (
	"math"
)

func GetEven (numbers []int) []int {
	var even []int
	for _, num := range numbers {
		if num % 2 == 0 {
			even = append(even, num)
		}
	}
	return even
}

func GetOdd (numbers []int) []int {
	var odd []int
	for _, num := range numbers {
		if num % 2 != 0 {
			odd = append(odd, num)
		}
	}
	return odd
}

func isPrime (number int) bool {
	if number < 2 {
		return false
	}

	limitation := int(math.Sqrt(float64(number)))

	for  i := 2; i <= limitation; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func GetPrime (numbers []int) []int {
	var primeNumbers []int
	for _, num := range numbers {
		if isPrime(num) {
			primeNumbers = append(primeNumbers, num)
		}
	}

	return primeNumbers
}

func OddPrime (numbers []int) []int {
	return GetOdd(GetPrime(numbers))
}

func GetFiveMultiple (numbers []int) []int {
	var fiveMultiple []int
	for _, num := range numbers {
		if num % 5 == 0 {
			fiveMultiple = append(fiveMultiple, num)
		}
	}
	return fiveMultiple
}