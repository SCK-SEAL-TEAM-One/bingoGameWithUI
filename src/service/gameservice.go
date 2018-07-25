package service

import (
	"bingogame"
)

const (
	numberOfGrid = 5
	numberInBox  = 75
)

type Service interface {
	NewGame(playerNames []string) error
	GetPlayerInfo() PlayerInfoResponse
	PlayGame() bingogame.PlayResponse
	ChangeTicket(playerName string) bingogame.Player
}

type PlayerInfoResponse struct {
	Players       []bingogame.Player `json:"players"`
	HistoryPickUp []int              `json:"historyPickUp"`
}

type MockGameService struct {
	Game bingogame.Game
}

func (gs *MockGameService) NewGame(playerNames []string) error {
	allPlayer := []bingogame.Player{}
	for index, playerName := range playerNames {
		ticket := bingogame.NewTicket(numberOfGrid)
		ticket = bingogame.MockTicketNumber(ticket, index+1)
		player := bingogame.NewPlayer(playerName, ticket)
		allPlayer = append(allPlayer, player)
	}
	numberBox := bingogame.MockNumberBox()
	gs.Game = bingogame.NewGame(allPlayer, numberBox)
	return nil
}

func (g MockGameService) GetPlayerInfo() PlayerInfoResponse {
	return PlayerInfoResponse{
		Players:       g.Game.Players,
		HistoryPickUp: g.Game.HistoryPickUp,
	}
}

func (g *MockGameService) PlayGame() bingogame.PlayResponse {
	return g.Game.Play()
}

func (g *MockGameService) ChangeTicket(playerName string) bingogame.Player {
	player := bingogame.Player{}
	for index := range g.Game.Players {
		if g.Game.Players[index].Name == playerName {
			player = g.Game.Players[index]
			player.Ticket = bingogame.MockTicketNumber(player.Ticket, 3)
			break
		}
	}
	return player
}

type GameService struct {
	Game bingogame.Game
}

func (gs *GameService) NewGame(playerNames []string) error {
	allPlayer := []bingogame.Player{}
	for _, playerName := range playerNames {
		ticket := bingogame.NewTicket(numberOfGrid)
		ticket = bingogame.GenerateTicketNumber(ticket)
		player := bingogame.NewPlayer(playerName, ticket)
		allPlayer = append(allPlayer, player)
	}
	numberBox := bingogame.NewNumberBox(numberInBox)
	gs.Game = bingogame.NewGame(allPlayer, numberBox)
	return nil
}

func (g GameService) GetPlayerInfo() PlayerInfoResponse {
	return PlayerInfoResponse{
		Players:       g.Game.Players,
		HistoryPickUp: g.Game.HistoryPickUp,
	}
}

func (g *GameService) PlayGame() bingogame.PlayResponse {
	return g.Game.Play()
}
func (g *GameService) ChangeTicket(playerName string) bingogame.Player {
	player := bingogame.Player{}
	for index := range g.Game.Players {
		if g.Game.Players[index].Name == playerName {
			player = g.Game.Players[index]
			player.Ticket = bingogame.GenerateTicketNumber(player.Ticket)
			break
		}
	}
	return player
}
