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
	A1              BitBoard = 1 << 0
	B1              BitBoard = 1 << 1
	C1              BitBoard = 1 << 2
	D1              BitBoard = 1 << 3
	E1              BitBoard = 1 << 4
	F1              BitBoard = 1 << 5
	G1              BitBoard = 1 << 6
	H1              BitBoard = 1 << 7
	A2              BitBoard = 1 << 8
	B2              BitBoard = 1 << 9
	C2              BitBoard = 1 << 10
	D2              BitBoard = 1 << 11
	E2              BitBoard = 1 << 12
	F2              BitBoard = 1 << 13
	G2              BitBoard = 1 << 14
	H2              BitBoard = 1 << 15
	A3              BitBoard = 1 << 16
	B3              BitBoard = 1 << 17
	C3              BitBoard = 1 << 18
	D3              BitBoard = 1 << 19
	E3              BitBoard = 1 << 20
	F3              BitBoard = 1 << 21
	G3              BitBoard = 1 << 22
	H3              BitBoard = 1 << 23
	A4              BitBoard = 1 << 24
	B4              BitBoard = 1 << 25
	C4              BitBoard = 1 << 26
	D4              BitBoard = 1 << 27
	E4              BitBoard = 1 << 28
	F4              BitBoard = 1 << 29
	G4              BitBoard = 1 << 30
	H4              BitBoard = 1 << 31
	A5              BitBoard = 1 << 32
	B5              BitBoard = 1 << 33
	C5              BitBoard = 1 << 34
	D5              BitBoard = 1 << 35
	E5              BitBoard = 1 << 36
	F5              BitBoard = 1 << 37
	G5              BitBoard = 1 << 38
	H5              BitBoard = 1 << 39
	A6              BitBoard = 1 << 40
	B6              BitBoard = 1 << 41
	C6              BitBoard = 1 << 42
	D6              BitBoard = 1 << 43
	E6              BitBoard = 1 << 44
	F6              BitBoard = 1 << 45
	G6              BitBoard = 1 << 46
	H6              BitBoard = 1 << 47
	A7              BitBoard = 1 << 48
	B7              BitBoard = 1 << 49
	C7              BitBoard = 1 << 50
	D7              BitBoard = 1 << 51
	E7              BitBoard = 1 << 52
	F7              BitBoard = 1 << 53
	G7              BitBoard = 1 << 54
	H7              BitBoard = 1 << 55
	A8              BitBoard = 1 << 56
	B8              BitBoard = 1 << 57
	C8              BitBoard = 1 << 58
	D8              BitBoard = 1 << 59
	E8              BitBoard = 1 << 60
	F8              BitBoard = 1 << 61
	G8              BitBoard = 1 << 62
	H8              BitBoard = 1 << 63
	RANK_1          BitBoard = 0x00000000000000FF
	RANK_2          BitBoard = 0x000000000000FF00
	RANK_3          BitBoard = 0x0000000000FF0000
	RANK_4          BitBoard = 0x00000000FF000000
	RANK_5          BitBoard = 0x000000FF00000000
	RANK_6          BitBoard = 0x0000FF0000000000
	RANK_7          BitBoard = 0x00FF000000000000
	RANK_8          BitBoard = 0xFF00000000000000
	FILE_A          BitBoard = A1 | A2 | A3 | A4 | A5 | A6 | A7 | A8
	FILE_B          BitBoard = B1 | B2 | B3 | B4 | B5 | B6 | B7 | B8
	FILE_C          BitBoard = C1 | C2 | C3 | C4 | C5 | C6 | C7 | C8
	FILE_D          BitBoard = D1 | D2 | D3 | D4 | D5 | D6 | D7 | D8
	FILE_E          BitBoard = E1 | E2 | E3 | E4 | E5 | E6 | E7 | E8
	FILE_F          BitBoard = F1 | F2 | F3 | F4 | F5 | F6 | F7 | F8
	FILE_G          BitBoard = G1 | G2 | G3 | G4 | G5 | G6 | G7 | G8
	FILE_H          BitBoard = H1 | H2 | H3 | H4 | H5 | H6 | H7 | H8
	CENTER          BitBoard = E4 | D4 | E5 | D5
	CENTER_EXPANDED BitBoard = (FILE_C | FILE_D | FILE_E | FILE_F) & (RANK_3 | RANK_4 | RANK_5 | RANK_6)
	WHITE_HALF      BitBoard = RANK_1 | RANK_2 | RANK_3 | RANK_4
	BLACK_HALF      BitBoard = RANK_5 | RANK_6 | RANK_7 | RANK_8
	QUEEN_SIDE      BitBoard = FILE_A | FILE_B | FILE_C | FILE_D
	KING_SIDE       BitBoard = FILE_E | FILE_F | FILE_G | FILE_H
	CORNER_A1       BitBoard = A1 | A2 | B1 | B2
	CORNER_A8       BitBoard = A7 | A8 | B7 | B8
	CORNER_H1       BitBoard = H1 | H2 | G1 | G2
	CORNER_H8       BitBoard = H7 | H8 | G7 | G8
	ALL_CELLS       BitBoard = 0xFFFFFFFFFFFFFFFF
	DIAGONAL_UP     BitBoard = A1 | B2 | C3 | D4 | E5 | F6 | G7 | H8
	DIAGONAL_DOWN   BitBoard = A8 | B7 | C6 | D5 | E4 | F3 | G2 | H1
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

