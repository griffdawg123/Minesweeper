package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
    "errors"
)

func debug_print_gameboard(board [][]int)  {
	for i := range board {
		var row []int = board[i]
		for j := range row {
			fmt.Printf("%v ", row[j])
		}
		fmt.Println()
	}
}

func getMineCoords(instream io.Reader) (*coords, error) {
    var reader = bufio.NewReader(instream)
    input, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }
    re := regexp.MustCompile(`(?P<row>[\d]+) (?P<col>[\d]+)`)
    match := re.Match([]byte(input))
    if match {
        matches := re.FindStringSubmatch(input)
        row, err1 := strconv.Atoi(matches[1])
        col, err2 := strconv.Atoi(matches[2])
        if err1 != nil || err2 != nil {
            return nil, errors.New("Could not parse mine coords")
        }
        return &coords{row, col}, nil
    } else {
        return nil, errors.New("Could not parse mine coords")
    }

}

func printGameboard(board [][]int) {
    // print column labels
    // print top border
    // before each mine row, print row number and left border
    // ## if hidden, <> if flag, '  ' if visible safe and no mines, else number of adjacent mines
    
    fmt.Print("    ")
    for i := range board[0] {
        if i < 10 {
            fmt.Print("0")
        }
        fmt.Print(strconv.Itoa(i), " ")
    }
    fmt.Print("\n")
    // one for left border, 3 for each number
    fmt.Println("   ", strings.Repeat("-", 3*len(board[0])+1)) 
    for i := range board {
        if i < 10 {
            fmt.Print("0")
        }
        fmt.Print(strconv.Itoa(i), " |")
        for j := range board {
            switch cell := board[i][j]; cell {
                case MINEFLAG:
                    fmt.Print("<> ")
                case WRONGFLAG:
                    fmt.Print("<> ")
                case HIDDEN_MINE:
                    fmt.Print("## ")
                case HIDDEN_SAFE:
                    fmt.Print("## ")
                case VISIBLE_SAFE:
                    adj, err := adjacentMines(board, &coords{i, j})
                    if err != nil {
                        panic(err)
                    }
                    if adj == 0 {
                        fmt.Print("   ")
                    } else {
                        fmt.Print("0", strconv.Itoa(adj), " ")
                    }
                default:
                    fmt.Print("?? ")
            } 
        }
        fmt.Print("\n")
    }
}


