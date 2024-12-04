package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func scanCorruptMemory(text string, pattern string) [][2]int32 {

	// will return the stuff in between the pattern

	r, err := regexp.Compile(pattern)
	check(err)

	var output [][2]int32

	finds := r.FindAllStringSubmatch(text, -1)
	for _, value := range finds {
		num1, err := strconv.ParseInt(value[1], 10, 32)
		check(err)
		num2, err := strconv.ParseInt(value[2], 10, 32)
		check(err)

		output = append(output, [2]int32{int32(num1), int32(num2)})
	}
	return output
}

func mul(numbers [][2]int32) int32 {
	var total int32 = 0
	for _, nums := range numbers {
		total += nums[0] * nums[1]
	}
	return total
}

func main() {

	const filePath = "./input.txt"

	file, err := os.Open(filePath)
	check(err)

	scanner := bufio.NewScanner(file)

	var builder strings.Builder

	for scanner.Scan() {
		builder.WriteString(scanner.Text())
	}

	text := builder.String()

	var total int32 = 0

	const (
		startInstruction = "do()"
		endInstruction   = "don't()"
		operationPattern = `mul\((\d+),(\d+)\)`
	)

	text = "do()" + text // to mimic the setting of do() instruction initially

	for len(text) > 0 {

		start, end := -1, -2 // reset the pointers (arbitrary numbers, to initiate index calculations)

		for end < start {
			start = strings.Index(text, startInstruction) // find the "do()" in the string, it'll give an index from where we can start considering the operations

			if start == -1 {
				fmt.Println(total)
				return // when "do()" is not found, it means, the operation will never resume
			}
			end = strings.Index(text, endInstruction) // find "don't()" in the string, it'll give an index until we can consider the operations
			if end == -1 {
				end = len(text) // when "don't()" is not found, it means, the operation will never stop
				break
			} else if end < start {
				text = text[end+len(endInstruction):] // we're still behind start, search for the next stop instruction
			}
		}

		numArray := scanCorruptMemory(text[start:end], operationPattern) // find valid operations within the conditional instructions
		total += mul(numArray)

		text = text[end+len(endInstruction):] // get rid of the processed text

	}
	fmt.Println(total)
}
