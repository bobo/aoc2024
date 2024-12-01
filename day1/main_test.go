package main

import (
	"testing"
)

func TestStep2(t *testing.T) {
	got := Step2("testdata/data.txt")
	want := 31
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
