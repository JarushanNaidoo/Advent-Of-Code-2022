package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var A Score = Score{"Rock", 1}
var B Score = Score{"Paper", 2}
var C Score = Score{"Scissors", 3}

var X Score = Score{"Rock", 1}
var Y Score = Score{"Paper", 2}
var Z Score = Score{shape: "Scissors", score: 3}

type Score struct {
	shape string
	score int
}

const (
	win  int = 6
	draw int = 3
	lose int = 0
)

type PlayerMoves struct {
	opponentMove Score
	myMove       Score
}

func castToScore(str string) Score {
	switch str {
	case "A":
		return A
	case "B":
		return B
	case "C":
		return C
	case "X":
		return X
	case "Y":
		return Y
	case "Z":
		return Z
	default:
		return Score{} // just returned anything because yolo
	}
}

func main() {
	// open file
	file, err := os.Open("Q2.txt")
	if err != nil {
		log.Fatal("Error opening file")
	}

	// Read contents from file
	fileScanner := bufio.NewScanner(file)

	// Split file into lines
	fileScanner.Split(bufio.ScanLines)

	var sum int = 0

	for fileScanner.Scan() {
		var line string = fileScanner.Text()
		str := strings.Split(line, " ")
		var move PlayerMoves = PlayerMoves{castToScore(str[0]), pickMyMove(str[0], str[1])}
		sum += gameResult(move)

	}

	fmt.Println(sum)

}

func pickMyMove(opponentMove string, result string) Score {
	if result == "X" && opponentMove == "A" {
		return Z
	} // losing case
	if result == "X" && opponentMove == "B" {
		return X
	} // losing case
	if result == "X" && opponentMove == "C" {
		return Y
	} // losing case
	if result == "Y" && opponentMove == "A" {
		return X
	} // draw case
	if result == "Y" && opponentMove == "B" {
		return Y
	} // draw case
	if result == "Y" && opponentMove == "C" {
		return Z
	} // draw case
	if result == "Z" && opponentMove == "A" {
		return Y
	} // winning case
	if result == "Z" && opponentMove == "B" {
		return Z
	} // winning case
	if result == "Z" && opponentMove == "C" {
		return X
	} // winning case
	return Score{} // just returned anything because yolo
}

// Score  = Shape + outcome of round
func gameResult(move PlayerMoves) int {
	// Draw cases
	if move.opponentMove == A && move.myMove == X {
		return draw + A.score
	}
	if move.opponentMove == B && move.myMove == Y {
		return draw + B.score
	}
	if move.opponentMove == C && move.myMove == Z {
		return draw + C.score
	}

	// Losing cases
	if move.opponentMove == A && move.myMove == Z {
		return lose + Z.score
	}
	if move.opponentMove == B && move.myMove == X {
		return lose + X.score
	}
	if move.opponentMove == C && move.myMove == Y {
		return lose + Y.score
	}

	// Winning cases
	if move.opponentMove == A && move.myMove == Y {
		return win + Y.score
	}
	if move.opponentMove == B && move.myMove == Z {
		return win + Z.score
	}
	if move.opponentMove == C && move.myMove == X {
		return win + X.score
	}

	return 0 // just returned anything because yolo
}
