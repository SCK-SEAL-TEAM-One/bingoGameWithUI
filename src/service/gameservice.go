package service

import (
	"api"
	"bingogame"
)

type GameService struct {
	Game bingogame.Game
}

func (g GameService) GetPlayerInfo() api.PlayerInfoResponse {
	return api.PlayerInfoResponse{
		PlayerOne:     g.Game.Players[0],
		PlayerTwo:     g.Game.Players[1],
		HistoryPickUp: g.Game.HistoryPickUp,
	}
}
