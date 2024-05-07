package main

import "flag"
import "errors"

func initFlags(height int, width int, numMines int) (int, int, int) {
	flag.IntVar(&height, "size", 10, "Height of the grid")
	flag.IntVar(&width, "size", 10, "Width of the grid")
	flag.IntVar(&numMines, "mines", 10, "Number of mines in the game")
	flag.Parse()
	return height, width, numMines
}

func initGameBoard(height int, width int) ([][]int, error) {
	if height <= 0 || width <= 0 {
		return nil, errors.New("size must not be negative")
	} else if height > MAXBOARDSIZE  || width > MAXBOARDSIZE {
		return nil, errors.New("maximum game size is 20 cells")
	}
	var gameboard [][]int = make([][]int, height)
	for i := range gameboard {
		gameboard[i] = make([]int, width)
		for j := range gameboard[i] {
			gameboard[i][j] = HIDDEN_SAFE 
		}
	}
	return gameboard, nil
}
func placeMines(mines []*coords, board [][]int) ([][]int, error) {
    var height int = len(board)
    var width int = len(board[0])

    for i := range mines {
        var mineCoords *coords = mines[i]
        var row = mineCoords.row
        var col = mineCoords.col
        if row < height && col < width {
            board[row][col] = HIDDEN_MINE
        }
    }
    return board, nil
}

func adjacentMines(board [][]int, square *coords) (int, error) {
    if square.row > len(board) || square.row < 0 {
        return -1, errors.New("Row out of bounds")
    } else if square.col > len(board[0]) || square.col < 0 {
        return -1, errors.New("Col out of bounds")
    }

    var count int = 0
    for row := max(square.row-1, 0); row < min(square.row+2, len(board)); row++ {
        for col := max(square.col-1, 0); col < min(square.col+2, len(board[0])); col++ {
            if board[row][col] == HIDDEN_MINE {
                count += 1
            }
        }
    }

    return count, nil
}

