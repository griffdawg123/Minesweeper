package main

import (
	"bufio"
	"io"
	"os"
)

func captureOutput(f func() error) (string, error) {
	orig := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	err := f()
	os.Stdout = orig
	w.Close()
	out, _ := io.ReadAll(r)
	return string(out), err
}

func mockStdinTesting(stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return string(text), err
}

type CoordsSort []*coords 

func (u CoordsSort) Len() int {
    return len(u)
}

func (u CoordsSort) Swap(i, j int) {
    u[i], u[j] = u[j], u[i]
}

func (u CoordsSort) Less(i, j int) bool {
    if u[i].row == u[j].row {
        return u[i].col < u[j].col
    }
    return u[i].row < u[j].row
}


