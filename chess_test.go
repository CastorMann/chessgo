package chessgo

import "testing"

func TestNewChessBoard(t *testing.T) {
	board := NewChessBoard()

	board.Print()
}
