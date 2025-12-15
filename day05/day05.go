package day05

import (
	"aoc2025/utils"
	"strconv"
	"strings"
)

type numberRange struct {
	start int
	end   int
}

func (numberRange *numberRange) contains(val int) bool {
	return val >= numberRange.start && val <= numberRange.end
}

func (numberRange *numberRange) containsRange(val numberRange) bool {
	return val.start > numberRange.start && val.end < numberRange.end
}

func (numberRange *numberRange) size() int {
	return numberRange.end - numberRange.start + 1
}

func mergeRanges(ranges ...numberRange) numberRange {
	start := ranges[0].start
	end := ranges[0].end

	for _, o := range ranges {
		if o.start < start {
			start = o.start
		}

		if o.end > end {
			end = o.end
		}
	}

	return numberRange{start: start, end: end}
}

func findSubRange(ranges map[numberRange]struct{}, value numberRange) []numberRange {
	res := make([]numberRange, 0)

	for r := range ranges {
		if value.containsRange(r) {
			res = append(res, r)
		}
	}

	return res
}

func findRange(ranges map[numberRange]struct{}, value int) (numberRange, bool) {
	for r := range ranges {
		if r.contains(value) {
			return r, true
		}
	}

	return numberRange{0, 0}, false
}

func createRange(str string) numberRange {
	parts := strings.Split(str, "-")

	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])

	return numberRange{start: start, end: end}
}

func Part1(inputPath string) int {
	lr, _ := utils.NewLineReader(inputPath)

	count := 0
	rangesDone := false

	ranges := make(map[numberRange]struct{}, 0)

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		if len(line) == 0 {
			rangesDone = true
			continue
		}

		if rangesDone {
			number, _ := strconv.Atoi(line)
			_, found := findRange(ranges, number)

			if found {
				count++
			}
		} else {
			// Add range
			r := createRange(line)

			for _, useless := range findSubRange(ranges, r) {
				delete(ranges, useless)
			}

			startOverlap, foundStart := findRange(ranges, r.start)
			endOverlap, foundEnd := findRange(ranges, r.end)

			if foundStart && foundEnd {
				delete(ranges, startOverlap)
				delete(ranges, endOverlap)

				newRange := mergeRanges(r, startOverlap, endOverlap)
				ranges[newRange] = struct{}{}
			} else if foundStart {
				delete(ranges, startOverlap)

				newRange := mergeRanges(r, startOverlap)
				ranges[newRange] = struct{}{}
			} else if foundEnd {
				delete(ranges, endOverlap)

				newRange := mergeRanges(r, endOverlap)
				ranges[newRange] = struct{}{}
			} else {
				ranges[r] = struct{}{}
			}
		}
	}

	return count
}

func Part2(inputPath string) int {
	lr, _ := utils.NewLineReader(inputPath)

	ranges := make(map[numberRange]struct{}, 0)

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		if len(line) == 0 {
			break
		}

		r := createRange(line)

		for _, useless := range findSubRange(ranges, r) {
			delete(ranges, useless)
		}

		startOverlap, foundStart := findRange(ranges, r.start)
		endOverlap, foundEnd := findRange(ranges, r.end)

		if foundStart && foundEnd {
			delete(ranges, startOverlap)
			delete(ranges, endOverlap)

			newRange := mergeRanges(r, startOverlap, endOverlap)
			ranges[newRange] = struct{}{}
		} else if foundStart {
			delete(ranges, startOverlap)

			newRange := mergeRanges(r, startOverlap)
			ranges[newRange] = struct{}{}
		} else if foundEnd {
			delete(ranges, endOverlap)

			newRange := mergeRanges(r, endOverlap)
			ranges[newRange] = struct{}{}
		} else {
			ranges[r] = struct{}{}
		}
	}

	count := 0

	for r := range ranges {
		count += r.size()
	}

	return count
}
