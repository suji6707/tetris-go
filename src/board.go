package main

import "fmt"

const (
	WIDTH      = 10
	HEIGHT     = 20
	SIDE_WIDTH = 5
)

type Board struct {
	backBoard    []int
	sideBoard    []int
	currentShape *Shape
	nextShape    *Shape
	currentX     int
	currentY     int
	score        int
}

func NewBoard() *Board {
	backBoard := make([]int, WIDTH*HEIGHT)
	sideBoard := make([]int, SIDE_WIDTH*HEIGHT)
	return &Board{
		backBoard:    backBoard,
		sideBoard:    sideBoard,
		currentShape: NewRandomShape(),
		nextShape:    NewRandomShape(),
		score:        0,
	}
}

func (b *Board) PrintCell(val int) {
	switch val {
	case 0:
		fmt.Print("⬛️")
	case 1:
		fmt.Print("🟦")
	case 2:
		fmt.Print("🟥")
	case 3:
		fmt.Print("🟩")
	case 4:
		fmt.Print("🟪")
	case 5:
		fmt.Print("🟨")
	case 6:
		fmt.Print("🟫")
	case 7:
		fmt.Print("⬜️")
	}
}

func (b *Board) Draw() {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH+1+SIDE_WIDTH; x++ {
			if x < WIDTH {
				b.PrintCell(b.backBoard[y*WIDTH+x])
			} else if x == WIDTH {
				fmt.Print("|")
			} else {
				b.PrintCell(b.sideBoard[y*SIDE_WIDTH+x-WIDTH-1])
			}
		}
		fmt.Println()
	}
	fmt.Println("Score: ", b.score)
}

func (b *Board) SetNewPiece() {
	b.currentShape = b.nextShape
	b.nextShape = NewRandomShape()
	b.currentX = WIDTH / 2
	b.currentY = 1

	coord := b.currentShape.GetCoord(0, b.currentX, b.currentY)
	// backBoard에 컬러 입히기
	for _, point := range coord {
		if point.x >= 0 && point.x < WIDTH && point.y >= 0 && point.y < HEIGHT {
			b.backBoard[point.y*WIDTH+point.x] = b.currentShape.ShapeType
		}
	}
}

// 테스트용: 보드 초기화
func (b *Board) ClearBoard() {
	for i := range b.backBoard {
		b.backBoard[i] = 0
	}
}
