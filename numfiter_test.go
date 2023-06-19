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