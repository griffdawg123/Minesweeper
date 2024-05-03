package main

import (
	"flag"
	"fmt"
)

func initFlags(size int, numMines int) (int, int) {
	flag.IntVar(&size, "size", 10, "Create an nxn grid with n=size")
	flag.IntVar(&numMines, "mines", 10, "Number of mines in the game")
	flag.Parse()
	return size, numMines
}

func initGameBoard(size int) [][]int {
	var gameboard [][]int = make([][]int, size)
	for i := range gameboard {
		gameboard[i] = make([]int, size)
	}
	return gameboard
}

func debug_print_gameboard(board [][]int) {
	for i := range board {
		var row []int = board[i]
		for j := range row {
			fmt.Printf("%v ", row[j])
		}
		fmt.Println()
	}
}
