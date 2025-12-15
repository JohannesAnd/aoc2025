package day04

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getValue_1(t *testing.T) {
	d := []byte{1, 2, 10, 4, 5, 10, 7, 8, 10}

	val, found := getValue(d, 2, 0, 1)

	assert.Equal(t, true, found)
	assert.Equal(t, byte(2), val)
}

func Test_getValue_2(t *testing.T) {
	d := []byte{1, 2, 10, 4, 5, 10, 7, 8, 10}

	val, found := getValue(d, 2, 1, 1)

	assert.Equal(t, true, found)
	assert.Equal(t, byte(5), val)
}

func Test_getValue_3(t *testing.T) {
	d := []byte{1, 2, 10, 4, 5, 10, 7, 8, 10}

	_, found := getValue(d, 2, -1, -1)

	assert.Equal(t, false, found)
}

func Test_getValue_4(t *testing.T) {
	d := []byte{1, 2, 10, 4, 5, 10, 7, 8, 10}

	val, found := getValue(d, 2, 2, 1)

	assert.Equal(t, true, found)
	assert.Equal(t, byte(8), val)
}

func Test_part1(t *testing.T) {
	result := Part1("day04.txt")

	assert.Equal(t, 1495, result)
}

func Test_part2(t *testing.T) {
	result := Part2("day04.txt")

	assert.Equal(t, 8768, result)
}
