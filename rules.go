package main

import (
	"fmt"
	"math/rand"
)

var BOARD = [10][10]int{}

func setUpBoard() {
	for i := 0; i < len(BOARD); i++ {
		for j := 0; j < len(BOARD[0]); j++ {
			BOARD[i][j] = rand.Intn(2)
		}
	}
}

func getNeighborCount(posX, posY int) int {

	var neighbors = [4]int{}
	var rows, cols int = 10, 10
	var sum int

	if posX == 0 {
		neighbors[1] = BOARD[posY][posX+1]
		neighbors[3] = BOARD[posY][len(BOARD)-1]
	} else if posX == cols-1 {
		neighbors[1] = BOARD[0][posX]
		neighbors[3] = BOARD[posY][posX-1]
	} else {
		neighbors[3] = BOARD[posY][posX-1]
		neighbors[1] = BOARD[posY][posX+1]
	}

	if posY == 0 {
		neighbors[0] = BOARD[len(BOARD[0])-1][posX]
		neighbors[2] = BOARD[posY+1][posX]
	} else if posY == rows-1 {
		neighbors[0] = BOARD[posY-1][posX]
		neighbors[2] = BOARD[posY][0]
	} else {
		neighbors[0] = BOARD[posY-1][posX]
		neighbors[2] = BOARD[posY+1][posX]
	}

	for i := 0; i < 3; i++ {
		if neighbors[i] == 1 {
			sum += 1
		}
	}

	return sum
}

func updateBoard() {
	for posY := 0; posY < len(BOARD); posY++ {
		for posX := 0; posX < len(BOARD[0]); posX++ {
			state := BOARD[posX][posY]
			neighbor_count := getNeighborCount(posX, posY)

			if state == 0 && neighbor_count == 3 {
				BOARD[posX][posY] = 1
			} else if state == 1 && (neighbor_count < 2 || neighbor_count > 3) {
				BOARD[posX][posY] = 0
			} else {
				BOARD[posX][posY] = state
			}
		}
	}
}

func checkIfAllDead() bool {
	num_of_alive := 0
	for i := 0; i < len(BOARD); i++ {
		for j := 0; j < len(BOARD[0]); j++ {
			if BOARD[i][j] == 1 {
				num_of_alive += 1
			}
		}
	}
	return num_of_alive > 0
}

func printBoard(iter int) {
	fmt.Printf(" ------Iter:%v------\n", iter)
	for i := 0; i < len(BOARD); i++ {
		fmt.Println(BOARD[i])
	}
}
