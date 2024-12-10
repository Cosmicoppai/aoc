package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func validPass(row int, col int, puzzle []string) int32 {
	searchFor := "XMAS"
	var total int32

	searchIdx := 0

	resetSearchIdx := func() {
		if searchIdx == len(searchFor) {
			total++
		}
		searchIdx = 0
	}

	// search horizontally (across column)
	for i := col; i < col+4; i++ {
		if i >= len(puzzle[row]) {
			break
		}
		if puzzle[row][i] != searchFor[searchIdx] {
			break
		}
		searchIdx++
	}

	resetSearchIdx()

	for i := col; i > col-4; i-- {
		if i < 0 {
			break
		}

		if puzzle[row][i] != searchFor[searchIdx] {
			break
		}
		searchIdx++
	}

	resetSearchIdx()

	// search vertically (across row)
	for i := row; i < row+4; i++ {
		if i >= len(puzzle) {
			break
		}

		if puzzle[i][col] != searchFor[searchIdx] {
			break
		}
		searchIdx++
	}

	resetSearchIdx()

	for i := row; i > row-4; i-- {
		if i < 0 {
			break
		}

		if puzzle[i][col] != searchFor[searchIdx] {
			break
		}
		searchIdx++
	}

	resetSearchIdx()

	// search diagonally backward(\)

	currRow, currCol := row, col

	for currRow < len(puzzle) && currCol < len(puzzle[row]) {
		if puzzle[currRow][currCol] != searchFor[searchIdx] {
			break
		}

		currRow++
		currCol++
		searchIdx++

		if searchIdx == len(searchFor) {
			break
		}
	}

	resetSearchIdx()

	currRow, currCol = row, col

	for currRow >= 0 && currCol >= 0 {
		if puzzle[currRow][currCol] != searchFor[searchIdx] {
			break
		}

		currRow--
		currCol--
		searchIdx++

		if searchIdx == len(searchFor) {
			break
		}
	}

	resetSearchIdx()

	// search diagonally forward (/)

	currRow, currCol = row, col

	for currRow >= 0 && currCol < len(puzzle[row]) {
		if puzzle[currRow][currCol] != searchFor[searchIdx] {
			break
		}

		currRow--
		currCol++
		searchIdx++

		if searchIdx == len(searchFor) {
			break
		}
	}

	resetSearchIdx()

	currRow, currCol = row, col

	for currRow < len(puzzle) && currCol >= 0 {
		if puzzle[currRow][currCol] != searchFor[searchIdx] {
			break
		}

		currRow++
		currCol--
		searchIdx++

		if searchIdx == len(searchFor) {
			break
		}
	}

	resetSearchIdx()

	return total

}

func xmas(row int, col int, puzzle []string) int32 {
	searchFor := "MAS"
	searchFor2 := "SAM"
	var total int32

	searchIdx := 0

	resetSearchIdx := func() {
		if searchIdx == len(searchFor) {
			total++
		}
		searchIdx = 0
	}

	// search diagonally backward(\)

	currRow, currCol := row, col
	var searchVar string

	if puzzle[currRow][currCol] == 'M' {
		searchVar = searchFor
	} else {
		searchVar = searchFor2
	}

	for currRow < len(puzzle) && currCol < len(puzzle[row]) {
		if puzzle[currRow][currCol] != searchVar[searchIdx] {
			break
		}

		currRow++
		currCol++
		searchIdx++

		if searchIdx == len(searchFor) {
			break
		}
	}

	resetSearchIdx()

	// search diagonally forward (/)

	if col+2 >= len(puzzle[row]) || total == 0 {
		return 0
	}

	col = col + 2

	currRow, currCol = row, col

	if puzzle[currRow][currCol] == 'M' {
		searchVar = searchFor
	} else {
		searchVar = searchFor2
	}

	for currRow < len(puzzle) && currCol >= 0 {
		if puzzle[currRow][currCol] != searchVar[searchIdx] {
			break
		}

		currRow++
		currCol--
		searchIdx++

		if searchIdx == len(searchFor) {
			break
		}
	}

	resetSearchIdx()

	return total / 2

}

func main() {

	const filePath = "C:/Users/cosmi/GolandProjects/aoc/4/input.txt"

	file, err := os.Open(filePath)
	check(err)

	scanner := bufio.NewScanner(file)

	var puzzleMatrix []string

	for scanner.Scan() {
		puzzleMatrix = append(puzzleMatrix, scanner.Text())
	}

	var total int32

	// part 1
	for i := 0; i < len(puzzleMatrix); i++ {
		for j := 0; j < len(puzzleMatrix[i]); j++ {
			count := validPass(i, j, puzzleMatrix)
			total += count
		}
	}

	fmt.Println(total)

	// part 2

	total = 0
	for i := 0; i < len(puzzleMatrix); i++ {
		for j := 0; j < len(puzzleMatrix[i]); j++ {
			count := xmas(i, j, puzzleMatrix)
			total += count
		}
	}

	fmt.Println(total)

}
