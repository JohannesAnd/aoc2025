package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := Part1("day03.txt")

	assert.Equal(t, 16973, result)
}

func Test_part2(t *testing.T) {
	result := Part2("day03.txt")

	assert.Equal(t, 168027167146027, result)
}
