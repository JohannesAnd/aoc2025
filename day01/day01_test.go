package day01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_part1(t *testing.T) {
	result := part1()

	assert.Equal(t, 1064, result)
}

func Test_part2(t *testing.T) {
	result := part2()

	assert.Equal(t, 6122, result)
}
