package main

import (
	"reflect"
	"testing"
)

// TestSplitLines tests the functionality of the SplitLines function.
func TestSplitLines(t *testing.T) {
	got := LinesFromFile("testdata/data.txt")
	want := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestLineToDigits(t *testing.T) {
	tt := []struct {
		line string
		want []int
	}{
		{"1abc2", []int{1, 2}},
		{"pqr3stu8vwx", []int{3, 8}},
		{"a1b2c3d4e5f", []int{1, 5}},
		{"treb7uchet", []int{7, 7}},
	}
	for _, tc := range tt {
		got := LineToDigits(tc.line)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestMergeDigits(t *testing.T) {
	tt := []struct {
		digits []int
		want   int
	}{
		{[]int{1, 2}, 12},
		{[]int{3, 8}, 38},
		{[]int{1, 5}, 15},
		{[]int{7, 7}, 77},
	}
	for _, tc := range tt {
		got := MergeDigits(tc.digits)
		if got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
