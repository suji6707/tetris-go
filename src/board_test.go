package main

import "testing"

func TestSetNewPiece(t *testing.T) {
	board := NewBoard()
	board.ClearBoard()

	// 새 블록 설정
	board.SetNewPiece()
	board.Draw()
}
