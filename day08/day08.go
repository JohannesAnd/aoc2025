package day08

import (
	"aoc2025/utils"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Box struct {
	x, y, z int
}

type Pair struct {
	a Box
	b Box
}

func (receiver Pair) distance() float64 {
	return math.Sqrt(
		math.Pow(math.Abs(float64(receiver.a.x-receiver.b.x)), 2) +
			math.Pow(math.Abs(float64(receiver.a.y-receiver.b.y)), 2) +
			math.Pow(math.Abs(float64(receiver.a.z-receiver.b.z)), 2))
}

func Part1(inputPath string) int {
	lr, _ := utils.NewLineReader(inputPath)

	boxes := make([]Box, 20)
	index := 0

	for {
		line, ok := lr.Next()

		if !ok {
			break
		}

		coords := strings.Split(line, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		boxes[index] = Box{x, y, z}

		index++

	}

	pairs := make(map[float64]Pair)

	for y := 0; y < len(boxes); y++ {
		for x := y + 1; x < len(boxes); x++ {
			pair := Pair{boxes[y], boxes[x]}
			distance := pair.distance()

			_, ok := pairs[distance]

			if ok {
				panic("Duplicate pair")
			}

			pairs[distance] = pair
		}
	}

	keys := slices.Sorted(maps.Keys(pairs))

	return len(keys)

}

func Part2(inputPath string) int {
	return -1
}
