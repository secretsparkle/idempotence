// Main function of chess engine
package main

import (
	"fmt"
	"strings"
)

type Tree struct {
	Board    [8][8]string
	Children []*Tree // we will use append and make with this slice
	// to check if it is a leaf, we can check if Children is nil
}

func main() {
	game := new(Tree)
	game.Board = buildChessBoard()
	generateMoves(game, "w")
	printBoard(game.Board)
	for _, possibilities := range game.Children {
		printBoard(possibilities.Board)
	}
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
}

// need to account for pawn moving two spaces
// also need to account if pawn would move off the board
func movePawn(board [8][8]string, row int, col int, player string, moveType string) [8][8]string {
	newRow := -1
	newCol := -1
	enemy := ""
	// create default empty board
	var emptyBoard [8][8]string
	emptyBoard[0][0] = "E"
	if player == "w" {
		enemy = "b"
		// cols aren't changing on forward moves, but
		// I kept in the corresponding operations to be
		// verbose and transparent
		switch moveType {
		case "forward":
			newRow = row + 1
			newCol = col
		case "forwardTwo":
			newRow = row + 2
			newCol = col
		case "leftAttack":
			newRow = row + 1
			newCol = col - 1
		case "rightAttack":
			newRow = row + 1
			newCol = col + 1
		}
	} else if player == "b" {
		enemy = "w"
		switch moveType {
		case "forward":
			newRow = row - 1
			newCol = col
		case "forwardTwo":
			newRow = row - 2
			newCol = col
		case "leftAttack":
			newRow = row - 1
			newCol = col - 1
		case "rightAttack":
			newRow = row - 1
			newCol = col + 1
		}
	}
	if newRow == -1 || !withinBoundaries(newRow, newCol) {
		return emptyBoard
	} else if strings.Contains(moveType, "forward") && board[newRow][newCol] == "_" {
		board[newRow][newCol] = board[row][col]
		board[row][col] = "_"
		return board
	} else if strings.Contains(moveType, "Attack") && string(board[newRow][newCol][0]) == enemy {
		board[newRow][newCol] = board[row][col]
		board[row][col] = "_"
		return board
	}
	return emptyBoard
}

func moveRook(board map[int]string, row int, col int, player string) map[int]string {
	fmt.Println("Moved rook!")
	return nil
}

func moveKnightHelper(board [8][8]string, row int, col int, newRow int, newCol int, enemy string) [8][8]string {
	var emptyBoard [8][8]string
	emptyBoard[0][0] = "E"
	if withinBoundaries(newRow, newCol) == true && (board[newRow][newCol] == "_" || string(board[newRow][newCol][0]) == enemy) {
		board[newRow][newCol] = board[row][col]
		board[row][col] = "_"
		return board
	} else {
		return emptyBoard
	}
}

// need to specifiy which move we want the knight to make as a parameter?
func moveKnight(board [8][8]string, row int, col int, main string, direction string, player string, enemy string) [8][8]string {
	// main = vert horz
	// direction = right left up down
	vertMajorUp := row + 2
	vertMajorDown := row - 2
	vertMinorUp := row + 1
	vertMinorDown := row - 1

	horzMajorRight := col + 2
	horzMajorLeft := col - 2
	horzMinorRight := col + 1
	horzMinorLeft := col - 1
	// create empty board in case checks fail
	var emptyBoard [8][8]string
	emptyBoard[0][0] = "E"
	switch {
	case main == "vertMajorUp" && direction == "right":
		return moveKnightHelper(board, row, col, vertMajorUp, horzMinorRight, enemy) // 1
	case main == "vertMajorUp" && direction == "left":
		return moveKnightHelper(board, row, col, vertMajorUp, horzMinorLeft, enemy) // 2
	case main == "vertMajorDown" && direction == "right":
		return moveKnightHelper(board, row, col, vertMajorDown, horzMinorRight, enemy) // 3
	case main == "vertMajorDown" && direction == "left":
		return moveKnightHelper(board, row, col, vertMajorDown, horzMinorLeft, enemy) // 4
	case main == "vertMinorUp" && direction == "right":
		return moveKnightHelper(board, row, col, vertMinorUp, horzMajorRight, enemy) // 5
	case main == "vertMinorUp" && direction == "left":
		return moveKnightHelper(board, row, col, vertMinorUp, horzMajorLeft, enemy) // 6
	case main == "vertMinorDown" && direction == "right":
		return moveKnightHelper(board, row, col, vertMinorDown, horzMajorRight, enemy) // 7
	case main == "vertMinorDown" && direction == "left":
		return moveKnightHelper(board, row, col, vertMinorDown, horzMajorLeft, enemy) // 8
	case main == "horzMajorRight" && direction == "up":
		return moveKnightHelper(board, row, col, horzMajorRight, vertMinorUp, enemy) // 9
	case main == "horzMajorRight" && direction == "down":
		return moveKnightHelper(board, row, col, horzMajorRight, vertMinorDown, enemy) // 10
	case main == "horzMajorLeft" && direction == "up":
		return moveKnightHelper(board, row, col, horzMajorLeft, vertMinorUp, enemy) // 11
	case main == "horzMajorLeft" && direction == "down":
		return moveKnightHelper(board, row, col, horzMajorLeft, vertMinorDown, enemy) // 12
	case main == "horzMinorRight" && direction == "up":
		return moveKnightHelper(board, row, col, horzMinorRight, vertMajorUp, enemy) // 13
	case main == "horzMinorRight" && direction == "down":
		return moveKnightHelper(board, row, col, horzMinorRight, vertMajorDown, enemy) // 14
	case main == "horzMinorLeft" && direction == "up":
		return moveKnightHelper(board, row, col, horzMinorLeft, vertMajorUp, enemy) // 15
	case main == "horzMinorLeft" && direction == "down":
		return moveKnightHelper(board, row, col, horzMinorLeft, vertMajorDown, enemy) // 16
	default:
		return emptyBoard
	}
}

