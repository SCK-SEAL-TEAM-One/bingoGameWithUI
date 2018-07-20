package bingogame
type Play struct{
	Number int `json:"number"`
	Winner string `json:"winner"`

}
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

func (g *Game) Play() PlayResponse {

	playResponse := PlayResponse{}

	pickupNumber := g.PickUpNumber()
	positionXPlayer1, positionYPlayer1 := g.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		g.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if g.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			playResponse.Winner = g.Players[0].Name
		}
	}

	positionXPlayer2, positionYPlayer2 := g.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		g.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if g.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			playResponse.Winner = g.Players[1].Name
		}
	}
	playResponse.Number = pickupNumber
	return playResponse
}

func MockNumberBox() []int {
	return []int{9, 51, 47, 29, 56, 49, 39, 58}
}

func MockNumberBoxForCase2() []int {
	return []int{45, 14, 41, 32, 11, 36}
}
