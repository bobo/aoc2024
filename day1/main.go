package main

import (
	"adventOfCode/internal/lists"
	"fmt"
)

func main() {
	// Your code for the day's problem goes here

	//part 1
	sortedLists := lists.SortLists(lists.CreateLists("data.txt"))
	sums := lists.SumList(lists.ToDistances(sortedLists[0], sortedLists[1]))
	fmt.Println(sums)

	//part 2
	similarity := Step2("data.txt")
	fmt.Println(similarity)
}

func Step2(filename string) int {
	similarityLists := lists.CreateLists(filename)
	similarity := lists.FindSimilarity(similarityLists[0], similarityLists[1])
	return similarity
}
