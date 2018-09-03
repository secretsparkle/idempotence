package moves

import (
	"../structures"
	"strings"
)

func move(board [8][8]string, row int, col int, newRow int, newCol int, enemy string) [8][8]string {
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

func genRookMoves(board [8][8]string, row int, col int, enemy string) [][8][8]string {
	var rookMoves [][8][8]string
	// vertical down moves
	for vertDown := col + 1; vertDown < 8; vertDown++ {
		rookMove := move(board, row, col, row, vertDown, enemy)
		// rooks can only move as far as the first piece they encounter
		if rookMove[0][0] == "E" {
			break
		} else if string(board[row][vertDown][0]) == enemy {
			// if a rook encounters an enemy, it won't be invalid,
			// but we don't want the rook to move beyond it
			rookMoves = append(rookMoves, rookMove)
			break
		} else {
			rookMoves = append(rookMoves, rookMove)
		}
	}
	// vertical up moves
	for vertUp := col - 1; vertUp >= 0; vertUp-- {
		rookMove := move(board, row, col, row, vertUp, enemy)
		if rookMove[0][0] == "E" {
			break
		} else if string(board[row][vertUp][0]) == enemy {
			rookMoves = append(rookMoves, rookMove)
			break
		} else {
			rookMoves = append(rookMoves, rookMove)
		}
	}
	// horizontal right moves
	for horzRight := row + 1; horzRight < 8; horzRight++ {
		rookMove := move(board, row, col, horzRight, col, enemy)
		if rookMove[0][0] == "E" {
			break
		} else if string(board[horzRight][col][0]) == enemy {
			rookMoves = append(rookMoves, rookMove)
			break
		} else {
			rookMoves = append(rookMoves, rookMove)
		}
	}
	// horizontal left moves
	for horzLeft := row - 1; horzLeft >= 0; horzLeft-- {
		rookMove := move(board, row, col, horzLeft, col, enemy)
		if rookMove[0][0] == "E" {
			break
		} else if string(board[horzLeft][col][0]) == enemy {
			rookMoves = append(rookMoves, rookMove)
			break
		} else {
			rookMoves = append(rookMoves, rookMove)
		}
	}
	return rookMoves
}

func genKnightMoves(board [8][8]string, row int, col int, enemy string) [][8][8]string {
	var knightMoves [][8][8]string
	vertMajorUp := row - 2
	vertMajorDown := row + 2
	vertMinorUp := row - 1
	vertMinorDown := row + 1

	horzMajorRight := col + 2
	horzMajorLeft := col - 2
	horzMinorRight := col + 1
	horzMinorLeft := col - 1

	knightMove := move(board, row, col, vertMajorUp, horzMinorRight, enemy) // 1
	knightMoves = append(knightMoves, knightMove)
	knightMove = move(board, row, col, vertMajorUp, horzMinorLeft, enemy) // 2
	knightMoves = append(knightMoves, knightMove)
	knightMove = move(board, row, col, vertMajorDown, horzMinorRight, enemy) // 3
	knightMoves = append(knightMoves, knightMove)
	knightMove = move(board, row, col, vertMajorDown, horzMinorLeft, enemy) // 4
	knightMoves = append(knightMoves, knightMove)
	knightMove = move(board, row, col, vertMinorUp, horzMajorRight, enemy) // 5
	knightMoves = append(knightMoves, knightMove)
	knightMove = move(board, row, col, vertMinorUp, horzMajorLeft, enemy) // 6
	knightMoves = append(knightMoves, knightMove)
	knightMove = move(board, row, col, vertMinorDown, horzMajorRight, enemy) // 7
	knightMoves = append(knightMoves, knightMove)
	knightMove = move(board, row, col, vertMinorDown, horzMajorLeft, enemy) // 8
	knightMoves = append(knightMoves, knightMove)
	return knightMoves
}

func genBishopMoves(board [8][8]string, row int, col int, enemy string) [][8][8]string {
	var bishopMoves [][8][8]string
	for vertUpRow, horzRightCol := row-1, col+1; vertUpRow >= 0 && horzRightCol < 8; vertUpRow, horzRightCol = vertUpRow-1, horzRightCol+1 {
		bishopMove := move(board, row, col, vertUpRow, horzRightCol, enemy)
		// bishops can only move as far as the first piece they encounter
		if bishopMove[0][0] == "E" {
			break
		} else if string(board[vertUpRow][horzRightCol][0]) == enemy {
			// if a bishop encounters an enemy, it won't be invalid,
			// but we don't want the rook to move beyond it
			bishopMoves = append(bishopMoves, bishopMove)
			break
		} else {
			bishopMoves = append(bishopMoves, bishopMove)
		}
	}
	for vertDownRow, horzRightCol := row+1, col+1; vertDownRow < 8 && horzRightCol < 8; vertDownRow, horzRightCol = vertDownRow+1, horzRightCol+1 {
		bishopMove := move(board, row, col, vertDownRow, horzRightCol, enemy)
		if bishopMove[0][0] == "E" {
			break
		} else if string(board[vertDownRow][horzRightCol][0]) == enemy {
			bishopMoves = append(bishopMoves, bishopMove)
			break
		} else {
			bishopMoves = append(bishopMoves, bishopMove)
		}
	}
	for vertDownRow, horzLeftCol := row+1, col-1; vertDownRow < 8 && horzLeftCol >= 0; vertDownRow, horzLeftCol = vertDownRow+1, horzLeftCol-1 {
		bishopMove := move(board, row, col, vertDownRow, horzLeftCol, enemy)
		if bishopMove[0][0] == "E" {
			break
		} else if string(board[vertDownRow][horzLeftCol][0]) == enemy {
			bishopMoves = append(bishopMoves, bishopMove)
			break
		} else {
			bishopMoves = append(bishopMoves, bishopMove)
		}
	}
	for vertUpRow, horzLeftCol := row-1, col-1; vertUpRow >= 0 && horzLeftCol >= 0; vertUpRow, horzLeftCol = vertUpRow-1, horzLeftCol-1 {
		bishopMove := move(board, row, col, vertUpRow, horzLeftCol, enemy)
		if bishopMove[0][0] == "E" {
			break
		} else if string(board[vertUpRow][horzLeftCol][0]) == enemy {
			bishopMoves = append(bishopMoves, bishopMove)
			break
		} else {
			bishopMoves = append(bishopMoves, bishopMove)
		}
	}
	return bishopMoves
}

func isCheck(board [8][8]string, kingMoveRow int, kingMoveCol int, player string, enemy string, genKing bool) bool {
	moveBoards := GenMoves(board, enemy, player, true)
	for _, moveBoard := range moveBoards.Children {
		if string(moveBoard.Board[kingMoveRow][kingMoveCol][0]) == enemy {
			return true
		}
	}
	return false
}

func genKingMove(board [8][8]string, row int, col int, newRow int, newCol int, player string, enemy string, genKing bool) [8][8]string {
	var emptyBoard [8][8]string
	emptyBoard[0][0] = "E"
	kingMove := move(board, row, col, newRow, newCol, enemy)
	// this logic will short-circuit if genKing is true. i.e., if we are already within a genKing
	// statement, we don't need to check if this new genKing state will put the king in check, because
	// we are only checking if the other king could take the our king on the next move, we are not actually
	// generating a move for the move tree
	if kingMove[0][0] != "E" && (genKing || !isCheck(kingMove, newRow, newCol, player, enemy, genKing)) {
		return kingMove
	} else {
		return emptyBoard
	}
}

func genKingMoves(board [8][8]string, row int, col int, player string, enemy string, genKing bool) [][8][8]string {
	var kingMoves [][8][8]string
	// up
	kingMove := genKingMove(board, row, col, row-1, col, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	// up-diagonal-right
	kingMove = genKingMove(board, row, col, row-1, col+1, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	// right
	kingMove = genKingMove(board, row, col, row, col+1, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	// down-diagonal-right
	kingMove = genKingMove(board, row, col, row+1, col+1, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	// down
	kingMove = genKingMove(board, row, col, row+1, col, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	// down-diagonal-left
	kingMove = genKingMove(board, row, col, row+1, col-1, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	// left
	kingMove = genKingMove(board, row, col, row, col-1, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	// up-diagonal-left
	kingMove = genKingMove(board, row, col, row-1, col-1, player, enemy, genKing)
	if kingMove[0][0] != "E" {
		kingMoves = append(kingMoves, kingMove)
	}
	return kingMoves
}

func genNewBranches(pieceMoves [][8][8]string, moves *structures.Tree) {
	for _, move := range pieceMoves {
		if move[0][0] != "E" {
			newBranch := new(structures.Tree)
			newBranch.Board = move
			moves.Children = append(moves.Children, newBranch)
		}
	}
}

// abstract move generation driver
func GenMoves(board [8][8]string, player string, enemy string, genKing bool) *structures.Tree {
	moves := new(structures.Tree)
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			switch {
			// Pawns
			// TODO Pawns will currently jump over opponents when moving two spaces
			// This is illegal and needs to be fixed!
			case string(board[row][col][0]) == player && string(board[row][col][1]) == "P":
				var pawnMoves [][8][8]string
				pawnMoveForward := movePawn(board, row, col, player, "forward")
				pawnMoves = append(pawnMoves, pawnMoveForward)
				pawnMoveForwardTwo := movePawn(board, row, col, player, "forwardTwo")
				pawnMoves = append(pawnMoves, pawnMoveForwardTwo)
				pawnMoveLeftAttack := movePawn(board, row, col, player, "leftAttack")
				pawnMoves = append(pawnMoves, pawnMoveLeftAttack)
				pawnMoveRightAttack := movePawn(board, row, col, player, "rightAttack")
				pawnMoves = append(pawnMoves, pawnMoveRightAttack)
				genNewBranches(pawnMoves, moves)
				// Knights
			case string(board[row][col][0]) == player && string(board[row][col][1]) == "N":
				knightMoves := genKnightMoves(board, row, col, enemy)
				genNewBranches(knightMoves, moves)
				// Bishops
			case string(board[row][col][0]) == player && string(board[row][col][1]) == "B":
				bishopMoves := genBishopMoves(board, row, col, enemy)
				genNewBranches(bishopMoves, moves)
				// Rooks
			case string(board[row][col][0]) == player && string(board[row][col][1]) == "R":
				rookMoves := genRookMoves(board, row, col, enemy)
				genNewBranches(rookMoves, moves)
				// Queens
			case string(board[row][col][0]) == player && string(board[row][col][1]) == "Q":
				// I'm proud of this part
				queenDiagMoves := genBishopMoves(board, row, col, enemy)
				queenStraightMoves := genRookMoves(board, row, col, enemy)
				genNewBranches(queenDiagMoves, moves)
				genNewBranches(queenStraightMoves, moves)
				// Kings
			case string(board[row][col][0]) == player && string(board[row][col][1]) == "K":
				kingMoves := genKingMoves(board, row, col, player, enemy, genKing)
				genNewBranches(kingMoves, moves)
			}
		}
	}
	return moves
}

func withinBoundaries(moveRow int, moveCol int) bool {
	if (moveRow >= 0 && moveRow < 8) && (moveCol >= 0 && moveCol < 8) {
		return true
	} else {
		return false
	}
}
