package main

func main() {
	// 테스트: SetNewPiece 함수가 좌표에 잘 찍히는지 확인
	board := NewBoard()

	// 2. 새 블록 설정하고 출력
	board.SetNewPiece()
	board.Draw()

}