func moveBishop(board map[int]string, row int, col int, player string) map[int]string {
	fmt.Println("Moved bishop!")
	return nil
}

func moveQueen(board map[int]string, row int, col int, player string) map[int]string {
	fmt.Println("Moved queen!")
	return nil
}

func moveKing(board map[int]string, row int, col int, player string) map[int]string {
	fmt.Println("Moved king!")
	return nil
}

// secondary move generation driver specific to white
func genWhite(board [8][8]string) *Tree {
	moves := new(Tree)
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			switch board[row][col] {
			case "wP":
				var pawnMoves [][8][8]string
				pawnMoveForward := movePawn(board, row, col, "w", "forward")
				pawnMoves = append(pawnMoves, pawnMoveForward)
				pawnMoveForwardTwo := movePawn(board, row, col, "w", "forwardTwo")
				pawnMoves = append(pawnMoves, pawnMoveForwardTwo)
				pawnMoveLeftAttack := movePawn(board, row, col, "w", "leftAttack")
				pawnMoves = append(pawnMoves, pawnMoveLeftAttack)
				pawnMoveRightAttack := movePawn(board, row, col, "w", "rightAttack")
				pawnMoves = append(pawnMoves, pawnMoveRightAttack)
				for _, move := range pawnMoves {
					if move[0][0] != "E" {
						newBranch := new(Tree)
						newBranch.Board = move
						moves.Children = append(moves.Children, newBranch)
					}
				}
				/*
					case "wR":
						rookMove := moveRook(board, row, col, "w")
						if rookMove != nil {
							moves.Children = append(moves.Children, rookMove)
						}
				*/
			case "wKn":
				var knightMoves [][8][8]string
				knightMoveVMaUR := moveKnight(board, row, col, "vertMajorUp", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveVMaUR)
				knightMoveVMaUL := moveKnight(board, row, col, "vertMajorUp", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveVMaUL)
				knightMoveVMaDR := moveKnight(board, row, col, "vertMajorDown", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveVMaDR)
				knightMoveVMaDL := moveKnight(board, row, col, "vertMajorDown", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveVMaDL)
				knightMoveMiUR := moveKnight(board, row, col, "vertMinorUp", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveMiUR)
				knightMoveVMiUL := moveKnight(board, row, col, "vertMinorUp", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveVMiUL)
				knightMoveVMiDR := moveKnight(board, row, col, "vertMinorDown", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveVMiDR)
				knightMoveVMiDL := moveKnight(board, row, col, "vertMinorDown", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveVMiDL)
				knightMoveHMaUR := moveKnight(board, row, col, "horzMajorUp", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMaUR)
				knightMoveHMaUL := moveKnight(board, row, col, "horzMajorUp", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMaUL)
				knightMoveHMaDR := moveKnight(board, row, col, "horzMajorDown", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMaDR)
				knightMoveHMaDL := moveKnight(board, row, col, "horzMajorDown", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMaDL)
				knightMoveHMiUR := moveKnight(board, row, col, "horzMinorUp", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMiUR)
				knightMoveHMiUL := moveKnight(board, row, col, "horzMinorUp", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMiUL)
				knightMoveHMiDR := moveKnight(board, row, col, "horzMinorDown", "right", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMiDR)
				knightMoveHMiDL := moveKnight(board, row, col, "horzMinorDown", "left", "w", "b")
				knightMoves = append(knightMoves, knightMoveHMiDL)
				for _, move := range knightMoves {
					if move[0][0] != "E" {
						newBranch := new(Tree)
						newBranch.Board = move
						moves.Children = append(moves.Children, newBranch)
					}
				}
				/*
					case "wB":
						bishopMove := moveBishop(board, row, col, "w")
						if bishopMove != nil {
							moves.Children = append(moves.Children, bishopMove)
						}
					case "wQ":
						queenMove := moveQueen(board, row, col, "w")
						if queenMove != nil {
							moves.Children = append(moves.Children, queenMove)
						}
					case "wK":
						kingMove := moveKing(board, row, col, "w")
						if kingMove != nil {
							moves.Children = append(moves.Children, kingMove)
						}
				*/

			}
		}
	}
	return moves
}

