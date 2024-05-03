package main

import "testing"

func TestGameBoardSetup(t *testing.T) {
	var size int = 5
	var board [][]int = initGameBoard(size)
	if len(board) != 5 {
		t.Errorf("Expected row count: %v, got: %v", size, len(board))
	}
	if len(board[0]) != 5 {
		t.Errorf("Expected row count: %v, got: %v", size, len(board[0]))
	}
}
