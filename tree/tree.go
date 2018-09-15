package tree

import (
	"../moves"
	"../structures"
	"fmt" // for debugging purposes
)

const uninitializedScore int = -200

// driver to produce all available moves, one level down, from a given board state
func genMovesLevel(tree *structures.Tree, player string) {
	if player == "w" {
		generatedBoards := moves.GenMoves(tree.Board, player, "b", false, false)
		for _, child := range generatedBoards.Children {
			child.Parent = tree
			child.Score = uninitializedScore
		}
		tree.Children = generatedBoards.Children
	} else if player == "b" {
		generatedBoards := moves.GenMoves(tree.Board, player, "w", false, false)
		for _, child := range generatedBoards.Children {
			child.Parent = tree
			child.Score = uninitializedScore
		}
		tree.Children = generatedBoards.Children
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

func MiniMax(tree *structures.Tree, levels int, player string, enemy string) {
	populateLowestLevelScores(tree, levels, player, enemy)
	for level := levels - 1; level >= 0; level-- {
		populateLevelScores(tree, level)
	}
}

// had to make the two following functions separate, because apparently
// you can't pass operators as params in go? UGH
// and also writing wrapper functions for gt and lt is way unsexy
func getMinLevel(children []*structures.Tree) int {
	// if inCheck eliminates all moves from one branch of tree
	if len(children) < 1 {
		return 1000
	}
	min := children[0].Score
	for _, state := range children {
		if min > state.Score {
			min = state.Score
		}
	}
	return min
}

func GetMaxLevel(children []*structures.Tree) int {
	// if inCheck eliminates all moves from one branch of tree
	if len(children) < 1 {
		return -1000
	}
	max := children[0].Score
	for _, state := range children {
		if max < state.Score {
			max = state.Score
		}
	}
	return max
}

func populateLevelScores(tree *structures.Tree, level int) {
	var children []*structures.Tree
	boardStates := tree.Children
	// first check if the final level before the top has been populated
	if tree.Children[0].Score != -200 && level == 0 {
		tree.Score = GetMaxLevel(tree.Children)
	}
	for i := 0; i < level; i++ {
		for _, state := range boardStates {
			for _, subState := range state.Children {
				children = append(children, subState)
			}
		}
		boardStates = nil
		for _, state := range children {
			boardStates = append(boardStates, state)
		}
		children = nil
	}
	// just in case end is reached on accident
	for _, state := range boardStates {
		if state.Score == -200 && level%2 != 0 {
			state.Score = getMinLevel(state.Children)
		} else if state.Score == -200 && level%2 == 0 {
			state.Score = GetMaxLevel(state.Children)
		}
	}
}

func populateLowestLevelScores(tree *structures.Tree, levels int, player string, enemy string) {
	boardStates := tree.Children
	if levels%2 == 0 {
		player = enemy
	}
	for true {
		var children []*structures.Tree
		if len(boardStates) < 1 {
			return
		}
		if boardStates[0].Children == nil {
			break
		}
		for _, state := range boardStates {
			for _, subState := range state.Children {
				children = append(children, subState)
			}
		}
		boardStates = nil
		for _, state := range children {
			boardStates = append(boardStates, state)
		}
	}
	for _, state := range boardStates {
		state.Score = genScore(state.Board, player)
	}
}

// currently a rudimentary way of score the board
func genScore(board [8][8]string, player string) int {
	score := 0
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if string(board[row][col][0]) != "_" {
				switch {
				case string(board[row][col][0]) == player && string(board[row][col][1]) == "P":
					score += 1
				case string(board[row][col][0]) != player && string(board[row][col][1]) == "P":
					score -= 1
				case string(board[row][col][0]) == player && string(board[row][col][1]) == "N":
					score += 3
				case string(board[row][col][0]) != player && string(board[row][col][1]) == "N":
					score -= 3
				case string(board[row][col][0]) == player && string(board[row][col][1]) == "B":
					score += 3
				case string(board[row][col][0]) != player && string(board[row][col][1]) == "B":
					score -= 3
				case string(board[row][col][0]) == player && string(board[row][col][1]) == "R":
					score += 5
				case string(board[row][col][0]) != player && string(board[row][col][1]) == "R":
					score -= 5
				case string(board[row][col][0]) == player && string(board[row][col][1]) == "Q":
					score += 9
				case string(board[row][col][0]) != player && string(board[row][col][1]) == "Q":
					score -= 9
				case string(board[row][col][0]) == player && string(board[row][col][1]) == "K":
					score += 100
				case string(board[row][col][0]) != player && string(board[row][col][1]) == "K":
					score -= 100
				}
			}
		}
	}
	return score
}
