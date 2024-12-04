package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isIncreasing(report []int32, dampener bool) bool {
	// check for increasing range with adjacent elements increased by at least one or at most three
	isDampened := false
	lastLevel := 0

	for level := 1; level < len(report); level++ {
		if report[level] > report[lastLevel] && (report[level]-report[lastLevel] >= 1) && (report[level]-report[lastLevel] <= 3) {
			lastLevel = level
			continue
		} else if dampener && !isDampened {
			isDampened = true

			if level+1 == len(report) { // case, when we reached the last level(index)
				continue
			}

			if level == 1 {
				if report[level] < report[level+1] && (report[level+1]-report[level] >= 1) && (report[level+1]-report[level] <= 3) {
					lastLevel = level
				} else {
					lastLevel = level - 1
				}
				continue
			}

			if report[level+1] > report[lastLevel] && (report[level+1]-report[lastLevel] >= 1) && (report[level+1]-report[lastLevel] <= 3) {
				lastLevel = level - 1 // to skip current level
				continue
			}
		} else {
			return false
		}
	}
	return true
}

func isDecreasing(report []int32, dampener bool) bool {
	// check for decreasing range with adjacent elements decreased by at least one or at most three
	isDampened := false
	lastLevel := 0

	for level := 1; level < len(report); level++ {
		if report[level] < report[lastLevel] && (report[lastLevel]-report[level] >= 1) && (report[lastLevel]-report[level] <= 3) {
			lastLevel = level
			continue
		} else if dampener && !isDampened {
			isDampened = true

			if level+1 == len(report) { // case, when we reached the last level(index)
				continue
			}

			if level == 1 {
				if (report[level] > report[level+1]) && (report[level]-report[level+1] >= 1) && (report[level]-report[level+1] <= 3) {
					lastLevel = level
				} else {
					lastLevel = level - 1
				}
				continue
			}

			if report[level+1] < report[level-1] && (report[level-1]-report[level+1] >= 1) && (report[level-1]-report[level+1] <= 3) {
				lastLevel = level - 1
				continue
			} // else {
			//	lastLevel = level
			//	continue
			//}
		} else {
			return false
		}
	}
	return true
}

func safeReports(reports [][]int32, dampener bool) int {

	// needs to be an either increasing or decreasing range
	// any two adjacent numbers must differ by at least one or at most three

	validReports := 0

	for _, report := range reports {
		if isIncreasing(report, dampener) || isDecreasing(report, dampener) {
			fmt.Println(report)
			validReports++
		}
	}
	return validReports

}

func main() {
	const filePath = "./input.txt"

	file, err := os.Open(filePath)
	check(err)

	var reports = make([][]int32, 1000)

	fileScanner := bufio.NewScanner(file)

	index := 0
	for fileScanner.Scan() {
		data := strings.Fields(fileScanner.Text())

		report := make([]int32, len(data))

		for i, num := range data {
			num, err := strconv.ParseInt(num, 10, 32)
			check(err)
			report[i] = int32(num)
		}

		reports[index] = report

		index++
	}
	err = file.Close()
	check(err)

	//fmt.Println(safeReports(reports, false))

	// part 2

	fmt.Println(safeReports(reports, true))

}
