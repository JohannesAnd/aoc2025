package day05

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := Part1()

	assert.Equal(t, 694, result)
}

func Test_part2(t *testing.T) {
	result := Part2()

	assert.Equal(t, 352716206375547, result)
}
