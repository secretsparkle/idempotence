// Main function of chess engine

package main

import (
	"fmt"
)

func main() {
	board := buildChessBoard()
	fmt.Println(winState(board))
}

// build the board!
func buildChessBoard() map[int]string {
	board := make(map[int]string)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch i {
			case 0:
				switch j {
				case 0:
					board[(i*10)+j] = "wR"
				case 1:
					board[(i*10)+j] = "wKn"
				case 2:
					board[(i*10)+j] = "wB"
				case 3:
					board[(i*10)+j] = "wQ"
				case 4:
					board[(i*10)+j] = "wK"
				case 5:
					board[(i*10)+j] = "wB"
				case 6:
					board[(i*10)+j] = "wKn"
				case 7:
					board[(i*10)+j] = "wR"
				}
			case 1:
				board[(i*10)+j] = "wP"
			case 6:
				board[(i*10)+j] = "bP"
			case 7:
				switch j {
				case 0:
					board[(i*10)+j] = "bR"
				case 1:
					board[(i*10)+j] = "bKn"
				case 2:
					board[(i*10)+j] = "bB"
				case 3:
					board[(i*10)+j] = "bQ"
				case 4:
					board[(i*10)+j] = "bK"
				case 5:
					board[(i*10)+j] = "bB"
				case 6:
					board[(i*10)+j] = "bKn"
				case 7:
					board[(i*10)+j] = "bR"
				}
			default:
				board[(i*10)+j] = "_" // or this could be nil too
			}

		}
	}
	return board
}

func winState(board map[int]string) string {
	whiteKing := false
	blackKing := false
	for _, square := range board {
		if square == "wK" {
			whiteKing = true
		} else if square == "bK" {
			blackKing = true
		}
	}
	if !whiteKing {
		return "b" // return black as the winner
	} else if !blackKing {
		return "w"
	} else {
		return "f" // return no one as the winner
	}
}
