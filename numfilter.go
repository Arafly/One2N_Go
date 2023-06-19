// https://playbook.one2n.in/go-bootcamp/go-projects/basic-number-filtering-exercise

package main

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

func GetPrime ()