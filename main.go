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
func printBoard(board [8][8]string) {
	for _, row := range board {
		for _, square := range row {
			fmt.Printf(square)
			if square == "wKn" || square == "bKn" {
				fmt.Printf("  ")
			} else if square == "_" {
				fmt.Printf("    ")
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
	}
	/*
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
*/
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
	// vert means goes up 2 and <direction> 1
	vertRight := key + 21
	vertLeft := key + 19
	// horz means goes <direction> 2 and up 1
	horzRight := key + 12
	horzLeft := key + 8
	if board[vertRight] == "_" {
		board[vertRight] = board[key]
		board[key] = "_"
		return board
	} else if board[vertLeft] == "_" {
		board[vertLeft] = board[key]
		board[key] = "_"
		return board
	} else if board[horzRight] == "_" {
		board[horzRight] = board[key]
		board[key] = "_"
		return board
	} else if board[horzLeft] == "_" {
		board[horzLeft] = board[key]
		board[key] = "_"
		return board
	}
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
	moves := make([]map[int]string, 0)
	for keepGoing {
		key := (row * 10) + column
		switch board[key] {
		case "bP":
			pawnMove := movePawn(board, key, "b")
			if pawnMove != nil {
				moves = append(moves, pawnMove)
			}
		case "bR":
			rookMove := moveRook(board, key, "b")
			if rookMove != nil {
				moves = append(moves, rookMove)
			}
		case "bKn":
			knightMove := moveKnight(board, key, "b")
			if knightMove != nil {
				moves = append(moves, knightMove)
			}
		case "bB":
			bishopMove := moveBishop(board, key, "b")
			if bishopMove != nil {
				moves = append(moves, bishopMove)
			}
		case "bQ":
			queenMove := moveQueen(board, key, "b")
			if queenMove != nil {
				moves = append(moves, queenMove)
			}
		case "bK":
			kingMove := moveKing(board, key, "b")
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

func copyBoard(board map[int]string) map[int]string {
	boardCopy := make(map[int]string)
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			key := (i * 10) + j
			boardCopy[key] = board[key]
		}
	}
	return boardCopy
}

// driver to produce all available moves from a given board state
func generateMoves(board map[int]string, player string) {
	boardCopy := copyBoard(board)
	if player == "w" {
		genWhite(boardCopy)
	} else if player == "b" {
		genBlack(boardCopy)
	}
=======
>>>>>>> b1cd9c191ebf8609f4a6b5f630f0e9a036076f7b
}

// build the board!
func buildChessBoard() [8][8]string {
	var board [8][8]string
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			switch i {
			case 0:
				switch j {
				case 0:
					board[i][j] = "wR"
				case 1:
					board[i][j] = "wKn"
				case 2:
					board[i][j] = "wB"
				case 3:
					board[i][j] = "wQ"
				case 4:
					board[i][j] = "wK"
				case 5:
					board[i][j] = "wB"
				case 6:
					board[i][j] = "wKn"
				case 7:
					board[i][j] = "wR"
				}
			case 1:
				board[i][j] = "wP"
			case 6:
				board[i][j] = "bP"
			case 7:
				switch j {
				case 0:
					board[i][j] = "bR"
				case 1:
					board[i][j] = "bKn"
				case 2:
					board[i][j] = "bB"
				case 3:
					board[i][j] = "bQ"
				case 4:
					board[i][j] = "bK"
				case 5:
					board[i][j] = "bB"
				case 6:
					board[i][j] = "bKn"
				case 7:
					board[i][j] = "bR"
				}
			default:
				board[i][j] = "_"
			}

		}
	}
	return board
}
