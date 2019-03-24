package main

import (
	"fmt"

	Logic "github.com/siredmar/connect-four-go/internal/logic"
	Player "github.com/siredmar/connect-four-go/internal/player"
)

func main() {

	P1 := Player.CreatePlayer("Player 1", "X")
	P2 := Player.CreatePlayer("Player 2", "O")

	var players []Player.Player
	players = append(players, P1, P2)
	fmt.Println(players)
	l := Logic.Create()
	l.Init(8, 8)

	l.InsertCoin(players[0], 0)
	l.InsertCoin(players[1], 1)
	l.InsertCoin(players[0], 2)
	l.InsertCoin(players[0], 3)
	l.InsertCoin(players[0], 2)
	l.InsertCoin(players[0], 3)
	l.InsertCoin(players[0], 3)
	l.InsertCoin(players[1], 1)
	l.InsertCoin(players[1], 2)
	l.InsertCoin(players[1], 3)
	l.InsertCoin(players[0], 4)
	l.InsertCoin(players[1], 4)
	l.InsertCoin(players[1], 4)
	l.InsertCoin(players[1], 4)
	l.InsertCoin(players[1], 4)
	l.Print()
	fmt.Println("Win: ", l.Check(players[1]))
	// var s bool
	// var err error
	// s, err = l.InsertCoin(players[1], 0)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// fmt.Println("Win: ", l.Check(P1))

	// s, err = l.InsertCoin(players[0], 1)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// fmt.Println("Win: ", l.Check(P1))
	// s, err = l.InsertCoin(players[0], 2)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// fmt.Println("Win: ", l.Check(P1))
	// s, err = l.InsertCoin(players[0], 3)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// fmt.Println("Win: ", l.Check(P1))
	// s, err = l.InsertCoin(players[0], 4)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// fmt.Println("Win: ", l.Check(P1))
	// s, err = l.InsertCoin(players[0], 7)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// s, err = l.InsertCoin(players[0], 7)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// s, err = l.InsertCoin(players[0], 7)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }
	// s, err = l.InsertCoin(players[0], 7)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }

	// s, err = l.InsertCoin(players[0], 1)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }

	// s, err = l.InsertCoin(players[1], 1)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }

	// s, err = l.InsertCoin(players[1], 2)
	// if err != nil {
	// 	fmt.Println(s, err)
	// } else {
	// 	fmt.Println(s)
	// 	l.Print()
	// }

	// l.Rotate()
	// l.Print()
	// fmt.Println("Win: ", l.Check(P1))
}
