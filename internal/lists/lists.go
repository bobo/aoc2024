package lists

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"adventOfCode/internal/memoize"
)

func CreateLists(filename string) [][]int {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading file %s: %v", filename, err)
	}
	lines := strings.Split(string(data), "\n")
	listA := []int{}
	listB := []int{}
	for _, line := range lines {
		digits := strings.Split(line, "   ")
		numA, _ := strconv.Atoi(digits[0])
		numB, _ := strconv.Atoi(digits[1])
		listA = append(listA, numA)
		listB = append(listB, numB)
	}
	return [][]int{listA, listB}
}

func SortLists(lists [][]int) [][]int {
	sort.Ints(lists[0])
	sort.Ints(lists[1])
	return lists
}

func ToDistances(listA, listB []int) []int {
	distances := []int{}
	for i := 0; i < len(listA); i++ {
		distances = append(distances, int(math.Abs(float64(listA[i]-listB[i]))))
	}
	return distances
}

func SumList(list []int) int {
	sum := 0
	for _, number := range list {
		sum += number
	}
	return sum
}

func CountNumberInList(number int, list []int) int {
	count := 0
	for _, num := range list {
		if num == number {
			count++
		}
	}
	return count
}

func FindSimilarity(listA, listB []int) int {
	momoizedCount := memoize.Memoize(CountNumberInList)
	similarity := 0
	for i := 0; i < len(listA); i++ {
		count := momoizedCount.Call(listA[i], listB)
		similarity += count * listA[i]
	}
	return similarity
}

func ReadReportsFromFile(filename string) [][]string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading file %s: %v", filename, err)
	}
	lines := strings.Split(string(data), "\n")
	reports := [][]string{}
	for _, line := range lines {
		reports = append(reports, strings.Split(line, " "))
	}
	return reports
}

func StringListToNumbers(strings []string) []int {
	numbers := []int{}
	for _, string := range strings {
		number, _ := strconv.Atoi(string)
		numbers = append(numbers, number)
	}
	return numbers
}

func AllIncreasing(list []int) bool {
	for i, candidate := range list {
		if i > 0 && candidate <= list[i-1] {
			return false
		}
	}
	return true
}

func AllDecreasing(list []int) bool {
	for i, candidate := range list {
		if i > 0 && candidate >= list[i-1] {
			return false
		}
	}
	return true
}

func DifferesByOneToThree(list []int) bool {
	for i := 1; i < len(list); i++ {
		diff := Abs(list[i] - list[i-1])
		if diff > 3 || diff < 1 {
			return false
		}

	}
	return true
}

func GetPossibleLists(list []int) [][]int {
	possibleLists := [][]int{}
	possibleLists = append(possibleLists, list)
	for i := 0; i < len(list); i++ {
		listToAdd := RemoveItemOnIndex(list, i)
		possibleLists = append(possibleLists, listToAdd)
	}
	return possibleLists
}

func RemoveItemOnIndex(slice []int, index int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	return append(sliceCopy[:index], sliceCopy[index+1:]...)
}

func IsSafeReport(report []int) bool {
	return (AllIncreasing(report) || AllDecreasing(report)) && DifferesByOneToThree(report)
}

func IsSafeReportIfRemoveOne(report []int) bool {
	possibleLists := GetPossibleLists(report)
	for _, possibleList := range possibleLists {
		if IsSafeReport(possibleList) {
			return true
		}
	}
	return false
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
