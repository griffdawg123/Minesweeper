package main

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
    for len(toCheck) > 0 {
        curr := toCheck[0]
        toCheck = toCheck[1:] 
        numAdjacent, err := adjacentMines(board, curr)
        if err != nil{
            return board, err
        }
        if numAdjacent > 0 {
            board[curr.col][curr.row] = VISIBLE_SAFE
        } else if numAdjacent == 0 {
            newSquares, err := getAdjacent(board, curr)
            if err != nil {
                return board, err
            }
            toCheck = append(toCheck, newSquares...)
        }
    }

    return board, nil 
}

func getAdjacent(board [][]int, square *coords) ([]*coords, error) {
    return []*coords{}, nil
}

