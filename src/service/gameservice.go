package service

import (
	"bingogame"
)

type PlayGameService struct {
	Game bingogame.Game
}

func (g *PlayGameService) PlayGame() bingogame.PlayResponse {
	playResponse := bingogame.PlayResponse{}
	pickupNumber := g.Game.PickUpNumber()

	positionXPlayer1, positionYPlayer1 := g.Game.Players[0].CheckNumber(pickupNumber)
	if positionXPlayer1 != -1 && positionYPlayer1 != -1 {
		g.Game.Players[0].Mark(positionXPlayer1, positionYPlayer1)
		if g.Game.Players[0].GetBingo(positionXPlayer1, positionYPlayer1) {
			playResponse.Winner = g.Game.Players[0].Name
		}
	}

	positionXPlayer2, positionYPlayer2 := g.Game.Players[1].CheckNumber(pickupNumber)
	if positionXPlayer2 != -1 && positionYPlayer2 != -1 {
		g.Game.Players[1].Mark(positionXPlayer2, positionYPlayer2)
		if g.Game.Players[1].GetBingo(positionXPlayer2, positionYPlayer2) {
			playResponse.Winner = g.Game.Players[1].Name
		}
	}
	playResponse.Number = pickupNumber
	return playResponse
}
