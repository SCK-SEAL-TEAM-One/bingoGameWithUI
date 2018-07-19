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
