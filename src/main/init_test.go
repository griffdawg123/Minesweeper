package main

import (
    "testing"
    "bytes"
	"github.com/stretchr/testify/assert"
)

func TestGameBoardSetup(t *testing.T) {
	var size int = 5
	var board [][]int
	board, _ = initGameBoard(size, size)
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
	_, err = initGameBoard(size, size)
	assert.Error(t, err)
}

func TestGameboardTooBig(t *testing.T) {
	var err error
	var size int = 50
	_, err = initGameBoard(size, size)
	assert.Error(t, err)
}

func TestMockStdinTesting(t *testing.T) {
	var stdin bytes.Buffer

	stdin.Write([]byte("Hello, World!\n"))
	result, err := mockStdinTesting(&stdin)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, World!\n", result)
}

func TestPlaceMines(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

	var mineCoords []*coords = []*coords{{1, 1}, {2, 2}, {3, 3}}
    newBoard, err := placeMines(mineCoords, board) 
    assert.Nil(t, err)

    assert.Equal(t, len(newBoard), 5)
    assert.Equal(t, len(newBoard[0]), 5)

    var expectedBoard = [][]int{{1, 1, 1, 1, 1},{1, 2, 1, 1, 1},{1, 1, 2, 1, 1},{1, 1, 1, 2, 1},{1, 1, 1, 1, 1}}
    for row := range newBoard {
        for col := range newBoard[row] {
            assert.Equal(t, newBoard[row][col], expectedBoard[row][col])
        }
    }
}

func TestPlaceMinesOutside(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

	var mineCoords []*coords = []*coords{{1, 1}, {2, 2}, {5, 5}}
    newBoard, err := placeMines(mineCoords, board) 
    assert.Nil(t, err)

    assert.Equal(t, len(newBoard), 5)
    assert.Equal(t, len(newBoard[0]), 5)

    var expectedBoard = [][]int{{1, 1, 1, 1, 1},{1, 2, 1, 1, 1},{1, 1, 2, 1, 1},{1, 1, 1, 1, 1},{1, 1, 1, 1, 1}}
    for row := range newBoard {
        for col := range newBoard[row] {
            assert.Equal(t, newBoard[row][col], expectedBoard[row][col])
        }
    }
}


