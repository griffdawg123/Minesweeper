package main

import (
	"testing"
    "sort"
	"github.com/stretchr/testify/assert"
)

func TestAdjacentFalse(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var coordsToCheck coords = coords{4, 4}
    var adjacent int
    adjacent, err =  adjacentMines(newBoard, &coordsToCheck)
    assert.Nil(t, err)
    assert.Equal(t, adjacent, 0)
}

func TestAdjacentTrue(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var coordsToCheck coords = coords{1, 1}
    var adjacent int
    adjacent, err =  adjacentMines(newBoard, &coordsToCheck)
    assert.Nil(t, err)
    assert.Equal(t, adjacent, 1)
}

func TestAdjacentTrueTop(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var coordsToCheck coords = coords{0, 1}
    var adjacent int
    adjacent, err =  adjacentMines(newBoard, &coordsToCheck)
    assert.Nil(t, err)
    assert.Equal(t, adjacent, 1)
}

func TestAdjacentTrueLeft(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var coordsToCheck coords = coords{1, 0}
    var adjacent int
    adjacent, err =  adjacentMines(newBoard, &coordsToCheck)
    assert.Nil(t, err)
    assert.Equal(t, adjacent, 1)
}

func TestAdjacentTrueRight(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{4, 4}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var coordsToCheck coords = coords{3, 4}
    var adjacent int
    adjacent, err =  adjacentMines(newBoard, &coordsToCheck)
    assert.Nil(t, err)
    assert.Equal(t, adjacent, 1)
}
func TestAdjacentTrueBottom(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{4, 4}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var coordsToCheck coords = coords{4, 3}
    var adjacent int
    adjacent, err =  adjacentMines(newBoard, &coordsToCheck)
    assert.Nil(t, err)
    assert.Equal(t, adjacent, 1)
}

func TestAdjacentOutOfBounrds(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var coordsToCheck coords = coords{10, 10}
    _, err =  adjacentMines(newBoard, &coordsToCheck)
    assert.Error(t, err)
}

func TestRevealAdjacent(t *testing.T) {

    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var selectedCoords coords = coords{1, 1}
    var revealed [][]int
    revealed, err = revealSquare(newBoard, &selectedCoords)
    
    assert.Nil(t, err)

    var expectedBoard = [][]int{{2, 1, 1, 1, 1},{1, 0, 1, 1, 1},{1, 1, 1, 1, 1},{1, 1, 1, 1, 1},{1, 1, 1, 1, 1}}
    for row := range revealed {
        for col := range revealed[row] {
            assert.Equal(t, expectedBoard[row][col], revealed[row][col])
        }
    } 
}


func TestRevealNotAdjacent(t *testing.T) {

    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var selectedCoords coords = coords{3, 3}
    var revealed [][]int
    revealed, err = revealSquare(newBoard, &selectedCoords)
    assert.Nil(t, err)
    var expectedBoard = [][]int{{2, 0, 0, 0, 0},{0, 0, 0, 0, 0},{0, 0, 0, 0, 0},{0, 0, 0, 0, 0},{0, 0, 0, 0, 0}}
    for row := range revealed {
        for col := range revealed[row] {
            assert.Equal(t, expectedBoard[row][col], revealed[row][col])
        }
    } 
}

func TestRevealMine(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)

    var mineCoords []*coords = []*coords{{0, 0}}
    newBoard, err := placeMines(mineCoords, board)
    assert.Nil(t, err)
   
    var selectedCoords coords = coords{0, 0}
    _, err = revealSquare(newBoard, &selectedCoords)
    if assert.Error(t, err) {
        assert.Equal(t, &GameOver{}, err)
    }
}

func TestCoordsSort(t *testing.T) {
    var toSort []*coords = []*coords{{1, 1}, {3, 2}, {3, 3}, {2, 2}, {4, 4}}
    sort.Sort(CoordsSort(toSort))
    expected := []*coords{{1, 1}, {2, 2}, {3, 2}, {3, 3}, {4, 4}}
    for i := range toSort {
        assert.Equal(t, toSort[i], expected[i])
    }
}

func TestGetAdjacentMiddle(t *testing.T) {
    var board [][]int
    var err error
    var size int = 5
    board, err = initGameBoard(size, size)
    assert.Nil(t, err)
    
    var selectedCoords coords = coords{3, 3}
    adjacent, err := getAdjacent(board, &selectedCoords)
    sort.Sort(CoordsSort(adjacent))
    assert.Equal(t, []*coords{{2, 2}, {2, 3}, {2, 4}, {3, 2}, {3, 4}, {4, 2}, {4, 3}, {4, 4}}, adjacent)
}

//func TestGetAdjacentEdge(t *testing.T) {
//
//}

//func TestGetAdjacentCorner(t *testing.T) {
//
//}
