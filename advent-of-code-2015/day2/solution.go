package day2

import (
	"bufio"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolveTotalWrappingArea(puzzle string) int {
	totalPaper := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzle))
	for scanner.Scan() {
		area, err := SolveWrappingArea(scanner.Text())
		if err != nil {
			fmt.Printf("failed to solve total wrapping area: %s\n", err)
			continue
		}
		totalPaper += area
	}

	return totalPaper
}

func SolveWrappingArea(puzzle string) (int, error) {
	dimensions := make([]int, 0, 3)
	for _, char := range strings.Split(puzzle, "x") {
		d, err := strconv.Atoi(char)
		if err != nil {
			return 0, fmt.Errorf("failed to parse dimension: %w", err)
		}
		dimensions = append(dimensions, d)
	}

	if len(dimensions) != 3 {
		return 0, fmt.Errorf("invalid input, expected len=3. got=%d", len(dimensions))
	}

	slices.Sort(dimensions)
	totalSurfaceArea := 2 * ((dimensions[0] * dimensions[1]) + (dimensions[1] * dimensions[2]) + (dimensions[0] * dimensions[2]))
	smallestArea := dimensions[0] * dimensions[1]
	return totalSurfaceArea + smallestArea, nil
}

func SolveTotalRibbonLength(puzzle string) int {
	totalLength := 0

	scanner := bufio.NewScanner(strings.NewReader(puzzle))
	for scanner.Scan() {
		length, err := SolveRibbonLength(scanner.Text())
		if err != nil {
			fmt.Printf("failed to solve total wrapping area: %s\n", err)
			continue
		}
		totalLength += length
	}

	return totalLength
}

func SolveRibbonLength(puzzle string) (int, error) {
	dimensions := make([]int, 0, 3)
	for _, char := range strings.Split(puzzle, "x") {
		d, err := strconv.Atoi(char)
		if err != nil {
			return 0, fmt.Errorf("failed to parse dimension: %w", err)
		}
		dimensions = append(dimensions, d)
	}

	if len(dimensions) != 3 {
		return 0, fmt.Errorf("invalid input, expected len=3. got=%d", len(dimensions))
	}

	slices.Sort(dimensions)
	smallestPerimeter := 2 * (dimensions[0] + dimensions[1])
	cubicVolume := dimensions[0] * dimensions[1] * dimensions[2]
	return smallestPerimeter + cubicVolume, nil
}
