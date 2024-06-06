package main

import (
	"time"
)

func gameTick(times int) {
	setUpBoard()
	for i := 0; i < times; i++ {
		printBoard(i + 1)
		time.Sleep(10 * time.Millisecond)
		updateBoard()
	}
}

func main() {
	gameTick(10)
}
