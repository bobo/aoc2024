package day2

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("testdata/data.txt")
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := Part2("testdata/data.txt")
	want := 4
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
