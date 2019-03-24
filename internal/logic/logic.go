package logic

import (
	"fmt"
	"math"

	Player "github.com/siredmar/connect-four-go/internal/player"
)

var (
	xMax = 8
	yMax = 8
)

type Logic struct {
	board [][]string
}

func Create() Logic {
	return Logic{}
}

// Rotate rotates the board to prepare the board for a diagonal win check
// based on information from:
// https://code.likeagirl.io/rotate-an-2d-matrix-90-degree-clockwise-without-create-another-array-49209ea8b6e6
func (l Logic) Rotate() {
	n := len(l.board)
	x := int(math.Floor(float64(n) / 2))
	y := n - 1
	for i := 0; i < x; i++ {
		for j := i; j < y-i; j++ {
			k := l.board[i][j]
			l.board[i][j] = l.board[y-j][i]
			l.board[y-j][i] = l.board[y-i][y-j]
			l.board[y-i][y-j] = l.board[j][y-i]
			l.board[j][y-i] = k
		}
	}
}

func (l *Logic) Init(x, y int) {
	l.board = make([][]string, yMax)
	for y := range l.board {
		l.board[y] = make([]string, xMax)
	}
	for y := 0; y < yMax; y++ {
		for x := 0; x < xMax; x++ {
			l.board[y][x] = "."
		}

	}
}

func (l Logic) Print() {
	fmt.Printf("  ")
	for i := 0; i < xMax; i++ {
		fmt.Printf("%v", i)
	}
	fmt.Println()
	for y := yMax - 1; y >= 0; y-- {
		fmt.Printf("%v ", y)
		for x := 0; x < xMax; x++ {
			fmt.Printf("%v", l.board[y][x])
		}
		fmt.Println()
	}
}

func (l Logic) getLastYPosition(x int) (int, error) {
	position := yMax - 1
	for y := yMax - 1; y >= 0; y-- {
		position = y
		if l.board[y][x] != "." {
			position++
			break
		}

	}
	if position >= yMax {
		return -1, fmt.Errorf("Column already filled up with coins")
	} else {
		return position, nil
	}
}

func (l *Logic) InsertCoin(p Player.Player, x int) (bool, error) {
	if x > xMax {
		return false, fmt.Errorf("x bigger than maximum board size: %v", xMax)
	}
	lastYPosition, err := l.getLastYPosition(x)
	if err != nil {
		return false, err
	} else {
		l.board[lastYPosition][x] = p.Sign
		return true, nil
	}
}
