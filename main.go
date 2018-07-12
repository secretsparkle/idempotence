// Main function of chess engine

package main

import (
	"fmt"
)

func main() {
	board := buildChessBoard()
	printBoard(board)
}

// prints a more readable board
func printBoard(board map[int]string) {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			key := (i * 10) + j
			fmt.Printf(board[key])
			if board[key] == "wKn" || board[key] == "bKn" {
				fmt.Printf("  ")
			} else if board[key] == "_" {
				fmt.Printf("    ")
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
}

func movePawn(board map[int]string, key int) {

}

func moveRook(board map[int]string, key int) {

}

func moveKnight(board map[int]string, key int) {

}

func moveBishop(board map[int]string, key int) {

}

func moveQueen(board map[int]string, key int) {

}

func moveKing(board map[int]string, key int) {

}

// secondary move generation driver specific to white
func genWhite(board map[int]string) {
	pieceCount := 0
	column := 0
	row := 0
	for pieceCount < 16 {
		key := (row * 10) + column
		switch board[key] {
		case "wP":
			movePawn(board, key)
		case "wR":
			moveRook(board, key)
		case "wKn":
			moveKnight(board, key)
		case "wB":
			moveBishop(board, key)
		case "wQ":
			moveQueen(board, key)
		case "wK":
			moveKing(board, key)
		}
		pieceCount = pieceCount + 1
	}
}

// secondary move generation driver specific to black
func genBlack(board map[int]string) {
	pieceCount := 0
	column := 0
	row := 0
	for pieceCount < 16 {
		key := (row * 10) + column
		switch board[key] {
		case "bP":
			movePawn(board, key)
		case "bR":
			moveRook(board, key)
		case "bKn":
			moveKnight(board, key)
		case "bB":
			moveBishop(board, key)
		case "bQ":
			moveQueen(board, key)
		case "bK":
			moveKing(board, key)
		}
		pieceCount = pieceCount + 1
	}
}

// driver to produce all available moves from a given board state
func generateMoves(board map[int]string, player string) {
	if player == "w" {
		genWhite(board)
	} else if player == "b" {
		genBlack(board)
	}
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
