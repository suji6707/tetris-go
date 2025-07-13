package main

import (
	"math/rand"
)

type Point struct {
	x, y int
}

type Shape struct {
	ShapeType int
}

const (
	ShapeNone = 0
	ShapeI    = 1
	ShapeL    = 2
	ShapeJ    = 3
	ShapeT    = 4
	ShapeO    = 5
	ShapeS    = 6
	ShapeZ    = 7
)

var ShapeCord = [8][4]Point{
	// ShapeNone
	{{0, 0}, {0, 0}, {0, 0}, {0, 0}},
	// ShapeI
	{{0, -1}, {0, 0}, {0, 1}, {0, 2}},
	// ShapeL
	{{0, -1}, {0, 0}, {0, 1}, {1, 1}},
	// ShapeJ
	{{0, -1}, {0, 0}, {0, 1}, {-1, 1}},
	// ShapeT
	{{0, -1}, {0, 0}, {0, 1}, {1, 0}},
	// ShapeO
	{{0, 0}, {0, -1}, {1, 0}, {1, -1}},
	// ShapeS
	{{0, 0}, {0, -1}, {-1, 0}, {1, -1}},
	// ShapeZ
	{{0, 0}, {0, -1}, {1, 0}, {-1, -1}},
}

func NewShape(shapeType int) *Shape {
	return &Shape{
		ShapeType: shapeType,
	}
}

func NewRandomShape() *Shape {
	return NewShape(rand.Intn(7) + 1) // 1~7
}

func (s *Shape) GetCoord(direction int, x, y int) [4]Point {
	shapeCord := ShapeCord[s.ShapeType]
	for i := range shapeCord {
		shapeCord[i].x += x
		shapeCord[i].y += y
	}
	return shapeCord
}
