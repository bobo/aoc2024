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

func TestReadReportsFromFile(t *testing.T) {
	got := ReadReportsFromFile("testdata/day2.txt")
	want := [][]string{{"7", "6", "4", "2", "1"}, {"1", "2", "7", "8", "9"}, {"9", "7", "6", "2", "1"}, {"1", "3", "2", "4", "5"}, {"8", "6", "4", "4", "1"}, {"1", "3", "6", "7", "9"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestStringListToNumbers(t *testing.T) {
	got := StringListToNumbers([]string{"7", "6", "4", "2", "1"})
	want := []int{7, 6, 4, 2, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAllIncreasing(t *testing.T) {
	got := AllIncreasing([]int{1, 2, 3, 4, 5})
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestAllDecreasing(t *testing.T) {
	got := AllDecreasing([]int{5, 4, 3, 2, 1})
	want := true
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDifferesByOneOrTwo(t *testing.T) {
	tt := []struct {
		list []int
		want bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{1, 2, 4, 7}, true},
		{[]int{1, 2, 3, 7}, false},
		{[]int{1, 2, 3, 3, 4}, false},
	}
	for _, tc := range tt {
		got := DifferesByOneToThree(tc.list)
		if got != tc.want {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestGetPossibleLists(t *testing.T) {
	got := GetPossibleLists([]int{1, 2, 3, 4, 5})
	want := [][]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5}, {1, 3, 4, 5}, {1, 2, 4, 5}, {1, 2, 3, 5}, {1, 2, 3, 4}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRemoveItemOnIndex(t *testing.T) {
	tt := []struct {
		list  []int
		index int
		want  []int
	}{
		{[]int{1, 2, 3, 4, 5}, 0, []int{2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{1, 2, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{1, 2, 3, 4}},
	}
	for _, tc := range tt {
		got := RemoveItemOnIndex(tc.list, tc.index)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("got %v, want %v", got, tc.want)
		}
	}
}

func TestIsSafeReport(t *testing.T) {
	tt := []struct {
		report []int
		want   bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{1, 3, 6, 7, 9}, true},
	}
	for _, tc := range tt {
		got := IsSafeReport(tc.report)
		if got != tc.want {
			t.Errorf("got %v, want %v for report %v", got, tc.want, tc.report)
		}
	}
}

func TestIsSafeReportIfRemoveOne(t *testing.T) {
	tt := []struct {
		report []int
		want   bool
	}{
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{7, 6, 4, 2, 1}, true},
		{[]int{1, 2, 7, 8, 9}, false},
		{[]int{9, 7, 6, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, true},
		{[]int{1, 3, 6, 7, 9}, true},
		{[]int{8, 6, 4, 4, 1}, true},
	}
	for _, tc := range tt {
		got := IsSafeReportIfRemoveOne(tc.report)
		if got != tc.want {
			t.Errorf("got %v, want %v for report %v", got, tc.want, tc.report)
		}
	}
}
