package day08

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := Part1("day08.txt")

	assert.Equal(t, 1518, result)
}

func Test_part2(t *testing.T) {
	result := Part2("day08.txt")

	assert.Equal(t, 25489586715621, result)
}
