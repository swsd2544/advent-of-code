package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"advent-of-code-2015/day1"
	"advent-of-code-2015/day2"
	"advent-of-code-2015/day3"
)

func main() {
	var day int
	flag.IntVar(&day, "day", 0, "day number to run puzzle input")

	flag.Parse()

	if day <= 0 {
		log.Fatalf("expected day equal or higher than one. got=%d", day)
	}

	puzzleInput, err := readPuzzleInput(day)
	if err != nil {
		log.Fatalf("failed to extract puzzle input string: %s", err)
	}

	switch day {
	case 1:
		deliveryFloor := day1.SolveDeliveryFloor(puzzleInput)
		characterIdx := day1.SolveFirstCharacterIdxToBasement(puzzleInput)
		log.Printf("From puzzle input, delivery floor is %d", deliveryFloor)
		log.Printf("From puzzle input, first character idx to go to basement is %d", characterIdx)
	case 2:
		totalPaperArea := day2.SolveTotalWrappingArea(puzzleInput)
		totalRibbonLength := day2.SolveTotalRibbonLength(puzzleInput)
		log.Printf("From puzzle input, total wrapping paper needed is %d", totalPaperArea)
		log.Printf("From puzzle input, total ribbon length needed is %d", totalRibbonLength)
	case 3:
		totalHouses := day3.SolveNumberOfHousesWithPresents(puzzleInput, 1)
		totalHousesWithRobo := day3.SolveNumberOfHousesWithPresents(puzzleInput, 2)
		log.Printf("From puzzle input, total houses with atleast one present is %d", totalHouses)
		log.Printf("From puzzle input, total houses with atleast one present (two workers) is %d", totalHousesWithRobo)
	default:
		log.Fatalf("the day %d is not implemented yet.", day)
	}
}

func readPuzzleInput(day int) (string, error) {
	puzzleInputFilepath := fmt.Sprintf("./day%d/puzzle.txt", day)
	puzzleInput, err := os.ReadFile(puzzleInputFilepath)
	if err != nil {
		return "", fmt.Errorf("failed to read puzzle input file %s: %w", puzzleInputFilepath, err)
	}
	return string(puzzleInput), nil
}
