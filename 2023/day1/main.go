package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := LinesFromFile("data.txt")
	sum := 0
	for _, line := range lines {
		sum += MergeDigits(LineToDigits(line))
	}
	fmt.Println(sum)
}

func LinesFromFile(filename string) []string {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error reading file %s: %v", filename, err)
	}
	return strings.Split(string(data), "\n")
}

func LineToDigits(line string) []int {
	digits := []int{}
	for _, char := range line {
		if char >= '0' && char <= '9' {
			digits = append(digits, int(char-'0'))
		}
	}
	if len(digits) > 2 {
		return []int{digits[0], digits[len(digits)-1]}
	}
	if len(digits) == 1 {
		return []int{digits[0], digits[0]}
	}
	return digits
}

func MergeDigits(digits []int) int {
	result := 0
	for _, digit := range digits {
		result = result*10 + digit
	}
	return result
}
