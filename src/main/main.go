package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var height int
	var width int
	var numMines int
	height, width, numMines = initFlags(height, width, numMines)

	var board [][]int
	var err error
	board, err = initGameBoard(height, width)
	if err != nil {
		panic(err)
	}
    fmt.Println("Height: ", strconv.Itoa(height), "Width: ", strconv.Itoa(width), "Num Mines:", strconv.Itoa(numMines))
	var mineCoords []*coords
	mineCoords, err = getMineCoords(numMines, os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(mineCoords)
    board, err = placeMines(mineCoords, board)
    if err != nil {
        panic(err)
    }
	err = debug_print_gameboard(board)
    if err != nil {
        panic(err)
    }
    board, err = revealSquare(board, &coords{3,3})
	err = debug_print_gameboard(board)
    if err != nil {
        panic(err)
    }
}