// secondary move generation driver specific to black
func genBlack(board [8][8]string) *Tree {
	moves := new(Tree)
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			switch board[row][col] {
			case "bP":
				pawnMove := movePawn(board, row, col, "b", "forward")
				if pawnMove[0][0] != "E" {
					newBranch := new(Tree)
					newBranch.Board = pawnMove
					moves.Children = append(moves.Children, newBranch)
				}
				/*
					case "bR":
						rookMove := moveRook(board, row, col, "b")
						if rookMove != nil {
							moves.Children = append(moves.Children, rookMove)
						}
				*/
				/*
					case "bKn":
						var knightMoves [][8][8]string
						knightMoveVUR := moveKnight(board, row, col, "vert", "up", "right", "b", "w")
						knightMoves = append(knightMoves, knightMoveVUR)
						knightMoveVUL := moveKnight(board, row, col, "vert", "up", "left", "b", "w")
						knightMoves = append(knightMoves, knightMoveVUL)
						knightMoveVDR := moveKnight(board, row, col, "vert", "down", "right", "b", "w")
						knightMoves = append(knightMoves, knightMoveVDR)
						knightMoveVDL := moveKnight(board, row, col, "vert", "down", "left", "b", "w")
						knightMoves = append(knightMoves, knightMoveVDL)
						knightMoveHUR := moveKnight(board, row, col, "horz", "up", "right", "b", "w")
						knightMoves = append(knightMoves, knightMoveHUR)
						knightMoveHUL := moveKnight(board, row, col, "horz", "up", "left", "b", "w")
						knightMoves = append(knightMoves, knightMoveHUL)
						knightMoveHDR := moveKnight(board, row, col, "horz", "down", "right", "b", "w")
						knightMoves = append(knightMoves, knightMoveHDR)
						knightMoveHDL := moveKnight(board, row, col, "horz", "down", "left", "b", "w")
						knightMoves = append(knightMoves, knightMoveHDL)
						for _, move := range knightMoves {
							if move[0][0] != "E" {
								newBranch := new(Tree)
								newBranch.Board = move
								moves.Children = append(moves.Children, newBranch)
							}
						}
				*/
				/*
					case "bB":
						bishopMove := moveBishop(board, row, col, "b")
						if bishopMove != nil {
							moves.Children = append(moves.Children, bishopMove)
						}
					case "bQ":
						queenMove := moveQueen(board, row, col, "b")
						if queenMove != nil {
							moves.Children = append(moves.Children, queenMove)
						}
					case "bK":
						kingMove := moveKing(board, row, col, "b")
						if kingMove != nil {
							moves.Children = append(moves.Children, kingMove)
						}
				*/

			}
		}
	}
	return moves
}

// driver to produce all available moves from a given board state
func generateMoves(tree *Tree, player string) {
	if player == "w" {
		generatedBoards := genWhite(tree.Board)
		tree.Children = generatedBoards.Children
		fmt.Println("genWhite")
	} else if player == "b" {
		generatedBoards := genBlack(tree.Board)
		tree.Children = generatedBoards.Children
	}
}

func withinBoundaries(moveRow int, moveCol int) bool {
	if (moveRow >= 0 && moveRow < 8) && (moveCol >= 0 && moveCol < 8) {
		return true
	} else {
		return false
	}
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
