package day02

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := Part1("day02.txt")

	assert.Equal(t, 28844599675, result)
}

func Test_part2(t *testing.T) {
	result := Part2("day02.txt")

	assert.Equal(t, 48778605167, result)
}
