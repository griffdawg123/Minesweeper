package main

import (
	"fmt"
	"strconv"
)

func main() {
	var size int
	var numMines int
	size, numMines = initFlags(size, numMines)

	var board [][]int = initGameBoard(size)

	fmt.Println("Size: ", strconv.Itoa(size), "Num Mines: %v", strconv.Itoa(numMines))
	debug_print_gameboard(board)
}
