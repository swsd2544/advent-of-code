package day1

func SolveDeliveryFloor(puzzle string) int {
	currFloor := 0
	for _, char := range puzzle {
		switch char {
		case '(':
			currFloor++
		case ')':
			currFloor--
		}
	}

	return currFloor
}

func SolveFirstCharacterIdxToBasement(puzzle string) int {
	currFloor := 0
	for idx, char := range puzzle {
		switch char {
		case '(':
			currFloor++
		case ')':
			currFloor--
		}
		if currFloor < 0 {
			return idx + 1
		}
	}
	return 0
}
