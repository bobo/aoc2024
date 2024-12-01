package lists

import (
	"reflect"
	"testing"
)

func TestCreateLists(t *testing.T) {
	got := CreateLists("testdata/data.txt")
	want := [][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
func TestSortLists(t *testing.T) {
	got := SortLists([][]int{{3, 4, 2, 1, 3, 3}, {4, 3, 5, 3, 9, 3}})
	want := [][]int{{1, 2, 3, 3, 3, 4}, {3, 3, 3, 4, 5, 9}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestToDistances(t *testing.T) {
	got := ToDistances([]int{1, 2, 3, 3, 3, 4}, []int{3, 3, 3, 4, 5, 9})
	want := []int{2, 1, 0, 1, 2, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestSumList(t *testing.T) {
	got := SumList([]int{2, 1, 0, 1, 2, 5})
	want := 11
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestCountNumberInList(t *testing.T) {
	tt := []struct {
		number int
		list   []int
		want   int
	}{
		{1, []int{1, 2, 1}, 2},
		{2, []int{1, 2, 1}, 1},
		{3, []int{1, 2, 1}, 0},
		{1, []int{1, 2, 1, 1, 1, 1, 2, 3, 1}, 6},
	}
	for _, tc := range tt {
		got := CountNumberInList(tc.number, tc.list)
		if got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestFindSimilarity(t *testing.T) {
	tt := []struct {
		listA []int
		listB []int
		want  int
	}{
		{[]int{1, 2, 3}, []int{1, 1, 2}, 1*2 + 2*1 + 3*0},
		{[]int{1, 2, 3}, []int{1, 2, 3}, 1*1 + 2*1 + 3*1},
		{[]int{1, 2, 3}, []int{4, 5, 6}, 1*0 + 2*0 + 3*0},
		{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}, 31},
	}
	for _, tc := range tt {
		got := FindSimilarity(tc.listA, tc.listB)
		if got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}
