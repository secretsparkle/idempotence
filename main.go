// Main function of chess engine
package main

import (
	"./structures"
	"./tree"
	"fmt"
)

func main() {
	game := new(structures.Tree)
	game.Board = buildChessBoard()
	tree.GenNLevels(game, "w", 4)
	tree.MiniMax(game, 4, "w", "b")
	move := genMove(game)
	printBoard(move)
	//printAllBoards(game)
}

// leftmost best move
func genMove(board *structures.Tree) [8][8]string {
	var nextMove [8][8]string
	max := tree.GetMaxLevel(board.Children)
	for _, state := range board.Children {
		if state.Score == max {
			nextMove = state.Board
			break
		}
	}
	return nextMove
}

// prints all boards generated
func printAllBoards(tree *structures.Tree) {
	printBoard(tree.Board)
	boardStates := tree.Children
	for true {
		if boardStates == nil {
			return
		}
		var children []*structures.Tree
		for _, state := range boardStates {
			fmt.Println(state.Score)
			printBoard(state.Board)
			fmt.Println()
			for _, subState := range state.Children {
				children = append(children, subState)
			}
		}
		boardStates = nil
		for _, state := range children {
			boardStates = append(boardStates, state)
		}
	}
}

// delete after debugging use
// prints a more readable board
func printBoard(board [8][8]string) {
	for _, row := range board {
		for _, square := range row {
			fmt.Printf(square)
			if square == "_" {
				fmt.Printf("    ")
			} else {
				fmt.Printf("   ")
			}
		}
		fmt.Println()
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
					board[i][j] = "wN"
				case 2:
					board[i][j] = "wB"
				case 3:
					board[i][j] = "wQ"
				case 4:
					board[i][j] = "wK"
				case 5:
					board[i][j] = "wB"
				case 6:
					board[i][j] = "wN"
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
					board[i][j] = "bN"
				case 2:
					board[i][j] = "bB"
				case 3:
					board[i][j] = "bQ"
				case 4:
					board[i][j] = "bK"
				case 5:
					board[i][j] = "bB"
				case 6:
					board[i][j] = "bN"
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

func winState(board [8][8]string) string {
	whiteKing := false
	blackKing := false
	for _, row := range board {
		for _, square := range row {
			if square == "wK" {
				whiteKing = true
			} else if square == "bK" {
				blackKing = true
			}
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
