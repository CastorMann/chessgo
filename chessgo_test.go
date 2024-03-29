package chessgo

import (
	"fmt"
	"testing"
)

func TestPawnMoves(t *testing.T) {
	var fromCell BitBoard = A2 | B3 | C7 | D8
	var expectedMoves BitBoard = A3 | A4 | B4 | C8
	var piece Piece = WHITE_PAWN
	var actualMoves BitBoard = piece.GetMoves(fromCell)

	if expectedMoves != actualMoves {
		t.Errorf("Expected %X moves, got %X", expectedMoves, actualMoves)
	}
}

func TestKnightMoves(t *testing.T) {
	var fromCell BitBoard = A2 | E5 | H7
	var expectedMoves BitBoard = B4 | C3 | D3 | C4 | C6 | D7 | F7 | G6 | G4 | F3 | G5 | F6 | F8 | C1
	var piece Piece = WHITE_KNIGHT
	var actualMoves BitBoard = piece.GetMoves(fromCell)
	var difference BitBoard = (expectedMoves ^ actualMoves)

	if expectedMoves != actualMoves {
		t.Errorf("Knights Placements:\n%s\nExpected:\n%s\nGot:\n%s\nDifference:\n%s\n", fromCell.Visualize(), expectedMoves.Visualize(), actualMoves.Visualize(), difference.Visualize())
	}
}

func TestBishopMoves(t *testing.T) {
	var fromCell BitBoard = A2 | E5 | H7
	var expectedMoves BitBoard = B1 | B3 | C4 | D5 | E6 | F7 | G8 | D6 | C7 | B8 | F4 | G3 | H2 | A1 | B2 | C3 | D4 | F6 | G7 | H8 | G8 | G6 | F5 | E4 | D3 | C2 | B1
	var piece Piece = WHITE_BISHOP
	var actualMoves BitBoard = piece.GetMoves(fromCell)
	var difference BitBoard = (expectedMoves ^ actualMoves)

	if expectedMoves != actualMoves {
		t.Errorf("Bishops Placements:\n%s\nExpected:\n%s\nGot:\n%s\nDifference:\n%s\n", fromCell.Visualize(), expectedMoves.Visualize(), actualMoves.Visualize(), difference.Visualize())
	}
}

func TestRookMoves(t *testing.T) {
	var fromCell BitBoard = A2 | E5 | H7
	var expectedMoves BitBoard = FILE_A | FILE_E | FILE_H | RANK_2 | RANK_5 | RANK_7 ^ (A2 | E5 | H7)
	var piece Piece = WHITE_ROOK
	var actualMoves BitBoard = piece.GetMoves(fromCell)
	var difference BitBoard = (expectedMoves ^ actualMoves)

	if expectedMoves != actualMoves {
		t.Errorf("Rooks Placements:\n%s\nExpected:\n%s\nGot:\n%s\nDifference:\n%s\n", fromCell.Visualize(), expectedMoves.Visualize(), actualMoves.Visualize(), difference.Visualize())
	}
}

func TestChessGame(t *testing.T) {
	cb := NewChessBoard()

	fmt.Printf("%v\n", cb.ToString())
	cb.MakeMove(E2, E4)
	fmt.Printf("%v\n", cb.ToString())
	fmt.Printf("check: %v\n", cb.IsBlackChecked())
	cb.MakeMove(D7, D5)
	fmt.Printf("%v\n", cb.ToString())
	fmt.Printf("check: %v\n", cb.IsBlackChecked())
	cb.MakeMove(G1, F3)
	fmt.Printf("%v\n", cb.ToString())
	fmt.Printf("check: %v\n", cb.IsBlackChecked())
	cb.MakeMove(E8, E7)
	fmt.Printf("%v\n", cb.ToString())
	fmt.Printf("check: %v\n", cb.IsBlackChecked())
	cb.MakeMove(F1, B5)
	fmt.Printf("%v\n", cb.ToString())
	fmt.Printf("check: %v\n", cb.IsBlackChecked())
}
