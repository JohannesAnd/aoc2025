package day06

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := Part1("day06.txt")

	assert.Equal(t, 5595593539811, result)
}

func Test_part2(t *testing.T) {
	result := Part2("day06.txt")

	assert.Equal(t, 10153315705125, result)
}
