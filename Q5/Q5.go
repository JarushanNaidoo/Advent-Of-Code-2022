package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

// Instructions will be stored here
type Instructions struct {
	number string
	from   string
	to     string
}

func main() {
	file, _ := os.Open("Q5.txt")

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)

	stacks := make(map[string][]string)

	var instructions []Instructions

	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		// Put stack structure from txt file into a map structure
		length := len(line)
		for i := 0; i < length; i++ {
			regex, _ := regexp.Compile("[A-Z]")
			if regex.MatchString(string(line[i])) {
				index := math.Ceil(float64(i/4.0)) + 1
				strindex := strconv.Itoa(int(index))
				stacks[strindex] = append(stacks[strindex], string(line[i]))
			}
		}
		// read and make sense of moves and save them
		if len(line) > 0 && line[0:4] == "move" {
			insRegex, _ := regexp.Compile("([1-9][0-9]|\\d)")
			var tempSlice []string
			for i := 0; i < len(line); i++ {
				if insRegex.MatchString(line) {
					tempSlice = append(tempSlice, insRegex.FindAllString(line, -1)...)
				}

			}
			instructions = append(instructions, Instructions{number: tempSlice[0], from: tempSlice[1], to: tempSlice[2]})
		}
	}

	// go through each move and apply them
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		number, _ := strconv.Atoi(instruction.number)

		popResult := stacks[instruction.from][:number]                          // get value from top of stack
		stacks[instruction.from] = stacks[instruction.from][len(popResult):]    // remove the top of the stack from the stack
		stacks[instruction.to] = appendStack(stacks[instruction.to], popResult) // append popResult to top of new stack
	}

	// get values of top crate in stack and concatenate them to get the solution
	var answer string
	for i := 1; i <= len(stacks); i++ {
		strI := strconv.Itoa(i)
		if len(stacks[strI]) > 0 {
			answer += stacks[strI][0]
		}
	}
	fmt.Println(answer)
}

func appendFirst(slice []string, popResult string) []string {
	var result []string
	result = append(result, popResult)
	result = append(result, slice[0:]...)

	return result
}

func appendStack(slice []string, popResult []string) []string {
	var result []string
	result = append(result, popResult...)
	result = append(result, slice...)

	return result
}
