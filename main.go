// Main function of chess engine
package main

import (
	"./structures"
	"./tree"
	"fmt"
	"time"
)

func main() {
	game := new(structures.Tree)
	game.Board = buildChessBoard()
	levels := 3
	player := "w"
	enemy := "b"
	//move := nextMove(game, levels, player, enemy)
	//printBoard(move.Board)
	for winState(game) != true {
		move := nextMove(game, levels, player, enemy)
		fmt.Println(move.Score)
		fmt.Println(player)
		printBoard(move.Board)
		fmt.Println()
		game = move
		if player == "w" {
			player = "b"
			enemy = "w"
		} else {
			player = "w"
			enemy = "b"
		}
		time.Sleep(1000 * time.Millisecond)
	}
	//printAllBoards(game)
}

func nextMove(game *structures.Tree, levels int, player string, enemy string) *structures.Tree {
	tree.GenNLevels(game, player, levels)
	tree.MiniMax(game, levels, player, enemy)
	move := genMove(game)
	return move
}

// leftmost best move
func genMove(board *structures.Tree) *structures.Tree {
	var nextMove *structures.Tree
	max := tree.GetMaxLevel(board.Children)
	for _, state := range board.Children {
		if state.Score == max {
			nextMove = state
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

func winState(game *structures.Tree) bool {
	whiteKing := false
	blackKing := false
	for _, row := range game.Board {
		for _, square := range row {
			if square == "wK" {
				whiteKing = true
			} else if square == "bK" {
				blackKing = true
			}
		}
	}
	if !whiteKing {
		return true // return black as the winner
	} else if !blackKing {
		return true
	} else {
		return false // return no one as the winner
	}
}
