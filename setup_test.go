package main

import (
	"bufio"
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameBoardSetup(t *testing.T) {
	var size int = 5
	var board [][]int
	board, _ = initGameBoard(size)
	if len(board) != 5 {
		t.Errorf("Expected row count: %v, got: %v", size, len(board))
	}
	if len(board[0]) != 5 {
		t.Errorf("Expected row count: %v, got: %v", size, len(board[0]))
	}
}

func TestGameBoardNegative(t *testing.T) {
	var err error
	var size int = -1
	_, err = initGameBoard(size)
	if err == nil {
		t.Errorf("Negative size game board should return error. Size: %v", size)
	}
}

func TestGameboardTooBig(t *testing.T) {
	var err error
	var size int = 50
	_, err = initGameBoard(size)
	if err == nil {
		t.Errorf("Gameboard size larger than %v should return an error. Size: %v", MAXBOARDSIZE, size)
	}
}

func mockStdinTesting(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return string(text), err
}

func TestMockStdinTesting(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("Hello, World!\n"))
	result, err := mockStdinTesting(&stdin)
	assert.NoError(t, err)
	assert.Equal(t, "Hello, World!\n", result)
}

func TestGetMineLocations(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("1 1\n2 2\n3 3\n"))
	result, err := getMineCoords(3, &stdin)
	assert.NoError(t, err)
	var expected []*coords = []*coords{{1, 1}, {2, 2}, {3, 3}}
	assert.Equal(t, len(result), len(expected))
	for i := range expected {
		assert.Equal(t, result[i], expected[i])
	}
}
