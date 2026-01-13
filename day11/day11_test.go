package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := Part1("day11.txt")

	assert.Equal(t, 603, result)
}

func Test_part2(t *testing.T) {
	result := Part2("day11.txt")

	assert.Equal(t, 2, result)
}
