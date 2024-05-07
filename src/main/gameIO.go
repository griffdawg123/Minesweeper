package main

import "io"
import "strconv"
import "bufio"
import "regexp"
import "fmt"

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

func printGameboard(board [][]int) error {
 
    return nil
}
