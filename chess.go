package chessgo

import "fmt"

type BitBoard uint64

func (b *BitBoard) SetBit(pos uint) {
	*b |= 1 << pos
}

func (b *BitBoard) ClearBit(pos uint) {
	*b &^= 1 << pos
}

func (b BitBoard) IsSet(pos uint) bool {
	return (b>>pos)&1 == 1
}

const (
	A1    BitBoard = 1 << 0
	B1             = 1 << 1
	C1             = 1 << 2
	D1             = 1 << 3
	E1             = 1 << 4
	F1             = 1 << 5
	G1             = 1 << 6
	H1             = 1 << 7
	A2             = 1 << 8
	B2             = 1 << 9
	C2             = 1 << 10
	D2             = 1 << 11
	E2             = 1 << 12
	F2             = 1 << 13
	G2             = 1 << 14
	H2             = 1 << 15
	A3             = 1 << 16
	B3             = 1 << 17
	C3             = 1 << 18
	D3             = 1 << 19
	E3             = 1 << 20
	F3             = 1 << 21
	G3             = 1 << 22
	H3             = 1 << 23
	A4             = 1 << 24
	B4             = 1 << 25
	C4             = 1 << 26
	D4             = 1 << 27
	E4             = 1 << 28
	F4             = 1 << 29
	G4             = 1 << 30
	H4             = 1 << 31
	A5             = 1 << 32
	B5             = 1 << 33
	C5             = 1 << 34
	D5             = 1 << 35
	E5             = 1 << 36
	F5             = 1 << 37
	G5             = 1 << 38
	H5             = 1 << 39
	A6             = 1 << 40
	B6             = 1 << 41
	C6             = 1 << 42
	D6             = 1 << 43
	E6             = 1 << 44
	F6             = 1 << 45
	G6             = 1 << 46
	H6             = 1 << 47
	A7             = 1 << 48
	B7             = 1 << 49
	C7             = 1 << 50
	D7             = 1 << 51
	E7             = 1 << 52
	F7             = 1 << 53
	G7             = 1 << 54
	H7             = 1 << 55
	A8             = 1 << 56
	B8             = 1 << 57
	C8             = 1 << 58
	D8             = 1 << 59
	E8             = 1 << 60
	F8             = 1 << 61
	G8             = 1 << 62
	H8             = 1 << 63
	ROW_1          = 0x00000000000000FF
	ROW_2          = 0x000000000000FF00
	ROW_3          = 0x0000000000FF0000
	ROW_4          = 0x00000000FF000000
	ROW_5          = 0x000000FF00000000
	ROW_6          = 0x0000FF0000000000
	ROW_7          = 0x00FF000000000000
	ROW_8          = 0xFF00000000000000
	COL_A          = A1 | A2 | A3 | A4 | A5 | A6 | A7 | A8
	COL_B          = B1 | B2 | B3 | B4 | B5 | B6 | B7 | B8
	COL_C          = C1 | C2 | C3 | C4 | C5 | C6 | C7 | C8
	COL_D          = D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8
	COL_E          = E1 | E2 | E3 | E4 | E5 | E6 | E7 | E8
	COL_F          = F1 | F2 | F3 | F4 | F5 | F6 | F7 | F8
	COL_G          = G1 | G2 | G3 | G4 | G5 | G6 | G7 | G8
	COL_H          = H1 | H2 | H3 | H4 | H5 | H6 | H7 | H8
)

type Piece uint

const (
	Pawn   Piece = 1 << 0
	Knight       = 1 << 1
	Bishop       = 1 << 2
	Rook         = 1 << 3
	Queen        = 1 << 4
	King         = 1 << 5
	Black        = 1 << 6 // flag for determining if the piece is white or black - if flag is set then the piece is black
	None         = 0
	All          = Pawn | Knight | Bishop | Rook | Queen | King
)

func (p *Piece) ToString() string {
	switch *p {
	case Pawn:
		return "P"
	case Knight:
		return "N"
	case Bishop:
		return "B"
	case Rook:
		return "R"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Pawn | Black:
		return "p"
	case Knight | Black:
		return "n"
	case Bishop | Black:
		return "b"
	case Rook | Black:
		return "r"
	case Queen | Black:
		return "q"
	case King | Black:
		return "k"
	case None:
		return " "
	}
	return "?"
}

type ChessBoard struct {
	whitePawns   BitBoard
	whiteKnights BitBoard
	whiteBishops BitBoard
	whiteRooks   BitBoard
	whiteQueens  BitBoard
	whiteKings   BitBoard

	blackPawns   BitBoard
	blackKnights BitBoard
	blackBishops BitBoard
	blackRooks   BitBoard
	blackQueens  BitBoard
	blackKings   BitBoard
}

func (cb *ChessBoard) GetPiece(cell BitBoard) Piece {
	if cb.whitePawns&cell == cell {
		return Pawn
	}
	if cb.whiteKnights&cell == cell {
		return Knight
	}
	if cb.whiteBishops&cell == cell {
		return Bishop
	}
	if cb.whiteRooks&cell == cell {
		return Rook
	}
	if cb.whiteQueens&cell == cell {
		return Queen
	}
	if cb.whiteKings&cell == cell {
		return King
	}

	if cb.blackPawns&cell == cell {
		return Pawn | Black
	}
	if cb.blackKnights&cell == cell {
		return Knight | Black
	}
	if cb.blackBishops&cell == cell {
		return Bishop | Black
	}
	if cb.blackRooks&cell == cell {
		return Rook | Black
	}
	if cb.blackQueens&cell == cell {
		return Queen | Black
	}
	if cb.blackKings&cell == cell {
		return King | Black
	}
	return None
}

func (cb *ChessBoard) Print() {
	for row := 7; row >= 0; row-- {
		for col := 0; col < 8; col++ {
			// Calculate the cell position
			cell := BitBoard(1 << (row*8 + col))

			// Get the piece at the cell
			piece := cb.GetPiece(cell)

			// Print the piece or a placeholder if the cell is empty
			fmt.Print(piece.ToString() + " ")
		}
		fmt.Println()
	}
}

func NewChessBoard() *ChessBoard {
	board := ChessBoard{
		whitePawns:   ROW_2,
		whiteKnights: B1 | G1,
		whiteBishops: C1 | F1,
		whiteRooks:   A1 | H1,
		whiteQueens:  D1,
		whiteKings:   E1,

		blackPawns:   ROW_7,
		blackKnights: B8 | G8,
		blackBishops: C8 | F8,
		blackRooks:   A8 | H8,
		blackQueens:  D8,
		blackKings:   E8,
	}

	return &board
}
