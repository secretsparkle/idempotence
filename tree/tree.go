package tree

import (
	"../moves"
	"../structures"
	"fmt" // for debugging purposes
)

// driver to produce all available moves, one level down, from a given board state
func genMovesLevel(tree *structures.Tree, player string) {
	if player == "w" {
		generatedBoards := moves.GenMoves(tree.Board, player, "b", false)
		tree.Children = generatedBoards.Children
		//fmt.Println("genWhite")
	} else if player == "b" {
		generatedBoards := moves.GenMoves(tree.Board, player, "w", false)
		tree.Children = generatedBoards.Children
		//fmt.Println("genBlack")
	}
}

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

// this function currently maxes out at two levels
// stack overflow!
func GenNLevels(tree *structures.Tree, player string, levels int) {
	genMovesLevel(tree, player)
	boardStates := tree.Children
	for i := 0; i < levels-1; i++ {
		if player == "w" {
			player = "b"
		} else {
			player = "w"
		}
		var children []*structures.Tree
		for _, state := range boardStates {
			//printBoard(state.Board)
			//fmt.Println()
			genMovesLevel(state, player)
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
