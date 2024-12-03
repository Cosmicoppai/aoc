package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func part1(arr1 []int32, arr2 []int32, index int) {

	slices.Sort(arr1)
	slices.Sort(arr2)

	var totalDistance float64 = 0

	index--
	for index >= 0 {
		totalDistance += math.Abs(float64(arr1[index] - arr2[index]))
		index--
	}

	fmt.Println("Total Distance:", strconv.FormatFloat(totalDistance, 'f', 0, 64))
}

func part2(arr1 []int32, countMap map[int32]int32) {

	similarity := float64(0)

	for i := 0; i < len(arr1); i++ {
		similarity += float64(arr1[i] * countMap[arr1[i]])
	}

	fmt.Println("Similarity:", strconv.FormatFloat(similarity, 'f', 0, 64))

}

func main() {
	const filePath = "./input.txt"

	file, err := os.Open(filePath)
	check(err)

	var arr1 []int32
	var arr2 []int32

	fileScanner := bufio.NewScanner(file)

	index := 0
	for fileScanner.Scan() {
		data := strings.Fields(fileScanner.Text())

		num1, err := strconv.ParseInt(data[0], 10, 32)
		check(err)
		num2, err := strconv.ParseInt(data[1], 10, 32)
		check(err)

		arr1 = append(arr1, int32(num1))
		arr2 = append(arr2, int32(num2))

		index++
	}
	err = file.Close()
	check(err)

	part1(arr1, arr2, index)

	// part 2

	countMap := make(map[int32]int32)

	for i := 0; i < len(arr2); i++ {
		if _, ok := countMap[arr2[i]]; ok {
			countMap[arr2[i]]++
		} else {
			countMap[arr2[i]] = 1
		}
	}

	part2(arr1, countMap)

}
