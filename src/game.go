package main

// import (
// 	"github.com/eiannone/keyboard"
// )

type Game struct {
	board *Board
	Score int
	// IsStarted bool
	// IsPaused  bool
	IsGameOver bool
	// Speed     time.Duration
}

func NewGame() *Game {
	return &Game{
		board: NewBoard(),
		Score: 0,
	}
}

func (g *Game) Start() {
	g.board.SetNewPiece()
	g.board.Draw()
}

func (g *Game) PlayLoop() {
	g.Start()

	// TODO: 키 입력에 따라 게임 상태 변경
	for !g.IsGameOver {
		success := g.board.GoDown()
		if !success {
			g.IsGameOver = true
		}
	}
}
