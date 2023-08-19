package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// open file
	file, err := os.Open("Q4.txt")
	if err != nil {
		log.Fatalf("Error opening file")
	}
	defer file.Close()

	// Scan file
	fileScanner := bufio.NewScanner(file)

	// Split file into lines
	fileScanner.Split(bufio.ScanLines)

	numOfOverlaps := 0

	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		splitLine := strings.Split(line, ",")
		bounds1 := strings.Split(splitLine[0], "-")
		bounds2 := strings.Split(splitLine[1], "-")

		lowerBound1, _ := strconv.Atoi(bounds1[0])
		upperBound1, _ := strconv.Atoi(bounds1[1])
		lowerBound2, _ := strconv.Atoi(bounds2[0])
		upperBound2, _ := strconv.Atoi(bounds2[1])

		// conditionA := (lowerBound1 >= lowerBound2 && upperBound1 <= upperBound2) || (lowerBound2 >= lowerBound1 && upperBound2 <= upperBound1)
		conditionB := (lowerBound1 >= lowerBound2 && lowerBound1 <= upperBound2) || (upperBound1 <= upperBound2 && upperBound1 >= lowerBound2) || (lowerBound2 >= lowerBound1 && lowerBound2 <= upperBound1) || (upperBound2 <= upperBound1 && upperBound2 >= lowerBound1)

		if conditionB {
			numOfOverlaps++
		}
	}

	fmt.Println(numOfOverlaps)
}
