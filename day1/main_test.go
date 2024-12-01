package main

import (
	"testing"
)

func TestStep2(t *testing.T) {
	tt := []struct {
		filename string
		want     int
	}{
		{"testdata/data.txt", 31},
		{"data.txt", 20719933},
	}
	for _, tc := range tt {
		got := Step2(tc.filename)
		if got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
