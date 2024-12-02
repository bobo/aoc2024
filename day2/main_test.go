package main

import "testing"

func TestPart1(t *testing.T) {
	got := Part1("testdata/data.txt")
	want := 2
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
