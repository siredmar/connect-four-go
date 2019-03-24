package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	Logic "github.com/siredmar/connect-four-go/internal/logic"
	Player "github.com/siredmar/connect-four-go/internal/player"
)

var (
	currentPlayer int
)

func main() {

	currentPlayer = 0

	var players []Player.Player
	players = append(players, Player.CreatePlayer("", "X"), Player.CreatePlayer("", "0"))

	players[0].Name = userInput("Player 1, enter you name: ")
	players[1].Name = userInput("Player 2, enter you name: ")
	fmt.Println(players[currentPlayer].Name, "begins. Your sign is:", players[currentPlayer].Sign)
	l := Logic.Create()
	l.Init(8, 8)
	l.Print()
	for {
		var column int
		for {
			fmt.Printf("%v, column: ", players[currentPlayer].Name)
			fmt.Scanf("%d", &column)
			// column-1 to compensate for beginning counting at zero
			r, err := l.InsertCoin(players[currentPlayer], column-1)
			if err != nil {
				fmt.Println(err)
			}
			if r == true {
				break
			}
		}
		l.Print()
		if l.Check(players[currentPlayer]) == true {
			fmt.Println(players[currentPlayer].Name, "won!")
			break
		}
		switchCurrentPlayer()
	}
}

func switchCurrentPlayer() {
	if currentPlayer == 0 {
		currentPlayer = 1
	} else {
		currentPlayer = 0
	}
}

func userInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	if prompt != "" {
		fmt.Print(prompt)
	}
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	return text

}
