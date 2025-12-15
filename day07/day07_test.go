package day07

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := Part1()

	assert.Equal(t, 1518, result)
}

func Test_part2(t *testing.T) {
	result := Part2()

	assert.Equal(t, 25489586715621, result)
}
