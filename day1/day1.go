package day1

import (
	"adventOfCode/internal/lists"
	"fmt"
)

func Run() {
	// Your code for the day's problem goes here

	//part 1
	sortedLists := lists.SortLists(lists.CreateLists("day1/data.txt"))
	sums := lists.SumList(lists.ToDistances(sortedLists[0], sortedLists[1]))
	fmt.Println("part1", sums)

	//part 2
	similarity := Step2("day1/data.txt")
	fmt.Println("part2", similarity)
}

func Step2(filename string) int {
	similarityLists := lists.CreateLists(filename)
	similarity := lists.FindSimilarity(similarityLists[0], similarityLists[1])
	return similarity
}
