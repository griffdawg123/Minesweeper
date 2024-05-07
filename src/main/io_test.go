package main

import (
    "testing"
    "bytes"

	"github.com/stretchr/testify/assert"
)

func TestGetMineLocations(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("1 1\n2 2\n3 3\n"))
	result, err := getMineCoords(3, &stdin)
	assert.Nil(t, err)
	var expected []*coords = []*coords{{1, 1}, {2, 2}, {3, 3}}
	assert.Equal(t, len(result), len(expected))
	for i := range expected {
		assert.Equal(t, result[i], expected[i])
	}
}

func TestDebugPrint(t *testing.T) {
	var board [][]int
	var err error
	var size int = 5
	board, err = initGameBoard(size, size)
	assert.Nil(t, err)
	output, err := captureOutput(func() error {
		err := debug_print_gameboard(board)
		return err
	})
	assert.Nil(t, err)
	assert.Equal(t, "1 1 1 1 1 \n1 1 1 1 1 \n1 1 1 1 1 \n1 1 1 1 1 \n1 1 1 1 1 \n", output)
}


