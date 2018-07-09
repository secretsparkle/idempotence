// Main function of chess engine

package main

import (
	"fmt"
)

func main() {
	board := buildChessBoard()
	fmt.Println("map: ", board)
}

func buildChessBoard() map[int]string {
	board := make(map[int]string)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch i {
			case 0:
				switch j {
				case 0:
				case 7:
					board[(i*10)+j] = "wR"
				case 1:
				case 6:
					board[(i*10)+j] = "wKn"
				case 2:
				case 5:
					board[(i*10)+j] = "wB"
				case 3:
					board[(i*10)+j] = "wQ"
				case 4:
					board[(i*10)+j] = "wK"
				}
			case 1:
				board[(i*10)+j] = "wP"
			case 6:
				board[(i*10)+j] = "bP"
			case 7:
				switch j {
				case 0:
				case 7:
					board[(i*10)+j] = "bR"
				case 1:
				case 6:
					board[(i*10)+j] = "bKn"
				case 2:
				case 5:
					board[(i*10)+j] = "bB"
				case 3:
					board[(i*10)+j] = "bQ"
				case 4:
					board[(i*10)+j] = "bK"
				}
			default:
				board[(i*10)+j] = "_" // or this could be nil too
			}

		}
	}
	return board
}
