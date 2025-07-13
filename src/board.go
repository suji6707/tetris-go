package main

import (
	"fmt"
	"time"
)

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
	sideX        int
	sideY        int
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
	b.sideX = SIDE_WIDTH / 2
	b.sideY = HEIGHT / 2

	// backBoard에 컬러 입히기
	b.Colorize(b.currentX, b.currentY)
	// sideBoard에 컬러 입히기
	sideCoord := b.nextShape.GetCoord(0, b.sideX, b.sideY)
	for _, point := range sideCoord {
		b.sideBoard[point.y*SIDE_WIDTH+point.x] = b.nextShape.ShapeType
	}
}

func (b *Board) Colorize(x, y int) {
	coord := b.currentShape.GetCoord(0, x, y)
	for _, point := range coord {
		b.backBoard[point.y*WIDTH+point.x] = b.currentShape.ShapeType
	}
}

func (b *Board) ClearPiece(x, y int) {
	coord := b.currentShape.GetCoord(0, x, y)
	for _, point := range coord {
		b.backBoard[point.y*WIDTH+point.x] = 0
	}
}

// 보드 초기화
func (b *Board) ClearBoard() {
	for i := range b.backBoard {
		b.backBoard[i] = 0
	}
}

func (b *Board) MovePiece(oldX, oldY, newX, newY int) {
	b.currentX = newX
	b.currentY = newY
	b.ClearPiece(oldX, oldY)
	b.Colorize(newX, newY)
}

func (b *Board) tryGoDown(newX, newY int) bool {
	coord := b.currentShape.GetCoord(0, newX, newY)
	for _, point := range coord {
		// 1. boundary check
		if point.x < 0 || point.x >= WIDTH || point.y < 0 || point.y >= HEIGHT {
			return false
		}
		// 2. collision check
		val := b.backBoard[point.y*WIDTH+point.x]
		if val != 0 && val != b.currentShape.ShapeType {
			return false
		}
	}
	return true
}

func (b *Board) GoDown() bool {
	oldX, oldY := b.currentX, b.currentY
	newX := oldX
	newY := oldY + 1

	if b.tryGoDown(newX, newY) {
		b.MovePiece(oldX, oldY, newX, newY)
		ClearScreen()
		b.Draw()
		time.Sleep(500 * time.Millisecond)
		return true
	} else {
		return false
	}
}
