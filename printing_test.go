package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDebugPrint(t *testing.T) {
	var board [][]int
	var err error
	var size int = 5
	board, err = initGameBoard(size)
	assert.Nil(t, err)
	output, err := captureOutput(func() error {
		err := debug_print_gameboard(board)
		return err
	})
	assert.Nil(t, err)
	assert.Equal(t, "1 1 1 1 1 \n1 1 1 1 1 \n1 1 1 1 1 \n1 1 1 1 1 \n1 1 1 1 1 \n", output)
}
