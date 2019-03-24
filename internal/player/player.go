package player

type Player struct {
	Name string
	Sign string
}

func CreatePlayer(name string, sign string) Player {
	return Player{
		Name: name,
		Sign: sign,
	}
}
