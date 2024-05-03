package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

const MAXBOARDSIZE = 20

func initFlags(size int, numMines int) (int, int) {
	flag.IntVar(&size, "size", 10, "Create an nxn grid with n=size")
	flag.IntVar(&numMines, "mines", 10, "Number of mines in the game")
	flag.Parse()
	return size, numMines
}

func initGameBoard(size int) ([][]int, error) {
	if size <= 0 {
		return nil, errors.New("gameboard size cannot be negative")
	} else if size > MAXBOARDSIZE {
		return nil, errors.New("maximum game size is 20 cells")
	}
	var gameboard [][]int = make([][]int, size)
	for i := range gameboard {
		gameboard[i] = make([]int, size)
		for j := range gameboard[i] {
			gameboard[i][j] = 1
		}
	}
	return gameboard, nil
}

func debug_print_gameboard(board [][]int) error {
	for i := range board {
		var row []int = board[i]
		for j := range row {
			fmt.Printf("%v ", row[j])
		}
		fmt.Println()
	}
	return nil
}

type coords struct {
	row, col int
}

func getMineCoords(numMines int, stdin io.Reader) ([]*coords, error) {
	var mineCoords []*coords = make([]*coords, numMines)
	var reader = bufio.NewReader(stdin)
	for i := range mineCoords {
		input, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		re := regexp.MustCompile(`(?P<row>[\d]+) (?P<col>[\d]+)`)
		match := re.Match([]byte(input))
		if match {
			matches := re.FindStringSubmatch(input)
			row, err := strconv.Atoi(matches[1])
			if err != nil {
				panic(err)
			}
			col, err := strconv.Atoi(matches[2])
			if err != nil {
				panic(err)
			}
			mineCoords[i] = &coords{row, col}
		}

	}
	return mineCoords, nil
}
