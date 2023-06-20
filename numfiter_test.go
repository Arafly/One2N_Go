package main

import (
	"testing"
	"reflect"
)

func TestGetEven (t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,10}
	expected := []int{2, 4, 6, 8, 10}

	output := GetEven(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("The Result was incorrect as GetEven(%v) failed, expected %v, but got %v",input, expected, output)
	}
}

func TestOdd (t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,10}
	expected := []int{1, 3, 5, 7, 9}

	output := GetOdd(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("The Result was incorrect as GetOdd(%v) failed, expected %v, but got %v",input, expected, output)
	}
}

func TestPrime (t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,10}
	expected := []int{2, 3, 5, 7}

	output := GetPrime(input)
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("The Result was incorrect as GetPrime(%v) failed, expected %v, but got %v",input, expected, output)
	}
}

func TestOddPrime (t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9,10}
	expected := []int{3, 5, 7}

	output := GetOdd(GetPrime(input))
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("The Result was incorrect as GetOdd(GetPrime(%v)) failed, expected %v, but got %v",input, expected, output)
	}
}

func TestEvenFiveMultiple (t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	expected := []int{10, 20}

	output := GetEven(GetFiveMultiple(input))
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("The Result was incorrect as GetEven(GetFiveMultiple(%v)) failed, expected %v, but got %v",input, expected, output)
	}
}