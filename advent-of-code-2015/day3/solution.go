package day3

type point struct {
	x, y int
}

func (p *point) add(ap point) {
	p.x += ap.x
	p.y += ap.y
}

func SolveNumberOfHousesWithPresents(puzzle string, workers int) int {
	if workers <= 0 {
		return 0
	}

	houses := make(map[point]struct{})

	currPoints := make([]point, workers)
	houses[point{0, 0}] = struct{}{}

	for idx, char := range puzzle {
		movement := decodeMovementUnit(char)
		currPoints[idx%len(currPoints)].add(movement)
		houses[currPoints[idx%len(currPoints)]] = struct{}{}
	}

	return len(houses)
}

func decodeMovementUnit(char rune) point {
	var p point
	switch char {
	case '^':
		p.y = 1
	case 'v':
		p.y = -1
	case '>':
		p.x = 1
	case '<':
		p.x = -1
	}
	return p
}
