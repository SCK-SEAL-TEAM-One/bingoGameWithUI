package bingogame

func NewNumberBox(endNumber int) []int {
	startNumber := 1
	return Shuffle(startNumber, endNumber)
}

func NewGame(players []Player, numberBox []int) Game {
	return Game{Players: players, NumberBox: numberBox}
}

func (g *Game) PickUpNumber() int {
	var pickUpNumber int
	pickUpNumber, g.NumberBox = g.NumberBox[0], g.NumberBox[1:]
	g.HistoryPickUp = append(g.HistoryPickUp, pickUpNumber)
	return pickUpNumber
}
func MockNumberBox() []int {
	return []int{9, 51, 47, 29, 56, 49, 39, 58}
}

func MockNumberBoxForCase2() []int {
	return []int{45, 14, 41, 32, 11, 36}
}
