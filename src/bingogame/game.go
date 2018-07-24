package bingogame

const notfound = -1

func NewNumberBox(endNumber int) []int {
	startNumber := 1
	return Shuffle(startNumber, endNumber)
}

func NewGame(players []Player, numberBox []int) Game {
	return Game{Players: players, NumberBox: numberBox, HistoryPickUp: []int{}}
}

func (g *Game) PickUpNumber() int {
	var pickUpNumber int
	pickUpNumber, g.NumberBox = g.NumberBox[0], g.NumberBox[1:]
	g.HistoryPickUp = append(g.HistoryPickUp, pickUpNumber)
	return pickUpNumber
}

func (g *Game) Play() PlayResponse {
	playResponse := PlayResponse{}
	pickupNumber := g.PickUpNumber()
	playResponse.Number = pickupNumber
	for index, _ := range g.Players {
		positionX, positionY := g.Players[index].CheckNumber(pickupNumber)
		if isInTicket(positionX, positionY) {
			g.Players[index].Mark(positionX, positionY)
			if g.Players[index].GetBingo(positionX, positionY) {
				playResponse.Winner = g.Players[index].Name
			}
		}
	}
	return playResponse
}

func isInTicket(positionX, positionY int) bool {
	return positionX != notfound && positionY != notfound
}
func MockNumberBox() []int {
	return []int{9, 51, 47, 29, 56, 49, 39, 58}
}

func MockNumberBoxForCase2() []int {
	return []int{45, 14, 41, 32, 11, 36}
}
