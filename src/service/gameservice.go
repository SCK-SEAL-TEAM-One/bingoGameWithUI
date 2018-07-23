package service

import (
	"bingogame"
)

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
