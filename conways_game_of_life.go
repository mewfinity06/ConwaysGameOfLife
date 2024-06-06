package main

import (
	"time"
)

func gameTick() {
	setUpBoard()
	i := 0
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
}

func main() {
	gameTick()
}