func (p *Piece) GetMoves(fromCell BitBoard) BitBoard {

	switch *p & All {
	case Pawn:
		return GetPawnMoves(fromCell, *p&Black == 0)
	case Knight:
		return GetKnightMoves(fromCell)
	case Bishop:
		return GetBishopMoves(fromCell)
	case Rook:
		return GetRookMoves(fromCell)
	case Queen:
		return GetQueenMoves(fromCell)
	case King:
		return GetKingMoves(fromCell)
	}
	return 0
}

func GetKingMoves(fromCell BitBoard) BitBoard {
	var bb BitBoard = 0
	bb |= (fromCell & (ALL_CELLS ^ RANK_8)) << 8
	bb |= (fromCell & (ALL_CELLS ^ RANK_1)) >> 8
	bb |= (fromCell & (ALL_CELLS ^ FILE_H)) << 1
	bb |= (fromCell & (ALL_CELLS ^ FILE_A)) >> 1
	bb |= (fromCell & (ALL_CELLS ^ (RANK_8 | FILE_H))) << 9
	bb |= (fromCell & (ALL_CELLS ^ (RANK_8 | FILE_A))) << 7
	bb |= (fromCell & (ALL_CELLS ^ (RANK_1 | FILE_H))) >> 7
	bb |= (fromCell & (ALL_CELLS ^ (RANK_1 | FILE_A))) >> 9
	return bb
}

func GetQueenMoves(fromCell BitBoard) BitBoard {
	return GetRookMoves(fromCell) | GetBishopMoves(fromCell)
}

func GetBishopMoves(fromCell BitBoard) BitBoard {
	var bb BitBoard = 0

	du := DIAGONAL_UP
	dd := DIAGONAL_DOWN
	if du&fromCell != 0 {
		bb |= du
	}
	if dd&fromCell != 0 {
		bb |= dd
	}
	for i := 0; i < 7; i++ {
		du >>= 1
		du &= ALL_CELLS ^ FILE_H
		dd >>= 1
		dd &= ALL_CELLS ^ FILE_H
		if du&fromCell != 0 {
			bb |= du
		}
		if dd&fromCell != 0 {
			bb |= dd
		}
	}
	du = DIAGONAL_UP
	dd = DIAGONAL_DOWN
	for i := 0; i < 7; i++ {
		du <<= 1
		du &= ALL_CELLS ^ FILE_A
		dd <<= 1
		dd &= ALL_CELLS ^ FILE_A
		if du&fromCell != 0 {
			bb |= du
		}
		if dd&fromCell != 0 {
			bb |= dd
		}
	}
	return bb ^ fromCell // todo: fix bug if 2 bishops on same diagonal
}

