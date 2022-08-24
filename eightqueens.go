package piscine

import (
	"github.com/01-edu/z01"
)

func EightQueens() {
	var positions [8][8]int
	var busyPositions [8][8]int

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			busyPositions[i][j] = 0
			positions[i][j] = 0
		}
	}
	count := 0
	SolvedQeen(&busyPositions, 0, &count)
}

func SolvedQeen(busyPositions *[8][8]int, n int, count *int) bool {
	extra := busyPositions
	if *count == 8 {
		return true
	}
	if n == 8 {
		return true
	}
	for j := 0; j < 8; j++ {
		if busyPositions[n][j] == 0 {
			if addBusyPositions(busyPositions, n, j) {
				if SolvedQeen(busyPositions, n+1, count) {
					for i := 0; i < 8; i++ {
						for j := 0; j < 8; j++ {
							if busyPositions[i][j] == i+11 {
								z01.PrintRune(rune(j+1) + 48)
							}
						}
					}
					z01.PrintRune(10)
					busyPositions = extra
				}
			}
			removeBusyPositions(busyPositions, n, j)
			*count--
		}
	}
	return false
}

func addBusyPositions(busyPositions *[8][8]int, first, second int) bool {
	// pluss position
	if busyPositions[first][second] != 0 {
		return false
	}
	for j := 0; j < 8; j++ {
		if busyPositions[first][j] == 0 {
			busyPositions[first][j] = first + 1
		}
		if busyPositions[j][second] == 0 {
			busyPositions[j][second] = first + 1
		}

	}
	// first /
	if first+second <= 7 {
		for j := 0; j <= second+first; j++ {
			if busyPositions[j][(second+first)-j] == 0 {
				busyPositions[j][(second+first)-j] = first + 1
			}
		}
	} else {
		for i := (first + second) - 7; i < 8; i++ {
			if busyPositions[i][(first+second)-i] == 0 {
				busyPositions[i][(first+second)-i] = first + 1
			}
		}
	}

	// second dioganal
	if first >= second {
		for i := first - second; i < 8; i++ {
			if busyPositions[i][i-(first-second)] == 0 {
				busyPositions[i][i-(first-second)] = first + 1
			}
		}
	} else if first < second {
		for i := second - first; i < 8; i++ {
			if busyPositions[i-(second-first)][i] == 0 {
				busyPositions[i-(second-first)][i] = first + 1
			}
		}
	}
	busyPositions[first][second] = first + 1 + 10
	return true
}

func removeBusyPositions(busyPositions *[8][8]int, first, second int) bool {
	// pluss position
	if busyPositions[first][second] == 0 {
		return false
	}
	for j := 0; j < 8; j++ {
		if busyPositions[first][j] == first+1 {
			busyPositions[first][j] = 0
		}
		if busyPositions[j][second] == first+1 {
			busyPositions[j][second] = 0
		}

	}
	// first
	if first+second <= 7 {
		for j := 0; j <= second+first; j++ {
			if busyPositions[j][(second+first)-j] == first+1 {
				busyPositions[j][(second+first)-j] = 0
			}
		}
	} else {
		for i := (first + second) - 7; i < 8; i++ {
			if busyPositions[i][(first+second)-i] == first+1 {
				busyPositions[i][(first+second)-i] = 0
			}
		}
	}

	// second dioganal
	if first >= second {
		for i := first - second; i < 8; i++ {
			if busyPositions[i][i-(first-second)] == first+1 {
				busyPositions[i][i-(first-second)] = 0
			}
		}
	} else if first < second {
		for i := second - first; i < 8; i++ {
			if busyPositions[i-(second-first)][i] == first+1 {
				busyPositions[i-(second-first)][i] = 0
			}
		}
	}
	busyPositions[first][second] = 0
	return true
}
