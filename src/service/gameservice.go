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
const (
	numberOfGrid = 5
	numberInBox  = 75
)

type GameService struct {
	Game bingogame.Game
}

type PlayerInfoResponse struct {
	PlayerOne     bingogame.Player `json:"playerOne"`
	PlayerTwo     bingogame.Player `json:"playerTwo"`
	HistoryPickUp []int            `json:"historyPickUp"`
}

func (g GameService) GetPlayerInfo() PlayerInfoResponse {
	return PlayerInfoResponse{
		PlayerOne:     g.Game.Players[0],
		PlayerTwo:     g.Game.Players[1],
		HistoryPickUp: g.Game.HistoryPickUp,
	}
}

func (gs GameService) NewGame(playerOneName, playerTwoName string) error {
	ticket1 := bingogame.NewTicket(numberOfGrid)
	ticket2 := bingogame.NewTicket(numberOfGrid)
	ticket1 = bingogame.MockTicketNumber(ticket1, 1)
	ticket2 = bingogame.MockTicketNumber(ticket2, 2)
	player1 := bingogame.NewPlayer(playerOneName, ticket1)
	player2 := bingogame.NewPlayer(playerTwoName, ticket2)
	numberBox := bingogame.MockNumberBox()
	allPlayer := []bingogame.Player{player1, player2}
	gs.Game = bingogame.NewGame(allPlayer, numberBox)
	return nil
}
