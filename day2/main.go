package main

import (
	"adventOfCode/internal/lists"
	"fmt"
)

func main() {
	Part1("data.txt")
	Part2("data.txt")
}

func Part1(filename string) int {
	safeReports := 0
	reports := lists.ReadReportsFromFile(filename)
	for _, report := range reports {
		numbers := lists.StringListToNumbers(report)
		if lists.IsSafeReport(numbers) {
			safeReports++
			fmt.Println(report)
		}
	}
	fmt.Println("part1", safeReports)
	return safeReports
}

func Part2(filename string) int {
	safeReports := 0
	reports := lists.ReadReportsFromFile(filename)
	for _, report := range reports {
		numbers := lists.StringListToNumbers(report)
		if lists.IsSafeReportIfRemoveOne(numbers) {
			safeReports++
		}
	}
	fmt.Println("part2", safeReports)
	return safeReports
}
