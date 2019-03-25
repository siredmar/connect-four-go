package logic

import (
	"fmt"
	"math"

	Player "github.com/siredmar/connect-four-go/internal/player"
)

var (
	xMax        int
	yMax        int
	numWinCoins = 4
)

type Logic struct {
	board [][]string
}

// Create makes an instance of the Logic struct and returns it
func Create() Logic {
	return Logic{}
}

// Rotate rotates the board to prepare the board for a horizontal/vertical/diagonal winning check
// Algorithm based on information from:
// https://code.likeagirl.io/rotate-an-2d-matrix-90-degree-clockwise-without-create-another-array-49209ea8b6e6
func (l Logic) rotate() {
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

// Check checks if the a given player won the game
func (l Logic) Check(p Player.Player) bool {
	win := false
	// We are rotating the board to make simple winning check functions.
	// Rotate the board 4 times. On each rotation the winning checks are performed.
	// On the last rotation the board is back in its original state.
	for r := 0; r < 4; r++ {
		if l.checkHorizontal(p) == true || l.checkDiagonal(p) == true {
			win = true
		}
		l.rotate()
	}
	return win
}

func (l Logic) checkHorizontal(p Player.Player) bool {
	for y := 0; y < yMax; y++ {
		connectedCoins := 0
		for x := 0; x < xMax-numWinCoins; x++ {
			for c := x; c < x+numWinCoins; c++ {
				if l.board[y][c] == p.Sign {
					connectedCoins++
				}
			}
			if connectedCoins >= numWinCoins {
				return true
			}
			connectedCoins = 0
		}
	}
	return false
}

func (l Logic) checkDiagonal(p Player.Player) bool {
	for y := 0; y < yMax-numWinCoins; y++ {
		for x := 0; x < xMax-numWinCoins; x++ {
			connectedCoins := 0
			for c := 0; c < numWinCoins; c++ {
				if l.board[y+c][x+c] == p.Sign {
					connectedCoins++
				}
			}
			if connectedCoins >= numWinCoins {
				return true
			}
		}
	}
	return false
}

// Init makes the board an initializes to default values
func (l *Logic) Init(x, y int) {
	xMax = x
	yMax = y
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

// Print prints the board at its current state
func (l Logic) Print() {
	fmt.Printf("  ")
	for i := 1; i <= xMax; i++ {
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
	fmt.Println()
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

// InsertCoin lets a given player insert a coin in a specific column of the board
func (l *Logic) InsertCoin(p Player.Player, x int) (bool, error) {
	if x >= xMax || x < 0 {
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
