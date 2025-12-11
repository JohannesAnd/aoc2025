package day01

import (
	"aoc2025/utils"
	"strconv"
	"strings"
	"unicode"
)

func stringsToInts(input []string) []int {
	res := make([]int, len(input))

	for i, s := range input {
		res[i], _ = strconv.Atoi(s)
	}

	return res
}

func sumInts(input []int) int {
	sum := 0
	for _, i := range input {
		sum += i
	}
	return sum
}

func Part1() int {
	lines, _ := utils.ReadFile("day06.txt")

	operators := strings.Fields(lines[len(lines)-1])
	results := stringsToInts(strings.Fields(lines[0]))

	for _, line := range lines[1 : len(lines)-1] {
		elements := stringsToInts(strings.Fields(line))

		for elementIndex, element := range elements {
			if operators[elementIndex] == "*" {
				results[elementIndex] = results[elementIndex] * element
			}
			if operators[elementIndex] == "+" {
				results[elementIndex] = results[elementIndex] + element
			}
		}
	}

	return sumInts(results)
}

func findNonSpaceIndecies(str string) []int {
	res := make([]int, 0)

	for i, s := range str {
		if !unicode.IsSpace(s) {
			res = append(res, i)
		}
	}

	res = append(res, len(str)+10)

	return res
}

func getColumnNumber(lines []string, index int) (int, error) {
	res := ""

	for _, line := range lines[0 : len(lines)-1] {
		if index >= len(line) {
			continue
		}

		char := line[index]
		if !unicode.IsSpace(rune(char)) {
			res += string(char)
		}
	}

	return strconv.Atoi(res)
}

func Part2() int {
	lines, _ := utils.ReadFile("day06.txt")

	chunkIndecies := findNonSpaceIndecies(lines[len(lines)-1])

	sum := 0

	for i := range chunkIndecies[0 : len(chunkIndecies)-1] {
		from := chunkIndecies[i]
		to := chunkIndecies[i+1]

		operator := lines[len(lines)-1][from]

		val := 0

		if operator == '*' {
			val = 1
		}

		for p := from; p < to-1; p++ {
			columnValue, _ := getColumnNumber(lines, p)

			if operator == '*' {
				val *= columnValue
			}
			if operator == '+' {
				val += columnValue
			}
		}

		sum += val
	}

	return sum
}
