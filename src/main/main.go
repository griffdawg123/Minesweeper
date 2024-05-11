package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var height, width, numMines int 
    height, width, numMines = initFlags(height, width, numMines)

    board, err := initGameBoard(height, width)
	if err != nil {
		panic(err)
	}

    fmt.Println("Height: ", strconv.Itoa(height), "Width: ", strconv.Itoa(width), "Num Mines:", strconv.Itoa(numMines))
    mineCoords := []*coords{} 
    for i := 0; i < numMines; i++ {
        coord, err := getMineCoords(os.Stdin)
        if err != nil {
            panic(err)
        }
        mineCoords = append(mineCoords, coord)
    }
	
    board = placeMines(mineCoords, board)
	debug_print_gameboard(board)
    // for !gameWon(board) {
        
    // }
    board, err = revealSquare(board, &coords{3,3})
    printGameboard(board)
}

