package chess

import (
	"strconv"
)

type PieceType uint8

const (
	Pawn PieceType = iota + 1
	Knight
	Bishop
	Rook
	Queen
	King
)

var Figurines = map[Color]map[PieceType]rune{
	Black: {Pawn: '♟', Knight: '♞', Bishop: '♝', Rook: '♜', Queen: '♛', King: '♚'},
	White: {Pawn: '♙', Knight: '♘', Bishop: '♗', Rook: '♖', Queen: '♕', King: '♔'},
}

func (t PieceType) String() string {
	if s, ok := i18n[Language][t]; ok {
		return s.Name
	}
	return "(" + strconv.Itoa(int(t)) + ")"
}

func (t PieceType) Letter() rune {
	if s, ok := i18n[Language][t]; ok {
		return s.Letter
	}
	return '?'
}

func (t PieceType) Figurine(color Color) rune {
	return Figurines[color][t]
}
