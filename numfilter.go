// https://playbook.one2n.in/go-bootcamp/go-projects/basic-number-filtering-exercise

package main

import "fmt"

func  GetEven (numbers []int) []int {
	var even []int
	for _, num := range numbers {
		if num % 2 == 0 {
			even = append(even, num)
		}
	}
	return even
}