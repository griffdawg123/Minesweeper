package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var size int
	var numMines int
	size, numMines = initFlags(size, numMines)

	var board [][]int
	var err error
	board, err = initGameBoard(size)
	if err != nil {
		panic(err)
	}
	fmt.Println("Size: ", strconv.Itoa(size), "Num Mines:", strconv.Itoa(numMines))
	var mineCoords []*coords
	mineCoords, err = getMineCoords(numMines, os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(mineCoords)
	debug_print_gameboard(board)
}
