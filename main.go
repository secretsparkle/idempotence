// Main function of chess engine

package main

import (
	"fmt"
	"strings"
)

func main() {
	board := buildChessBoard()
	generateMoves(board, "w")
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

func movePawn(board map[int]string, key int, player string) map[int]string {
	forward := 0
	enemy := ""
	if player == "w" {
		forward = key + 10
		enemy = "b"
	} else if player == "b" {
		forward = key - 10
		enemy = "w"
	}
	if board[forward] == "_" {
		board[forward] = board[key]
		board[key] = "_"
	} else if strings.Contains(board[forward], enemy) && board[forward+10] == "_" {
		board[forward] = "_"
		board[forward+10] = board[key]
		board[key] = "_"
	}
	return board
}

func moveRook(board map[int]string, key int, player string) {
	fmt.Println("Moved rook!")
}

func moveKnight(board map[int]string, key int, player string) {
	fmt.Println("Moved knight!")
}

func moveBishop(board map[int]string, key int, player string) {
	fmt.Println("Moved bishop!")
}

func moveQueen(board map[int]string, key int, player string) {
	fmt.Println("Moved queen!")
}

func moveKing(board map[int]string, key int, player string) {
	fmt.Println("Moved king!")
}

// secondary move generation driver specific to white
func genWhite(board map[int]string) {
	pieceCount := 0
	column := 0
	row := 0
	callCount := 0
	exitCondition := true
	for exitCondition {
		key := (row * 10) + column
		switch board[key] {
		case "wP":
			movePawn(board, key, "w")
			callCount = callCount + 1
		case "wR":
			moveRook(board, key, "w")
		case "wKn":
			moveKnight(board, key, "w")
		case "wB":
			moveBishop(board, key, "w")
		case "wQ":
			moveQueen(board, key, "w")
		case "wK":
			moveKing(board, key, "w")
		}
		pieceCount = pieceCount + 1
		if column == 7 {
			column = 0
			row = row + 1
		} else {
			column = column + 1
		}
		if row >= 16 || pieceCount >= 16 {
			exitCondition = false
		}
	}
}

// secondary move generation driver specific to black
func genBlack(board map[int]string) {
	pieceCount := 0
	column := 0
	row := 0
	exitCondition := true
	for exitCondition {
		key := (row * 10) + column
		switch board[key] {
		case "bP":
			movePawn(board, key, "b")
		case "bR":
			moveRook(board, key, "b")
		case "bKn":
			moveKnight(board, key, "b")
		case "bB":
			moveBishop(board, key, "b")
		case "bQ":
			moveQueen(board, key, "b")
		case "bK":
			moveKing(board, key, "b")
		}
		pieceCount = pieceCount + 1
		if column == 7 {
			column = 0
			row = row + 1
		} else {
			column = column + 1
		}
		if row >= 16 || pieceCount >= 16 {
			exitCondition = false
		}
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
