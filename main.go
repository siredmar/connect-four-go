package main

import (
	"fmt"

	Logic "github.com/siredmar/connect-four/internal/logic"
	Player "github.com/siredmar/connect-four/internal/player"
)

func main() {

	P1 := Player.CreatePlayer("Player 1", "X")
	P2 := Player.CreatePlayer("Player 2", "O")

	var players []Player.Player
	players = append(players, P1, P2)
	fmt.Println(players)
	l := Logic.Create()
	l.Init(8, 8)
	s, err := l.InsertCoin(players[0], 0)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}

	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}
	s, err = l.InsertCoin(players[0], 7)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}

	s, err = l.InsertCoin(players[0], 1)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}

	s, err = l.InsertCoin(players[1], 1)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}

	s, err = l.InsertCoin(players[1], 2)
	if err != nil {
		fmt.Println(s, err)
	} else {
		fmt.Println(s)
		l.Print()
	}

	l.Rotate()
	l.Print()
}
