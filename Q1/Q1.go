package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// Open file
	file, err := os.Open("Q1.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}

	// Read contents from file
	fileScanner := bufio.NewScanner(file)

	// Split file into lines
	fileScanner.Split(bufio.ScanLines)

	var max int = 0
	var sum int = 0
	var totals []int
	// Loop through each line and evaluate
	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		if line != "" {
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("CONVERSION ERROR")
			}
			sum += num
		} else {
			totals = append(totals, sum)
			if max < sum {
				max = sum
			}
			sum = 0
		}
	}

	fmt.Println(findMaxs(totals))
	defer file.Close()
}

func findMaxs(totals []int) int {
	sumMax := 0
	var count int = 0
	for count < 3 {
		max := 0
		var index int = 0
		for i := 0; i < len(totals); i++ {
			if max < totals[i] {
				max = totals[i]
				index = i
			}
		}
		totals[index] = 0
		sumMax += max
		count++
	}
	return sumMax
}
