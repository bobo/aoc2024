package lists

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
	similarity := 0
	for i := 0; i < len(listA); i++ {
		count := CountNumberInList(listA[i], listB)
		similarity += count * listA[i]
	}
	return similarity
}
