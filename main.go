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
		return board
	} else if strings.Contains(board[forward], enemy) && board[forward+10] == "_" {
		board[forward] = "_"
		board[forward+10] = board[key]
		board[key] = "_"
		return board
	}
	return nil
}

func moveRook(board map[int]string, key int, player string) map[int]string {
	fmt.Println("Moved rook!")
	return nil
}

func moveKnight(board map[int]string, key int, player string) map[int]string {
	fmt.Println("Moved knight!")
	return nil
}

func moveBishop(board map[int]string, key int, player string) map[int]string {
	fmt.Println("Moved bishop!")
	return nil
}

func moveQueen(board map[int]string, key int, player string) map[int]string {
	fmt.Println("Moved queen!")
	return nil
}

func moveKing(board map[int]string, key int, player string) map[int]string {
	fmt.Println("Moved king!")
	return nil
}

// secondary move generation driver specific to white
func genWhite(board map[int]string) {
	pieceCount := 0
	column := 0
	row := 0
	keepGoing := true
	moves := make([]map[int]string, 0)
	for keepGoing {
		key := (row * 10) + column
		switch board[key] {
		case "wP":
			pawnMove := movePawn(board, key, "w")
			if pawnMove != nil {
				moves = append(moves, pawnMove)
			}
		case "wR":
			rookMove := moveRook(board, key, "w")
			if rookMove != nil {
				moves = append(moves, rookMove)
			}
		case "wKn":
			knightMove := moveKnight(board, key, "w")
			if knightMove != nil {
				moves = append(moves, knightMove)
			}
		case "wB":
			bishopMove := moveBishop(board, key, "w")
			if bishopMove != nil {
				moves = append(moves, bishopMove)
			}
		case "wQ":
			queenMove := moveQueen(board, key, "w")
			if queenMove != nil {
				moves = append(moves, queenMove)
			}
		case "wK":
			kingMove := moveKing(board, key, "w")
			if kingMove != nil {
				moves = append(moves, kingMove)
			}
		}
		pieceCount = pieceCount + 1
		if column == 7 {
			column = 0
			row = row + 1
		} else {
			column = column + 1
		}
		if row >= 16 || pieceCount >= 16 {
			keepGoing = false
		}
	}
}

// secondary move generation driver specific to black
func genBlack(board map[int]string) {
	pieceCount := 0
	column := 0
	row := 0
	keepGoing := true
	for keepGoing {
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
			keepGoing = false
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
