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
	var game *Tree
	game.Board = buildChessBoard()
	generateMoves(game, "w")
	printBoard(game.Board)
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
func movePawn(board [8][8]string, row int, col int, player string) [8][8]string {
	forward := 0
	enemy := ""
	if player == "w" {
		forward = row + 1
		enemy = "b"
	} else if player == "b" {
		forward = row - 1
		enemy = "w"
	}
	if board[forward][col] == "_" {
		board[forward][col] = board[row][col]
		board[row][col] = "_"
		return board
	} else if strings.Contains(board[forward][col], enemy) && board[forward+1][col] == "_" { // what circumstance does this code account for?
		board[forward][col] = "_"
		board[forward+1][col] = board[row][col]
		board[row][col] = "_"
		return board
	}
	return nil // what should be returned here? (Can't be nil)
}

func moveRook(board map[int]string, row int, col int, player string) map[int]string {
	fmt.Println("Moved rook!")
	return nil
}

/*
func 8=D~(.Y.)() {
      fmt.Println("Aww yeeee")
}
*/

// need to specifiy which move we want the knight to make as a parameter?
func moveKnight(board [8][8]string, row int, col int, main string, modifier string, direction string, player string, enemy string) [8][8]string {
	// TODO: STILL NEED TO CHECK IF MOVE IS LEGAL (ON BOARD), MAKE SEP FUNCTION?
	// vertMain means go up or down 2
	vertMainUp := row + 2
	vertMainDown := row - 2
	vertUp := row + 1
	vertDown := row - 1
	// horzMain means goes right or left 2
	horzMainRight := col + 2
	horzMainLeft := col - 2
	horzRight := col + 1
	horzLeft := col - 1
	// could probably make a function to handle the gruntwork of each case
	// as the same lines of code are used over and over again
	// also, this could easily be made into an attacking function as well,
	// just by checking if the space to be moved to is an enemy space
	// TODO: give this function attacking functionality
	if main == "vert" && modifier == "up" {
		if direction == "right" && (board[vertMainUp][horzRight] == "_" || string(board[vertMainUp][horzRight][0]) == enemy) && withinBoundaries(vertMainUp, horzRight) == true {
			board[vertMainUp][horzRight] = board[row][col]
			board[row][col] = "_"
			return board
		} else if direction == "left" && (board[vertMainUp][horzLeft] == "_" || string(board[vertMainUp][horzLeft][0]) == enemy) && withinBoundaries(vertMainUp, horzLeft) == true {
			board[vertMainUp][horzLeft] = board[row][col]
			board[row][col] = "_"
			return board
		}
	} else if main == "vert" && modifier == "down" {
		if direction == "right" && (board[vertMainDown][horzRight] == "_" || string(board[vertMainDown][horzRight][0]) == enemy) && withinBoundaries(vertMainDown, horzRight) == true {
			board[vertMainDown][horzRight] = board[row][col]
			board[row][col] = "_"
			return board
		} else if direction == "left" && (board[vertMainDown][horzLeft] == "_" || string(board[vertMainDown][horzLeft][0]) == enemy) && withinBoundaries(vertMainDown, horzLeft) == true {
			board[vertMainDown][horzLeft] = board[row][col]
			board[row][col] = "_"
			return board
		}
	} else if main == "horz" && modifier == "right" {
		if direction == "up" && (board[horzMainRight][vertUp] == "_" || string(board[horzMainRight][vertUp][0]) == enemy) && withinBoundaries(horzMainRight, vertUp) == true {
			board[horzMainRight][vertUp] = board[row][col]
			board[row][col] = "_"
			return board
		} else if direction == "down" && (board[horzMainRight][vertDown] == "_" || string(board[horzMainRight][vertDown][0]) == enemy) && withinBoundaries(horzMainRight, vertDown) == true {
			board[horzMainRight][vertDown] = board[row][col]
			board[row][col] = "_"
			return board
		}
	} else if main == "horz" && modifier == "left" {
		if direction == "up" && (board[horzMainLeft][vertUp] == "_" || string(board[horzMainLeft][vertUp][0]) == enemy) && withinBoundaries(horzMainLeft, vertUp) == true {
			board[horzMainLeft][vertUp] = board[row][col]
			board[row][col] = "_"
			return board
		} else if direction == "down" && (board[horzMainLeft][vertDown] == "_" || string(board[horzMainLeft][vertDown][0]) == enemy) && withinBoundaries(horzMainLeft, vertDown) == true {
			board[horzMainLeft][vertDown] = board[row][col]
			board[row][col] = "_"
			return board
		}
	}
	return nil // same question as movePawn, what should the default return type be?
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
func genWhite(board [8][8]string) {
	pieceCount := 0
	row := 0
	col := 0
	keepGoing := true
	var moves []*Tree
	for keepGoing {
		switch board[row][col] {
		case "wP":
			pawnMove := movePawn(board, row, col, "w")
			if pawnMove != nil {
				var move *Tree
				move.Board = pawnMove
				moves = append(moves, move)
			}
		case "wR":
			rookMove := moveRook(board, row, col, "w")
			if rookMove != nil {
				moves.Children.append(rookMove)
			}
		case "wKn":
			knightMove := moveKnight(board, row, col, "vert", "up", "right", "w", "b")
			if knightMove != nil {
				moves.Children.append(knightMove)
			}
		case "wB":
			bishopMove := moveBishop(board, row, col, "w")
			if bishopMove != nil {
				moves.Children.append(bishopMove)
			}
		case "wQ":
			queenMove := moveQueen(board, row, col, "w")
			if queenMove != nil {
				moves.Children.append(queenMove)
			}
		case "wK":
			kingMove := moveKing(board, row, col, "w")
			if kingMove != nil {
				moves.Children.append(kingMove)
			}
		}
		// we need to be sure to come back and deal with pieceCount
		pieceCount = pieceCount + 1
		if column == 7 {
			column = 0
			row += 1
		} else {
			column += 1
		}
		if row >= 8 || pieceCount >= 16 {
			keepGoing = false
		}
	}
}

// secondary move generation driver specific to black
func genBlack(board [8][8]string) {
	pieceCount := 0
	row := 0
	col := 0
	keepGoing := true
	var moves Tree
	for keepGoing {
		switch board[row][col] {
		case "bP":
			pawnMove := movePawn(board, row, col, "b")
			if pawnMove != nil {
				moves.Children = append(moves.Children, pawnMove)
			}
		case "bR":
			rookMove := moveRook(board, row, col, "b")
			if rookMove != nil {
				moves.Children.append(rookMove)
			}
		case "bKn":
			knightMove := moveKnight(board, row, col, "vert", "up", "right", "b", "w")
			if knightMove != nil {
				moves.Children.append(knightMove)
			}
		case "bB":
			bishopMove := moveBishop(board, row, col, "b")
			if bishopMove != nil {
				moves.Children.append(bishopMove)
			}
		case "bQ":
			queenMove := moveQueen(board, row, col, "b")
			if queenMove != nil {
				moves.Children.append(queenMove)
			}
		case "bK":
			kingMove := moveKing(board, row, col, "b")
			if kingMove != nil {
				moves.Children.append(kingMove)
			}
		}
		// we need to be sure to come back and deal with pieceCount
		pieceCount = pieceCount + 1
		if col == 7 {
			col = 0
			row += 1
		} else {
			col += 1
		}
		if row >= 8 || pieceCount >= 16 {
			keepGoing = false
		}
	}
}

// driver to produce all available moves from a given board state
func generateMoves(tree *Tree, player string) {
	if player == "w" {
		tree.Children = genWhite(tree.Board)
	} else if player == "b" {
		tree.Children.append(genBlack(tree.Board))
	}
	// END OF NEW CODE
}

func withinBoundaries(moveRow int, moveCol int) bool {
	if moveRow >= 0 || moveRow < 8 || moveCol >= 0 || moveCol < 8 {
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
