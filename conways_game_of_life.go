package main

import (
	"time"
)

func gameTick() {
	setUpBoard()
	printBoard(0)
	i := 1
	for {
		printBoard(i)
		updateBoard()
		all_dead := checkIfAllDead()
		if !all_dead {
			break
		}
		time.Sleep(10 * time.Millisecond)
		i += 1
	}
	printBoard(i)
}

func main() {
	gameTick()
}
