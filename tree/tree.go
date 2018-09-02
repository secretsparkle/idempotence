package tree

import (
	"../moves"
	"../structures"
	//"fmt"
)

// driver to produce all available moves, one level down, from a given board state
func genMovesLevel(tree *structures.Tree, player string) {
	if player == "w" {
		generatedBoards := moves.GenMoves(tree.Board, player, "b")
		tree.Children = generatedBoards.Children
		//fmt.Println("genWhite")
	} else if player == "b" {
		generatedBoards := moves.GenMoves(tree.Board, player, "w")
		tree.Children = generatedBoards.Children
		//fmt.Println("genBlack")
	}
}

// this function currently maxes out at two
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
