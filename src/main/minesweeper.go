package main

const MAXBOARDSIZE = 20

const VISIBLE_SAFE = 0
const HIDDEN_SAFE = 1
const HIDDEN_MINE = 2

type coords struct {
	row, col int
}
type GameOver struct {}

func (m *GameOver) Error() string {
    return "Game Over!"
}

