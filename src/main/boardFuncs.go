package main

import (
	"errors"
	"fmt"
)

func containsCoord(coordsList []*coords, coord *coords) (bool) {
    for i := range coordsList {
        c := coordsList[i]
        if c.row == coord.row && c.col == coord.col {
            return true
        }
    }
    return false
}

func revealSquare(board [][]int, square *coords) ([][]int, error) {
    // DFS
    // if mine - game over
    // if not mine 
        // if adjacent to a mine, reveal square and return
        // if not adjacent to a min, reveal square and all surrounding squares. 
    if board[square.row][square.col] == HIDDEN_MINE {
        return board, &GameOver{}
    }

    var toCheck []*coords = make([]*coords, 0) 
    toCheck = append(toCheck, square)
    // loop := 0
    for len(toCheck) > 0 {
        fmt.Println("To Check:")
        for i := range toCheck {
            fmt.Printf("%v, %v\n", toCheck[i].row, toCheck[i].col)
        }
        curr := toCheck[0]
        toCheck = toCheck[1:] 
        numAdjacent, err := adjacentMines(board, curr)
        fmt.Printf("%v\n", numAdjacent)
        if err != nil{
            return board, err
        }
        board[curr.col][curr.row] = VISIBLE_SAFE
        if numAdjacent == 0 {
            newSquares, err := getAdjacent(board, curr)

            for i := range newSquares {
                fmt.Printf("%v, %v\n", newSquares[i].row, newSquares[i].col)
            }
            if err != nil {
                return board, err
            }
            for i := range newSquares {
                newSquare := newSquares[i]
                if containsCoord(toCheck,newSquare)|| board[newSquare.row][newSquare.col] == VISIBLE_SAFE{
                    continue
                }
                toCheck = append(toCheck, newSquares[i])
            }
        }
        // if loop > 2 {
        //     break
        // }
        // loop += 1
    }

    return board, nil 
}

func getAdjacent(board [][]int, square *coords) ([]*coords, error) {

    var adjacent []*coords = []*coords{}

    if square.row < 0 || square.row >= len(board) {
        return adjacent, errors.New("Row value out of bounds") 
    } else if square.col < 0 || square.col >= len(board[0]) {
        return adjacent, errors.New("Col value out of bounds")
    }

    for i := max(0, square.row - 1); i < min(len(board), square.row + 2); i++ {
        for j := max(0, square.col - 1); j < min(len(board[0]), square.col + 2); j++ {
            if i == square.row && j == square.col {
                continue
            }
            adjacent = append(adjacent, &coords{i, j})
        }
    }

    return adjacent, nil
}