func GetRookMoves(fromCell BitBoard) BitBoard {
	var bb BitBoard = 0
	for rank := RANK_1; rank != 0; rank <<= 8 {
		if fromCell&rank != 0 {
			bb |= rank ^ (fromCell & rank)
		}
	}
	for file := FILE_A; file != 0; file <<= 1 {
		if fromCell&file != 0 {
			bb |= file ^ (fromCell & file)
		}
	}
	return bb
}

func GetKnightMoves(fromCell BitBoard) BitBoard {
	var bb BitBoard = 0
	bb |= (fromCell & (ALL_CELLS ^ (FILE_H | RANK_8 | RANK_7))) << 17
	bb |= (fromCell & (ALL_CELLS ^ (FILE_A | RANK_8 | RANK_7))) << 15
	bb |= (fromCell & (ALL_CELLS ^ (FILE_H | FILE_G | RANK_8))) << 10
	bb |= (fromCell & (ALL_CELLS ^ (FILE_A | FILE_B | RANK_8))) << 6
	bb |= (fromCell & (ALL_CELLS ^ (FILE_A | RANK_2 | RANK_1))) >> 17
	bb |= (fromCell & (ALL_CELLS ^ (FILE_H | RANK_2 | RANK_1))) >> 15
	bb |= (fromCell & (ALL_CELLS ^ (FILE_H | FILE_G | RANK_1))) >> 6
	bb |= (fromCell & (ALL_CELLS ^ (FILE_A | FILE_B | RANK_1))) >> 10
	return bb
}

func GetPawnMoves(fromCell BitBoard, white bool) BitBoard {
	if white {
		return GetWhitePawnMoves(fromCell)
	}
	return GetBlackPawnMoves(fromCell)
}

func GetBlackPawnMoves(fromCell BitBoard) BitBoard {
	return fromCell>>8 | fromCell&RANK_7>>16
}

func GetWhitePawnMoves(fromCell BitBoard) BitBoard {
	return fromCell<<8 | fromCell&RANK_2<<16
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

func (cb *ChessBoard) ToString() string {
	var sb string = ""
	sb += "   ------------------------------\n"
	for row := 7; row >= 0; row-- {
		sb += fmt.Sprintf("%d |", row+1)
		for col := 0; col < 8; col++ {
			// Calculate the cell position
			cell := BitBoard(1 << (row*8 + col))

			// Get the piece at the cell
			piece := cb.GetPiece(cell)

			// Print the piece or a placeholder if the cell is empty
			sb += fmt.Sprint(" " + piece.ToString() + " |")
		}
		sb += "\n   ------------------------------\n"
	}
	sb += fmt.Sprintln("    A   B   C   D   E   F   G   H")
	return sb
}

func (bb *BitBoard) Visualize() string {
	var sb string = ""
	sb += "   ------------------------------\n"
	for row := 7; row >= 0; row-- {
		sb += fmt.Sprintf("%d |", row+1)
		for col := 0; col < 8; col++ {
			var val string = " "
			if bb.IsSet(uint(row*8 + col)) {
				val = "X"
			}
			// Print the piece or a placeholder if the cell is empty
			sb += fmt.Sprint(" " + val + " |")
		}
		sb += "\n   ------------------------------\n"
	}
	sb += "    A   B   C   D   E   F   G   H"
	return sb
}

func NewChessBoard() *ChessBoard {
	board := ChessBoard{
		whitePawns:   RANK_2,
		whiteKnights: B1 | G1,
		whiteBishops: C1 | F1,
		whiteRooks:   A1 | H1,
		whiteQueens:  D1,
		whiteKings:   E1,

		blackPawns:   RANK_7,
		blackKnights: B8 | G8,
		blackBishops: C8 | F8,
		blackRooks:   A8 | H8,
		blackQueens:  D8,
		blackKings:   E8,
	}

	return &board
}
